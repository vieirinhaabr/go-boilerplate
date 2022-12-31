package redisUserRepo

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go-boilerplate/domain/user/entities"
	"go-boilerplate/infrastructure/modules/redis/utils"
	"go-boilerplate/utils"
)

type repo struct {
	client *redis.Client
	field  string
}

var Repo repo

func (r repo) Setup(client *redis.Client) {
	Repo = repo{client: client, field: "user"}
}

func (r repo) Set(user *userEntity.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	ctx, _ := utils.CreateContext()
	err = r.client.Set(ctx, user.Name, string(jsonUser), redisUtils.RedisDefaultTTL).Err()

	return err
}

func (r repo) Get(name string) (*userEntity.User, error) {
	ctx, _ := utils.CreateContext()

	data, err := r.client.Get(ctx, name).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if data == "" {
		return nil, fmt.Errorf("get: not found")
	}

	user := &userEntity.User{}
	err = utils.UnmarshalString(data, user)

	return user, err
}
