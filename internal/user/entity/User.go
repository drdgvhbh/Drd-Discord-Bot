package entity

type User struct {
	id     string
	tokens float64
}

func (user User) ID() string {
	return user.id
}

func (user User) Tokens() float64 {
	return user.tokens
}

func CreateUser(id string, startingTokens float64) *User {
	return &User{
		id:     id,
		tokens: startingTokens,
	}
}
