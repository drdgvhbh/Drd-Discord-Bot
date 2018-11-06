package entity

type Character struct {
	id   string
	name string
}

func (character Character) ID() string {
	return character.id
}

func (character Character) Name() string {
	return character.Name()
}

func CreateCharacter(id string, name string) *Character {
	return &Character{
		id:   id,
		name: name,
	}
}
