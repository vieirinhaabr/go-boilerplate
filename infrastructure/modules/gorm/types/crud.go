package types

type Crud interface {
	Create() (bool, error)
	Update() (bool, error)
	Delete() (bool, error)
}
