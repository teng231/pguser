package main

import (
	"log"
	"net"
	"sync"

	"gitlab.com/my0sot1s/pguser/db"
	"gitlab.com/my0sot1s/pguser/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type IDatabase interface {
	ReadUser(*pb.User) (*pb.User, error) // *pb.User
	IsUserExisted(u *pb.User) bool
	InsertUser(*pb.User) error
	UpdateUser(*pb.User, *pb.User) error
	ListUsers(rq *pb.UserRequest) ([]*pb.User, error)
	CountUsers(rq *pb.UserRequest) (int64, error)
	ScanUserTable(cond *pb.User, buf chan *pb.User, wg *sync.WaitGroup) error
	TransUserCreate(users ...*pb.User) (int64, error)
}

type User struct {
	Db IDatabase
}

func makeHandler(cf Configs) (*User, error) {
	d := &db.DB{}
	if err := d.ConnectDb(cf.DbPath, cf.DbName); err != nil {
		return nil, err
	}
	log.Print("Connect db successful")

	return &User{
		Db: d,
	}, nil
}

func GRPCServe(port string, handler *User) error {
	listen, err := net.Listen("tcp", ":"+cf.GrpcPort)
	if err != nil {
		return err
	}
	serve := grpc.NewServer()
	pb.RegisterUserServiceServer(serve, handler)
	reflection.Register(serve)
	return serve.Serve(listen)
}
