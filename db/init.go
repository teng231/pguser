package db

import (
	"log"

	"gitlab.com/my0sot1s/pguser/pb"
)

const (
	tblUser = "user"
)

// CreateDb func
func (d *DB) MigrationDb() error {
	err := d.engine.Table(tblUser).Migrator().AutoMigrate(&pb.User{})
	if err != nil {
		log.Print(err)
	}
	err = d.engine.Table("partner").Migrator().AutoMigrate(&pb.Partner{})
	if err != nil {
		log.Print(err)
	}
	err = d.engine.Table("product_type").Migrator().AutoMigrate(&pb.ProductType{})
	if err != nil {
		log.Print(err)
	}
	return nil
}
