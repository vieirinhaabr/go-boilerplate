package repo

import (
	"gorm.io/gorm"
)

type transaction struct {
	conn *gorm.DB
}

var Transaction transaction

func (t *transaction) Setup(conn *gorm.DB) {
	*t = transaction{
		conn,
	}
}

func (t transaction) OpenTransaction() *gorm.DB {
	return t.conn.Begin()
}
