package domain

type CharacterID string

type Character struct {
	id CharacterID
}

func (c *Character) ID() CharacterID {
	return c.id
}

type CharacterBuilder struct {
	id *CharacterID
}

func (cb *CharacterBuilder) WithID(id CharacterID) *CharacterBuilder {
	cb.id = &id
	return cb
}

func (cb *CharacterBuilder) Build() (*Character, error) {
	if cb.id == nil {
		return nil, NewUnexpectedError("id is required to build a character")
	}
	return &Character{
		id: *cb.id,
	}, nil
}
