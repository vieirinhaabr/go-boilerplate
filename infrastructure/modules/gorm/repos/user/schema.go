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
	ID        string    `gorm:"column:id;primaryKey"`
	Email     string    `gorm:"column:email;not null"`
	Name      string    `gorm:"column:name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}
