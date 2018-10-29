package middleware

import (
	"github.com/bwmarrin/discordgo"
)

type MessageMiddleware interface {
	ProcessMessage(
		session *discordgo.Session,
		messageCreateEvent *discordgo.MessageCreate) bool
}
