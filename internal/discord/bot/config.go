package bot

import (
	"os"
)

type Config struct {
	BotToken  string
	CmdPrefix string
	AppName   string
	EOFDelim  string
	DBPort    int
}

func CreateConfig() *Config {
	return &Config{
		BotToken:  os.Getenv("BOT_TOKEN"),
		CmdPrefix: os.Getenv("CMD_PREFIX"),
		AppName:   "Ruin and Grief Bot",
		EOFDelim:  "EOF",
	}
}

var config *Config

func ProvideConfig() *Config {
	if config != nil {
		return config
	}

	config = CreateConfig()

	return config
}
