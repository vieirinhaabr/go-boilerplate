package userEntity

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(Email string, Name string) User {
	return User{
		Email:     Email,
		Name:      Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (e User) Update(Email string, Name string) {
	e.Email = Email
	e.Name = Name
	e.UpdatedAt = time.Now()
}
