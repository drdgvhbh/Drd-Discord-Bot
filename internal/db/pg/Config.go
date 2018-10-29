package pg

import "drdgvhbh/discordbot/internal/env"

type Config struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func (config Config) Host() string {
	return config.host
}

func (config Config) Port() string {
	return config.port
}

func (config Config) User() string {
	return config.user
}

func (config Config) Password() string {
	return config.password
}

func (config Config) Database() string {
	return config.database
}

func CreateConfig() *Config {
	return &Config{
		host:     env.Lookup("DB_HOST"),
		port:     env.Lookup("DB_PORT"),
		user:     env.Lookup("DB_USERNAME"),
		password: env.Lookup("DB_PASSWORD"),
		database: env.Lookup("DB_NAME"),
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
