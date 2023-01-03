package gormTypes

import "github.com/jinzhu/gorm"

type Connection interface {
	Setup(conn *gorm.DB)
}
