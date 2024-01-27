package inmemory

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

type Repository struct {
	store *sync.Map
}

type Opt func(*Repository) error

func NewRepository(ctx context.Context, worldsFilePath string) (*Repository, error) {
	worlds, err := newWorldsLoader(worldsFilePath).load(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("error initializing worlds")
		return nil, ErrUnableToParseYmlWorlds
	}
	store := &sync.Map{}
	store.Store("worlds", worlds)
	return &Repository{
		store: store,
	}, nil
}

func (r *Repository) GetWorlds(context.Context) ([]domain.World, error) {
	worldsInStore, ok := r.store.Load("worlds")
	if !ok {
		return nil, ErrUnableToRetrieveWorlds
	}
	worldsByID := worldsInStore.(map[string]domain.World)
	worlds := make([]domain.World, len(worldsByID))
	i := 0
	for _, world := range worldsByID {
		worlds[i] = world
		i++
	}

	return worlds, nil
}
