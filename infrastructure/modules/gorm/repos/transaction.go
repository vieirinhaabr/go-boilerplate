package transaction

import (
	"gorm.io/gorm"
)

type props struct {
	conn *gorm.DB
}

var Connection props

func Setup(conn *gorm.DB) {
	Connection = props{
		conn,
	}
}

func (p props) OpenTransaction() *gorm.DB {
	return p.conn.Begin()
}
