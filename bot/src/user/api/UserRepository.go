package api

import (
	"Drd-Discord-Bot/bot/src/db/pg"
	"Drd-Discord-Bot/bot/src/user/entity"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	pq "github.com/lib/pq"
)

type UserRepository struct {
	db     *pg.Connector
	mapper UserMapper
}

var userRepositoryInstance *UserRepository

func CreateUserRepository(
	db *pg.Connector,
	mapper UserMapper,
) *UserRepository {
	if userRepositoryInstance != nil {
		return userRepositoryInstance
	}

	userRepositoryInstance = &UserRepository{db, mapper}
	return userRepositoryInstance
}

func (userRepository UserRepository) InsertUser(user *entity.User) {
	db := userRepository.db
	mapper := userRepository.mapper
	dbUser := mapper.MapFrom(user)

	query := fmt.Sprintf(`INSERT INTO users(id)
	VALUES 
		('%s');
	`, dbUser.ID)

	err := db.QueryRow(query).Scan()
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			panic(fmt.Sprintf(
				"Error: %s\nCode: %s\n%s\n",
				err.Error(),
				err.Code,
				spew.Sdump(user)))
		}
	}
}
