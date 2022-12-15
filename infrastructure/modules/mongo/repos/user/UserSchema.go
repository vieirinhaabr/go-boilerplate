package mongoUserRepo

import (
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/mongo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IRepo interface {
	Setup(db *mongo.Database)
	Create(*userEntity.User) error
	Update(*userEntity.User) error
	GetById(id string) (userEntity.User, error)
}

func fromEntity(user *userEntity.User) *Schema {
	return utils.ConvertSchema[Schema](user)
}

func toEntity(schema *Schema) *userEntity.User {
	return utils.ConvertSchema[userEntity.User](schema)
}

const CollectionName = "users"

type Schema struct {
	ID  		primitive.ObjectID 	`bson:"_id,omitempty" json:"id,omitempty"`
	Email     	string    			`bson:"email" json:"email"`
	Name      	string    			`bson:"name" json:"name"`
	CreatedAt 	time.Time 			`bson:"created_at" json:"created_at"`
	UpdatedAt 	time.Time 			`bson:"updated_at" json:"updated_at"`
}
