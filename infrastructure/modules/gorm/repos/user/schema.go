package userRepo

import (
	user "go-boilerplate/domain/user/entities"
	"time"
)

type Repo interface {
	Create(user.Entity) error
	Update(user.Entity) error
}

func (Schema) TableName() string {
	return "users"
}

type Schema struct {
	ID        string    `gorm:"column:id;primary_key"`
	Email     string    `gorm:"column:email"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
