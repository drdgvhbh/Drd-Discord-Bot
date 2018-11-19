package main

import (
	"drdgvhbh/discordbot/internal/cli/anime/mal/commands"
	"drdgvhbh/discordbot/internal/di"
	realCommands "drdgvhbh/discordbot/internal/discord/bot/commands"
	messageMiddleware "drdgvhbh/discordbot/internal/discord/message/middleware"
	"drdgvhbh/discordbot/internal/discord/writer"
	"drdgvhbh/discordbot/internal/entity"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
 ` + fmt.Sprintf("\n%s", os.Getenv("EOF_DELIM"))

	cli.CommandHelpTemplate = `NAME:
 {{.HelpName}} - {{.Usage}}

USAGE:
 {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

CATEGORY:
 {{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
 {{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:
 {{range .VisibleFlags}}{{.}}
 {{end}}{{end}}
` + fmt.Sprintf("\n%s", os.Getenv("EOF_DELIM"))
	cli.SubcommandHelpTemplate = `NAME:
 {{.HelpName}} - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

USAGE:
 {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
 {{.Name}}:{{end}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{end}}{{if .VisibleFlags}}
OPTIONS:
 {{range .VisibleFlags}}{{.}}
 {{end}}{{end}}
` + fmt.Sprintf("\n%s", os.Getenv("EOF_DELIM"))
}

func main() {
	dg := di.InitializeDiscordBot()

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	err := dg.Open()
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

	authorID := message.Author.ID
	channelID := message.ChannelID

	channel, _ := session.Channel(channelID)
	serverID := channel.GuildID

	userID := fmt.Sprintf("%s-%s", serverID, authorID)

	cmdPrefix := os.Getenv("CMD_PREFIX")
	if strings.HasPrefix(message.Content, cmdPrefix) {
		cmdStr := strings.TrimPrefix(message.Content, cmdPrefix)
		args := strings.Split(cmdStr, " ")
		cliApp := di.InitializeCLI()

		writer := writer.CreateDiscordWriter(
			session, message.ChannelID, os.Getenv("EOF_DELIM"))
		cliApp.Writer = writer

		cliApp.Action = func(c *cli.Context) error {
			return nil
		}
		cliApp.Commands = []cli.Command{}
		commandCB := func(result *discordgo.MessageEmbed) (*discordgo.Message, error) {
			return session.ChannelMessageSendEmbed(message.ChannelID, result)
		}

		sendChannelMessage := func(msg string) {
			session.ChannelMessageSend(
				channelID,
				msg)
		}

		sendChannelMessageWithAuthorMention := func(msg string) {
			sendChannelMessage(fmt.Sprintf(
				"%s %s", message.Author.Mention(), msg))
		}

		sendChannelMessageEmbed := func(msg *discordgo.MessageEmbed) {
			session.ChannelMessageSendEmbed(channelID, msg)
		}

		user := entity.NewUser(userID, 1000.0)

		cliApp.Commands = append(
			cliApp.Commands,
			*di.InitializeRegisterUserCommandFactory().Construct(
				user, sendChannelMessageWithAuthorMention,
			),
			*di.InitializeAnimeCommand([]cli.Command{
				*di.InitializeAnimeStockQuoteCommandFactory().Construct(
					sendChannelMessage,
					sendChannelMessageEmbed,
				),
				commands.GetAnimeProfileCommand(commandCB),
			}),
			*realCommands.NewAuctionCommand(
				[]cli.Command{
					*realCommands.NewCurrentAuctionCommand([]cli.Command{}),
					*realCommands.NewStartAuctionCommand([]cli.Command{}),
				},
			),
		)

		err := cliApp.Run(args)
		if err != nil {
			log.Println(err)
		}

	}
}
