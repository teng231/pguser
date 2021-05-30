package db

import (
	"log"
	"sync"
	"testing"

	"gitlab.com/my0sot1s/pguser/pb"
)

func Test_connection(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb(
		"postgresql://postgres:123456@localhost:5432/pguser", "pguser",
	)
	if err != nil {
		log.Print(err)
	}
}

func Test_listUsers(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb(
		"postgresql://postgres:123456@localhost:5432/pguser", "pguser",
	)
	if err != nil {
		log.Print(err)
	}
	list, err := d.ListUsers(&pb.UserRequest{Limit: 5})
	log.Print(list, err)

	list, err = d.ListUsers(&pb.UserRequest{Limit: 5, Skip: 1})
	log.Print(list, err)
}

func Test_readUser(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb(
		"postgresql://postgres:123456@localhost:5432/pguser", "pguser",
	)
	if err != nil {
		log.Print(err)
	}
	list, err := d.ListUsers(&pb.UserRequest{Limit: 5})
	log.Print(list, err)

	list, err = d.ListUsers(&pb.UserRequest{Limit: 5, Skip: 1})
	log.Print(list, err)
}

func Test_readIterUser(t *testing.T) {
	d := &DB{}
	err := d.ConnectDb(
		"postgresql://postgres:123456@localhost:5432/pguser", "pguser",
	)
	if err != nil {
		log.Print(err)
	}
	wg := &sync.WaitGroup{}

	buf := make(chan *pb.User, 20)

	go func() {
		for {
			user := <-buf
			log.Print(user)
			wg.Done()
		}
	}()

	d.ScanUserTable(&pb.User{}, buf, wg)
	wg.Wait()
}
