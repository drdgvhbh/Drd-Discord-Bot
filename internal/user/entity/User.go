package entity

type User struct {
	id string
}

func (user User) ID() string {
	return user.id
}

func CreateUser(id string) *User {
	return &User{
		id: id,
	}
}
