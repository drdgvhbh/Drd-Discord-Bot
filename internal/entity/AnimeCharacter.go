package entity

type AnimeCharacter struct {
	id   string
	name string
}

func (character AnimeCharacter) ID() string {
	return character.id
}

func (character AnimeCharacter) Name() string {
	return character.name
}

func CreateCharacter(id string, name string) *AnimeCharacter {
	return &AnimeCharacter{
		id:   id,
		name: name,
	}
}
