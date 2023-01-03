package grpcModule

import (
	"go-boilerplate/config/environment"
	grpcClient "go-boilerplate/infrastructure/modules/grpc/clients"
	"go-boilerplate/infrastructure/modules/grpc/servers"
	"go-boilerplate/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type Module struct {
	// Client
	boilerplateConn *grpc.ClientConn

	// Server
	server   *grpc.Server
	listener net.Listener
}

func NewGrpcModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	utils.Log.Warn("[GRPC] ᛃ Configuring\n")

	var config = environment.Vars.Grpc
	var services = environment.Vars.Services

	// Client
	boilerplateConn, err := grpc.Dial(services.BoilerplateHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("[GRPC] Listen: " + err.Error())
	}

	// Server
	server := grpc.NewServer()
	grpcServer.UserServer.ConfigureServer(server)

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		panic("[GRPC] Listen: " + err.Error())
	}

	*handler = Module{boilerplateConn, server, lis}
}

func (handler *Module) Start() error {
	utils.Log.Warn("[GRPC] ▶ Starting\n")

	var config = environment.Vars.Grpc

	// Client
	grpcClient.ConfigureUserClient(handler.boilerplateConn)

	// Server
	utils.Log.Warn("[GRPC] Listening on %s\n", config.Port)
	return handler.server.Serve(handler.listener)
}

func (handler *Module) Finish() {
	utils.Log.Warn("[GRPC] ■ Finishing\n")

	// Server
	handler.server.Stop()
}
