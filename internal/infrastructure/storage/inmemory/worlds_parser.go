package inmemory

import (
	"context"
	"fmt"
	"os"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"gopkg.in/yaml.v3"
)

type worldsLoader struct {
	worldsFilePath string
}

func newWorldsLoader(worldsFilePath string) *worldsLoader {
	return &worldsLoader{worldsFilePath: worldsFilePath}
}

func (wl *worldsLoader) load(ctx context.Context) (map[string]domain.World, error) {
	rawWorlds, err := wl.parseYmlWorlds()
	if err != nil {
		return nil, fmt.Errorf("error parsing yml worlds: %w", err)
	}
	worlds := make(map[string]domain.World, len(rawWorlds))
	for _, world := range rawWorlds {
		mappedWorld, err := world.toDomain(ctx)
		if err != nil {
			return nil, err
		}
		worlds[world.ID] = *mappedWorld
	}
	return worlds, nil
}

func (wl *worldsLoader) parseYmlWorlds() ([]serializableWorld, error) {
	file, err := os.ReadFile(wl.worldsFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading worlds file: %w", err)
	}

	worlds := Worlds{}
	err = yaml.Unmarshal(file, &worlds)
	return worlds.Worlds, nil
}
