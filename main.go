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
	user, err := makeHandler(cf)
	if err != nil {
		panic(err)
	}
	if err := GRPCServe(cf.GrpcPort, user); err != nil {
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
		{Name: "createDb", Usage: "create table in db with proto", Action: migrationDb},
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
