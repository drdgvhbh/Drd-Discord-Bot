package cli

type Config struct {
	name        string
	description string
	helpPrefix  string
}

func (config Config) Name() string {
	return config.name
}

func (config Config) Description() string {
	return config.description
}

func (config Config) HelpPrefix() string {
	return config.helpPrefix
}

var configInstance *Config

func CreateConfig() *Config {
	return &Config{
		name:        "Ruin and Grief Bot",
		description: "General purpose bot for ruining and griefing",
		helpPrefix:  "!ruin",
	}
}

func ProvideConfig() *Config {
	if configInstance != nil {
		return configInstance
	}

	configInstance = CreateConfig()

	return configInstance
}
