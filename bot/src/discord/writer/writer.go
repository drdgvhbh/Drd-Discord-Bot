package writer

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type discordWriter struct {
	strBuilder *strings.Builder
	session    *discordgo.Session
	channelID  string
	eofDelim   string
}

func CreateDiscordWriter(
	session *discordgo.Session,
	channelID string,
	eofDelim string,
) io.Writer {
	discordWriter := new(discordWriter)
	discordWriter.strBuilder = new(strings.Builder)
	discordWriter.eofDelim = eofDelim
	discordWriter.channelID = channelID
	discordWriter.session = session

	return discordWriter
}

func (writer discordWriter) Write(p []byte) (numBytesWritten int, err error) {
	EOF := writer.eofDelim
	if binary.Size(p) == len(EOF) && string(p[:]) == EOF {
		writer.session.ChannelMessageSend(
			writer.channelID, fmt.Sprintf("```%s```", writer.strBuilder.String()))
		writer.strBuilder.Reset()
	} else {
		bytes, err := writer.strBuilder.Write(p)
		return bytes, err
	}
	return 0, io.EOF
}
