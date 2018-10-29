package api

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/user/entity"
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

func (userRepository UserRepository) InsertUser(user *entity.User) error {
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
			return fmt.Errorf(
				"error: %s\nCode: %s\n%s",
				err.Error(),
				err.Code,
				spew.Sdump(user))
		}
	}

	return nil
}
