package world

import (
	"context"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"
	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

//go:generate mockery --name Repository --output mocks --outpkg mocks --case underscore
type Repository interface {
	GetWorlds(ctx context.Context) ([]domain.World, error)
}

type GetWorlds struct {
	worldRepository Repository
}

func NewGetWorlds(worldRepository Repository) *GetWorlds {
	return &GetWorlds{worldRepository: worldRepository}
}

func (g *GetWorlds) Execute(ctx context.Context) ([]WorldDTO, error) {
	worlds, err := g.worldRepository.GetWorlds(ctx)
	if err != nil {
		return nil, application.ErrInternalServer
	}
	return toDTOList(worlds), nil
}

func toDTOList(worlds []domain.World) []WorldDTO {
	result := make([]WorldDTO, len(worlds))
	for i, world := range worlds {
		result[i] = WorldDTO{
			ID:   string(world.ID()),
			Name: world.Name(),
		}
	}
	return result
}
