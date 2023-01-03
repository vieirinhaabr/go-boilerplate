package mongoUserRepo

import (
	"errors"
	"go-boilerplate/domain/user/entities"
	errorsType "go-boilerplate/infrastructure/global/errors"
	"go-boilerplate/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	database   *mongo.Database
	collection *mongo.Collection
}

var Repo repo

func (r repo) Setup(db *mongo.Database) {
	collection := db.Collection(CollectionName)
	Repo = repo{db, collection}
}

func (r repo) Create(user *userEntity.User) error {
	ctx, cancel := utils.CreateContext()
	defer cancel()

	schema := fromEntity(user)
	result, err := r.collection.InsertOne(ctx, schema)
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return err
}

func (r repo) Update(user *userEntity.User) error {
	ctx, cancel := utils.CreateContext()
	defer cancel()

	schema := fromEntity(user)
	objId, _ := primitive.ObjectIDFromHex(user.ID)
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", schema}}
	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r repo) GetById(id string) (userEntity.User, error) {
	ctx, cancel := utils.CreateContext()
	defer cancel()

	var schema = new(Schema)
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	result := r.collection.FindOne(ctx, filter)
	_ = result.Decode(schema)
	user := toEntity(schema)

	if result.Err() == mongo.ErrNoDocuments {
		return *user, errors.New(string(errorsType.NoItemsFound))
	}

	return *user, result.Err()
}
