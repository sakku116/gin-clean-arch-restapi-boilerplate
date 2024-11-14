package main

import (
	"context"
	"fmt"
	"os"
	"restapi/config"
)

func CliHandler(args []string) {
	args = args[1:]
	ctx := context.Background()

	mongoConn := config.NewMongoConn(ctx)
	defer mongoConn.Close(ctx)
	mongoDB := mongoConn.Database(config.Envs.MONGO_DB_NAME)
	_ = mongoDB

	switch args[0] {
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
