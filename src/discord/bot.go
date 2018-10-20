package main

import (
	botcli "cli"
	"discord/anime/mal"
	messageMiddleware "discord/message/middleware"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

func init() {
	cli.AppHelpTemplate = `NAME:
	{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

 USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

 VERSION:
	{{.Version}}{{end}}{{end}}{{if .Description}}

 DESCRIPTION:
	{{.Description}}{{end}}{{if len .Authors}}

 AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
	{{range $index, $author := .Authors}}{{if $index}}
	{{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

 COMMANDS:{{range .VisibleCategories}}{{if .Name}}

	{{.Name}}:{{end}}{{range .VisibleCommands}}
	  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

 GLOBAL OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

 COPYRIGHT:
	{{.Copyright}}{{end}}
EOF
 `
}

func main() {
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
func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if (!messageMiddleware.IgnoreOwnMessagesMiddleware{}.ProcessMessage(session, message)) {
		return
	}
	if message.Content == "apes" {
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
		session.ChannelMessageSendEmbed(message.ChannelID, embed)
	}

	// If the message is "ping" reply with "Pong!"
	if message.Content == "Ruin" {
		session.ChannelMessageSend(message.ChannelID, "Grief")
	}

	// If the message is "pong" reply with "Ping!"
	if message.Content == "Grief" {
		session.ChannelMessageSend(message.ChannelID, "Ruin")
	}

	prefix := "!ruin"
	if strings.HasPrefix(message.Content, prefix) {
		cmdStr := strings.TrimPrefix(message.Content, prefix)
		args := strings.Split(cmdStr, " ")
		cliApp := botcli.CreateCLI()

		var sb strings.Builder
		writer := customWriter{&sb, session, message.ChannelID}
		cliApp.Writer = writer

		err := cliApp.Run(args)
		if err != nil {
			log.Fatal(err)
		}

	}
}

type customWriter struct {
	strBuilder *strings.Builder
	session    *discordgo.Session
	channelID  string
}

func (w customWriter) Write(p []byte) (n int, err error) {
	print(string(p[:]))

	if binary.Size(p) == 3 && string(p[:]) == "EOF" {
		w.session.ChannelMessageSend(w.channelID, fmt.Sprintf("```%s```", w.strBuilder.String()))
		w.strBuilder.Reset()
	} else {
		bytes, err := w.strBuilder.Write(p)
		return bytes, err
	}
	return 0, io.EOF
}
