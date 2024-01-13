package domain

// WorldID represents the ID of a world
type WorldID string

// World represents a world in the game
type World struct {
	id   WorldID
	name string
}

func (w World) Name() string {
	return w.name
}

func (w World) ID() WorldID {
	return w.id
}

type WorldBuilder struct {
	id   *WorldID
	name *string
}

func (wb WorldBuilder) WithID(id WorldID) WorldBuilder {
	wb.id = &id
	return wb
}

func (wb WorldBuilder) WithName(name string) WorldBuilder {
	wb.name = &name
	return wb
}

func (wb WorldBuilder) Build() *World {
	return &World{
		id:   *wb.id,
		name: *wb.name,
	}
}
