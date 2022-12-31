package types

type Module interface {
	Configure()
	Start() error
	Finish()
}
