package inmemory

import (
	"context"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

type serializableWorld struct {
	ID      string
	Name    string
	Species []string
}

func (sw serializableWorld) toDomain(ctx context.Context) (*domain.World, error) {
	mappedSpecies, err := mapSpecies(ctx, sw.Species)
	if err != nil {
		return nil, err
	}
	return new(domain.WorldBuilder).
		WithID(domain.WorldID(sw.ID)).
		WithName(sw.Name).
		WithSpecies(mappedSpecies).
		Build(), nil
}

type Worlds struct {
	Worlds []serializableWorld `yaml:"worlds"`
}
