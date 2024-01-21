package domain

type CharacterID string

type Character struct {
	id   CharacterID
	name string
}

func (c *Character) ID() CharacterID {
	return c.id
}

func (c *Character) Name() string {
	return c.name
}

type CharacterBuilder struct {
	id   *CharacterID
	name *string
}

func (cb *CharacterBuilder) WithID(id CharacterID) *CharacterBuilder {
	cb.id = &id
	return cb
}

func (cb *CharacterBuilder) WithName(name string) *CharacterBuilder {
	cb.name = &name
	return cb
}

func (cb *CharacterBuilder) Build() (*Character, error) {
	if cb.id == nil {
		return nil, NewUnexpectedError("id is required to build a character")
	}
	return &Character{
		id:   *cb.id,
		name: *cb.name,
	}, nil
}
