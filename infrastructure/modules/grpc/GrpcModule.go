package grpcModule

import (
	"go-boilerplate/config/environment"
	"go-boilerplate/infrastructure/modules/grpc/servers"
	"go-boilerplate/utils"
	"google.golang.org/grpc"
	"net"
)

type Module struct {
	// Server
	server   *grpc.Server
	listener net.Listener
}

func NewGrpcModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {
	var config = environment.Vars.Grpc

	// Server
	server := grpc.NewServer()
	grpcServer.UserServer.ConfigureServer(server)

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		panic("[GRPC] Listen: " + err.Error())
	}

	*handler = Module{server, lis}
}

func (handler *Module) Start() error {
	var config = environment.Vars.Grpc

	// Server
	utils.Log.Info("[GRPC] Listening on %s\n", config.Port)
	return handler.server.Serve(handler.listener)
}

func (handler *Module) Finish() {
	// Server
	handler.server.Stop()

	utils.Log.Error("[GRPC] Finished\n")
}
