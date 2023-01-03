package redisModule

import (
	"github.com/go-redis/redis/v9"
	"go-boilerplate/config/environment"
	redisUserRepo "go-boilerplate/infrastructure/modules/redis/repos/user"
	"go-boilerplate/utils"
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

	ctx, cancel := utils.CreateContext()
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}

func (handler *Module) Configure() {
	utils.Log.Warn("[REDIS] ᛃ Configuring\n")

	userClient := CreateClient(0)
	testClient := CreateClient(1)

	*handler = Module{clients: struct {
		user *redis.Client
		test *redis.Client
	}{user: userClient, test: testClient}}
}

func (handler *Module) Start() error {
	utils.Log.Warn("[REDIS] ▶ Starting\n")

	redisUserRepo.Repo.Setup(handler.clients.user)

	return nil
}

func (handler *Module) Finish() {
	utils.Log.Warn("[REDIS] ■ Finishing\n")

	_ = handler.clients.user.Close()
	_ = handler.clients.test.Close()
}
