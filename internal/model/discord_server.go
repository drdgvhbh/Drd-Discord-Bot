package model

type DiscordServer struct {
	id string
}

func NewDiscordServer(id string) *DiscordServer {
	return &DiscordServer{
		id: id,
	}
}

func (discordServer *DiscordServer) ID() string {
	return discordServer.id
}
