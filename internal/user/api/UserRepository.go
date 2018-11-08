package api

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	pq "github.com/lib/pq"
)

type UserRepository struct {
	db     pg.Connector
	mapper UserMapper
}

func CreateUserRepository(
	db pg.Connector,
	mapper UserMapper,
) *UserRepository {
	return &UserRepository{db, mapper}
}

var userRepositoryInstance *UserRepository

func ProvideUserRepository(
	db pg.Connector,
	mapper UserMapper,
) *UserRepository {
	if userRepositoryInstance != nil {
		return userRepositoryInstance
	}

	userRepositoryInstance = CreateUserRepository(db, mapper)
	return userRepositoryInstance
}

func (userRepository UserRepository) InsertUser(user *entity.User) error {
	db := userRepository.db
	mapper := userRepository.mapper
	dbUser := mapper.MapFrom(user)

	query := fmt.Sprintf(`INSERT INTO users(id, tokens)
	VALUES 
		('%s', %f);
	`, dbUser.ID, dbUser.Tokens)

	err := db.QueryRow(query).Scan()
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code.Name() {
			case "unique_violation":
				return &domain.UserAlreadyExistsError{}
			default:
				log.Fatalf("Error: %s\nCode: %s\n%s\n",
					err.Error(),
					err.Code,
					spew.Sdump(user))
			}
		}
	}

	return nil
}
