package user

import "time"

type Entity struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Create(Email string, Name string) Entity {
	return Entity{
		Email:     Email,
		Name:      Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (e Entity) Update(Email string, Name string) {
	e.Email = Email
	e.Name = Name
	e.UpdatedAt = time.Now()
}
