package inmemory

import (
	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

type serializableWorld struct {
	ID   string
	Name string
}

func (sw serializableWorld) toDomain() domain.World {
	return *new(domain.WorldBuilder).
		WithID(domain.WorldID(sw.ID)).
		WithName(sw.Name).
		Build()
}

type Worlds struct {
	Worlds []serializableWorld `yaml:"worlds"`
}
