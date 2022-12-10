package user

func (User) TableName() string {
	return "users"
}

type User struct {
	id    string `gorm:"primary_key"`
	Email string
	Name  string
}
