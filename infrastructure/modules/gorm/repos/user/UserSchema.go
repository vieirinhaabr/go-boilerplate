package gormUserRepo

import (
	"github.com/gofrs/uuid"
	"go-boilerplate/domain/user/entities"
	"gorm.io/gorm"
	"time"
)

type IRepo interface {
	Setup(conn *gorm.DB)
	Create(*userEntity.User) error
	Update(*userEntity.User) error
	GetById(id string) (userEntity.User, error)
}

func (Schema) TableName() string {
	return "users"
}

type Schema struct {
	ID        uuid.UUID `gorm:"<-:create;column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"column:email;type:varchar;not null"`
	Name      string    `gorm:"column:name;type:varchar;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null"`
}
