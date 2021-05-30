package db

import (
	"log"

	"gitlab.com/my0sot1s/pguser/pb"
)

// CreateDb func
func (d *DB) MigrationDb() error {
	err := d.engine.Migrator().AutoMigrate(&pb.User{})
	if err != nil {
		log.Print(err)
	}
	err = d.engine.Migrator().AutoMigrate(&pb.Partner{})
	if err != nil {
		log.Print(err)
	}
	err = d.engine.Migrator().AutoMigrate(&pb.ProductType{})
	if err != nil {
		log.Print(err)
	}
	return nil
}
