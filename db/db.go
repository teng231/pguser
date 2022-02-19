package db

import (
	"errors"
	"log"
	"sync"
	"time"

	"gitlab.com/my0sot1s/pguser/pb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type DB struct {
	engine *gorm.DB
}

// ConnectDb expose ...
func (d *DB) ConnectDb(sqlPath, dbName string) error {
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  sqlPath,
			PreferSimpleProtocol: true,
		}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			if err := sqlDb.Ping(); err != nil {
				log.Print(err)
			}
		}
	}()
	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDb.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDb.SetMaxOpenConns(100)
	d.engine = db
	return nil
}

func (d *DB) listUsersQuery(rq *pb.UserRequest) *gorm.DB {
	ss := d.engine.Table(tblUser)
	if rq.GetUsername() != "" {
		ss.Where("username = ?", rq.GetUsername())
	}
	if rq.GetFullname() != "" {
		ss.Where("fullname like ?", "%"+rq.GetFullname()+"%")
	}
	if rq.GetPhone() != "" {
		ss.Where("phone = ?", rq.GetPhone())
	}
	if rq.GetId() != "" {
		ss.Where("id = ?", rq.GetId())
	}
	if len(rq.GetIds()) != 0 {
		ss.Where("id IN ?", rq.GetIds())
	}
	if len(rq.GetNotIds()) != 0 {
		ss.Where("id NOT IN ?", rq.GetNotIds())
	}
	if rq.GetState() != "" {
		ss.Where("state = ?", rq.GetState())
	}
	return ss
}

// da check day la cau truc dung
func (d *DB) ListUsers(rq *pb.UserRequest) ([]*pb.User, error) {
	ss := d.listUsersQuery(rq)
	if rq.GetLimit() != 0 {
		ss.Limit(int(rq.GetLimit()))
	}
	if rq.GetSkip() != 0 {
		ss.Offset(int(rq.GetLimit()) * int(rq.GetSkip()))
	}
	users := make([]*pb.User, 0)
	err := ss.Order("id desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// da check day la cau truc dung
func (d *DB) ScanUserTable(cond *pb.User, buf chan *pb.User, wg *sync.WaitGroup) error {
	rows, err := d.engine.Table(tblUser).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	wg.Add(1)
	defer wg.Done()
	bean := new(pb.User)
	for rows.Next() {
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		err := d.engine.ScanRows(rows, bean)
		if err != nil {
			log.Print(err)
			continue
		}
		buf <- bean
		wg.Add(1)
	}
	log.Print("xxxx")
	return nil
}

// da check day la cau truc dung
func (d *DB) IsUserExisted(u *pb.User) bool {
	has := false
	d.engine.Table(tblUser).Select("count(id) > 0").Where(u).Take(&has)
	return has
}

// da check day la cau truc dung
func (d *DB) InsertUser(u *pb.User) error {
	tx := d.engine.Table(tblUser).Create(u)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no affected")
	}
	return nil
}

// da check day la cau truc dung
func (d *DB) UpdateUser(updator, selector *pb.User) error {

	tx := d.engine.Table(tblUser).Model(selector).Updates(updator)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no affected")
	}
	return nil
}

// da check day la cau truc dung
func (d *DB) ReadUser(u *pb.User) (*pb.User, error) {
	err := d.engine.Table(tblUser).Where(u).Take(u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("no affected")
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

// da check day la cau truc dung
func (d *DB) CountUsers(rq *pb.UserRequest) (int64, error) {
	ss := d.listUsersQuery(rq)
	var c int64
	if err := ss.Count(&c).Error; err != nil {
		return 0, err
	}
	return c, nil
}

// da check day la cau truc dung
func (d *DB) TransUserCreate(users ...*pb.User) (int64, error) {
	tx := d.engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	for _, item := range users {
		if err := tx.Table(tblUser).Create(item).Error; err != nil {
			log.Print(err)
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return 0, nil
}
