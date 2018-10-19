package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"../anime/mal"
	messageMiddleware "../message/middleware"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func init() {
}

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	if (!messageMiddleware.IgnoreOwnMessagesMiddleware{}.ProcessMessage(session, m)) {
		return
	}
	if m.Content == "apes" {
		var userProfile = mal.GetProfile("Genshyguy")
		log.Println(userProfile)

		var favouriteAnimes string
		for _, animes := range userProfile.Favorites.Anime {
			favouriteAnimes += animes.Name
			favouriteAnimes += "\n"
		}
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x2E51A2, // Green
			Description: "My Anime List Profile",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Last Online",
					Value:  userProfile.LastOnline.Format("01/02/2006 15:04"),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Gender",
					Value:  userProfile.Gender,
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Birthday",
					Value:  userProfile.Birthday.Format("01/02/2006"),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Joined",
					Value:  userProfile.Joined.Format("01/02/2006"),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Favorite Animes",
					Value:  favouriteAnimes,
					Inline: false,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: userProfile.ImageURL,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: userProfile.ImageURL,
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     userProfile.Username,
		}
		session.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "Ruin" {
		session.ChannelMessageSend(m.ChannelID, "Grief")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "Grief" {
		session.ChannelMessageSend(m.ChannelID, "Ruin")
	}
}
