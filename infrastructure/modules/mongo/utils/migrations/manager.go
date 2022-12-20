package main

import (
	"context"
	"fmt"
	migrate "github.com/xakep666/mongo-migrate"
	"go-boilerplate/config/environment"
	mongoModule "go-boilerplate/infrastructure/modules/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	environment.SetupVars()
	config := environment.Vars.Mongo

	if len(os.Args) == 1 {
		panic("Missing arguments")
	}

	option := os.Args[1]

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
	if err != nil {
		panic(err)
	}

	database := client.Database(config.DBName)
	migrate.SetDatabase(database)
	migrate.SetMigrationsCollection(mongoModule.MigrationCollectionName)

	switch option {
	case "new":
		if len(os.Args) != 3 {
			panic("Should pass description of the migration")
		}

		_, filename, _, _ := runtime.Caller(0)
		dir := path.Dir(filename)

		fName := fmt.Sprintf("%s/../../migrations/%s_%s.go", dir, time.Now().Format("20060102150405"), os.Args[2])
		from, err := os.Open(fmt.Sprintf("%s/template.go", dir))
		if err != nil {
			panic(err)
		}
		defer from.Close()

		to, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 066)
		if err != nil {
			panic(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			panic(err)
		}

		log.Printf("New migration created: %s\n", fName)

	case "up":
		err = migrate.Up(migrate.AllAvailable)

	case "down":
		err = migrate.Down(migrate.AllAvailable)
	}

	if err != nil {
		panic(err)
	}
}
