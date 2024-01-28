package domain

type CharacterID string

type Character struct {
	id      CharacterID
	name    string
	species Species
}

func (c *Character) ID() CharacterID {
	return c.id
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) Species() Species {
	return c.species
}

type CharacterBuilder struct {
	id      *CharacterID
	name    *string
	species *Species
}

func (cb *CharacterBuilder) WithID(id CharacterID) *CharacterBuilder {
	cb.id = &id
	return cb
}

func (cb *CharacterBuilder) WithName(name string) *CharacterBuilder {
	cb.name = &name
	return cb
}

func (cb *CharacterBuilder) WithSpecies(species Species) *CharacterBuilder {
	cb.species = &species
	return cb
}

func (cb *CharacterBuilder) Build() (*Character, error) {
	if cb.id == nil {
		return nil, NewUnexpectedError("id is required to build a character")
	}
	if cb.name == nil {
		return nil, NewUnexpectedError("name is required to build a character")
	}
	if cb.species == nil {
		return nil, NewUnexpectedError("species is required to build a character")
	}
	return &Character{
		id:      *cb.id,
		name:    *cb.name,
		species: *cb.species,
	}, nil
}
