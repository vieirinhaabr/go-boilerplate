package types

import "github.com/jinzhu/gorm"

type Connection interface {
	Setup(*gorm.DB)
}
