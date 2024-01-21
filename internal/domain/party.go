package domain

type Party struct {
	characters []Character
}

func (p *Party) Characters() []Character {
	return p.characters
}

func NewParty(characters []Character) (*Party, error) {
	if len(characters) == 0 {
		return nil, NewUnexpectedError("characters are required to build a party")
	}
	return &Party{
		characters: characters,
	}, nil
}
