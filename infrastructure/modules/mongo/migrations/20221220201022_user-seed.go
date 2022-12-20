package migrations

import (
	"context"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type user struct {
	Email     string    `bson:"email"`
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

const collectionName = "users"

func init() {
	migrate.Register(func(db *mongo.Database) error {
		var newUsers = []interface{}{
			user{
				Email:     "teste",
				Name:      "teste@email.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			user{
				Email:     "teste2",
				Name:      "teste2@email.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		_, err := db.Collection(collectionName).InsertMany(context.TODO(), newUsers)
		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(collectionName).DeleteMany(context.TODO(), bson.D{{}})
		if err != nil {
			return err
		}
		return nil
	})
}
