package domain

// WorldID represents the ID of a world
type WorldID string

// World represents a world in the game
type World struct {
	id      WorldID
	name    string
	species []Species
}

func (w World) Name() string {
	return w.name
}

func (w World) ID() WorldID {
	return w.id
}

func (w World) Species() []Species {
	return w.species
}

type WorldBuilder struct {
	id      *WorldID
	name    *string
	species []Species
}

func (wb *WorldBuilder) WithID(id WorldID) *WorldBuilder {
	wb.id = &id
	return wb
}

func (wb *WorldBuilder) WithName(name string) *WorldBuilder {
	wb.name = &name
	return wb
}

func (wb *WorldBuilder) WithSpecies(species []Species) *WorldBuilder {
	wb.species = species
	return wb
}

func (wb *WorldBuilder) Build() *World {
	return &World{
		id:      *wb.id,
		name:    *wb.name,
		species: wb.species,
	}
}
