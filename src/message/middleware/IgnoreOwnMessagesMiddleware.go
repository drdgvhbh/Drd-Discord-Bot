package middleware

import (
	"github.com/bwmarrin/discordgo"
)

type IgnoreOwnMessagesMiddleware struct{}

func (middleware IgnoreOwnMessagesMiddleware) ProcessMessage(
	session *discordgo.Session,
	messageCreateEvent *discordgo.MessageCreate) bool {
	return !(messageCreateEvent.Author.ID == session.State.User.ID)
}
