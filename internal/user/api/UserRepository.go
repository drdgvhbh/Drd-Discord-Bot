package api

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"
	"fmt"

	pq "github.com/lib/pq"
)

type UserRepository struct {
	databaseConnector      pg.Connector
	userDataTransferMapper UserDataTransferMapper
}

func CreateUserRepository(
	databaseConnector pg.Connector,
	userDataTransferMapper UserDataTransferMapper,
) *UserRepository {
	return &UserRepository{databaseConnector, userDataTransferMapper}
}

func (userRepository UserRepository) InsertUser(user *entity.User) error {
	databaseConnector := userRepository.databaseConnector
	userDataTransferMapper := userRepository.userDataTransferMapper

	userDTO := userDataTransferMapper.CreateDTOFrom(user)

	insertUserQuery := fmt.Sprintf(`INSERT INTO users(id, tokens)
		VALUES
			('%s', %f);
	`,
		userDTO.ID,
		userDTO.Tokens)

	insertUserError := databaseConnector.QueryRow(insertUserQuery).Scan()

	if insertUserError != nil {
		if insertUserPGError, isPGError := insertUserError.(*pq.Error); isPGError {
			switch insertUserPGError.Code.Name() {
			case "unique_violation":
				return domain.CreateDuplicateUserInsertionError(user)
			default:
				return insertUserPGError
			}
		}
	}

	return nil
}
