package redisModule

import (
	"github.com/go-redis/redis/v9"
	"go-boilerplate/config/environment"
	redisUserRepo "go-boilerplate/infrastructure/modules/redis/repos/user"
	"go-boilerplate/utils"
	"log"
)

type Module struct {
	clients struct {
		user *redis.Client
		test *redis.Client
	}
}

func NewRedisModule() *Module {
	return &Module{}
}

func CreateClient(DB int) *redis.Client {
	var config = environment.Vars.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Pass,
		DB:       DB,
	})

	ctx, _ := utils.CreateContext()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}

func (handler *Module) Configure() {
	userClient := CreateClient(0)
	testClient := CreateClient(1)

	*handler = Module{clients: struct {
		user *redis.Client
		test *redis.Client
	}{user: userClient, test: testClient}}
}

func (handler *Module) Start() error {
	redisUserRepo.Repo.Setup(handler.clients.user)

	return nil
}

func (handler *Module) Finish() {
	_ = handler.clients.user.Close()
	_ = handler.clients.test.Close()

	log.Printf("[REDIS] Finished\n")
}
