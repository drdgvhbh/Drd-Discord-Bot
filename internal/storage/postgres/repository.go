package postgres

import (
	"drdgvhbh/discordbot/internal/model"

	"github.com/google/uuid"

	"github.com/jackc/pgx"
)

type Auction = model.Auction

type Repository struct {
	dbConnection *pgx.Conn
}

func NewRepository(dbConnection *pgx.Conn) *Repository {
	return &Repository{dbConnection}
}

func (repository *Repository) AddAuction(auction *Auction) error {
	dbConnection := repository.dbConnection

	bidIds := []uuid.UUID{}
	for _, value := range auction.Bids() {
		bidIds = append(bidIds, value.ID())
	}

	_, err := dbConnection.Exec(`
	INSERT INTO auctions(
		id, ends_at, bids, anime_character_id, discord_server_id
	)
	VALUES (
		$1, $2, $3, $4, $5
	)`, auction.ID(),
		auction.EndsAt(),
		bidIds,
		auction.Item().ID(),
		auction.Location().ID())

	if err != nil {
		return err
	}

	return nil
}
