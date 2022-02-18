package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/caarlos0/env"
	"github.com/urfave/cli/v2"
	"gitlab.com/my0sot1s/pguser/db"
	"gitlab.com/my0sot1s/pguser/pb"
)

var cf Configs

type Configs struct {
	DbPath               string `env:"DB_PATH"`
	DbName               string `env:"DB_NAME"`
	HttpPort             string `env:"HTTP_PORT"`
	ServiceName          string `env:"SERVICE_NAME"`
	SyncerServerUrl      string `env:"SYNCER_SERVER_URL"`
	RedisAddr            string `env:"REDIS_ADDR"`
	RedisPw              string `env:"REDIS_PW"`
	KafkaBrokers         string `env:"KAFKA_BROKERS"`
	PubsubSubscription   string `env:"PUBSUB_SUBSCRIPTION"`
	AccessTokenDuration  int    `env:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration int    `env:"REFRESH_TOKEN_DURATION"`
	RsaPrivatePath       string `env:"RSA_PRIVATE_PATH"`
	RsaPublicPath        string `env:"RSA_PUBLIC_PATH"`
	Secret               string `env:"SECRET"`
	GrpcPort             string `env:"GRPC_PORT"`
	PermGrpcServer       string `env:"PERM_GRPC_SERVER"`
}

func migrationDb(ctx *cli.Context) error {
	d := &db.DB{}
	if err := d.ConnectDb(cf.DbPath, cf.DbName); err != nil {
		debug.PrintStack()
		return err
	}
	if err := d.MigrationDb(); err != nil {
		log.Print(err)
	}
	log.Print("Tables migration done!")
	return nil
}

func startApp(ctx *cli.Context) error {
	log.Print("STart app")
	h, err := makeHandler(cf)
	if err != nil {
		panic(err)
	}
	// data, err := h.Db.ReadUser(&pb.User{Id: "usrbqj4l1920hsad3tp6n17g"})
	// log.Print(data, err)

	// u := &pb.User{Fullname: "teng1"}
	// err = h.Db.UpdateUser(u, &pb.User{Id: "usrbru30a6lss08bj5adn5g1"})
	// log.Print(err, u)

	// u := &pb.User{Id: "usrbxxx"}
	// err = h.Db.InsertUser(u)
	// log.Print(err, u)

	// log.Print(h.Db.IsUserExisted(&pb.User{Id: "usrbxxx2"}))

	// users := make([]*pb.User, 0)
	// users, err = h.Db.ListUsers(&pb.UserRequest{Limit: 2, Username: "0983369399", Ids: []string{"usrbcg6tsp20hsad3si9bm0"}})
	// log.Print(err, users)

	// data, err := h.Db.CountUsers(&pb.UserRequest{Id: "usrblk8snh20hsad3su4eo0"})
	// log.Print(data, err)

	// users := make([]*pb.User, 0)
	// users, err = h.Db.ListUsers(&pb.UserRequest{Limit: 2, Username: "0983369399", Ids: []string{"usrbcg6tsp20hsad3si9bm0"}})
	// log.Print(err, users)

	// data := make(chan *pb.User, 100)
	// wg := &sync.WaitGroup{}
	// go func() {
	// 	for {
	// 		x := <-data
	// 		log.Print(x)
	// 		wg.Done()
	// 	}
	// }()
	// h.Db.ScanUserTable(&pb.User{}, data, wg)
	// wg.Wait()
	// log.Print(h.Db.TransUserCreate(&pb.User{Id: "1", Username: "1"}, &pb.User{Id: "2", Username: "2"}))
	log.Print(h.Db.TransUserCreate(&pb.User{Id: "1", Username: "1"}, &pb.User{Id: "2", Username: "2"}))
	if err := GRPCServe(cf.GrpcPort, h); err != nil {
		return err
	}
	return nil
}

func appRoot() error {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		return errors.New("Wow, ^.^ dumb")
	}
	app.Commands = []*cli.Command{
		{Name: "start", Usage: "start up running app", Action: startApp},
		{Name: "migrationDb", Usage: "create table in db with proto", Action: migrationDb},
	}

	return app.Run(os.Args)
}

func main() {
	log.Print("$ STARTED")
	if err := env.Parse(&cf); err != nil {
		log.Fatal(err)
	}
	go freeMemory()
	if err := appRoot(); err != nil {
		panic(err)
	}
}

func freeMemory() {
	for {
		fmt.Println("run gc")
		start := time.Now()
		runtime.GC()
		debug.FreeOSMemory()
		elapsed := time.Since(start)
		fmt.Printf("gc took %s\n", elapsed)
		time.Sleep(15 * time.Minute)
	}
}
