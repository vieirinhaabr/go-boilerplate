package grpcModule

type Module struct {
}

func NewGrpcModule() *Module {
	return &Module{}
}

func (handler *Module) Configure() {}

func (handler *Module) Start() {}

func (handler *Module) Finish() {}
