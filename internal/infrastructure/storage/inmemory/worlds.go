package inmemory

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

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

func initWorlds() (map[string]domain.World, error) {
	rawWorlds, err := parseYmlWorlds()
	if err != nil {
		return nil, fmt.Errorf("error parsing yml worlds: %w", err)
	}
	worlds := make(map[string]domain.World, len(rawWorlds))
	for _, world := range rawWorlds {
		worlds[world.ID] = world.toDomain()
	}
	return worlds, nil
}

func parseYmlWorlds() ([]serializableWorld, error) {
	//TODO move to an assets folder and parametrize the route
	file, err := os.ReadFile("internal/infrastructure/storage/inmemory/worlds.yml")
	if err != nil {
		return nil, fmt.Errorf("error reading worlds file: %w", err)
	}

	worlds := Worlds{}
	err = yaml.Unmarshal(file, &worlds)
	return worlds.Worlds, nil
}
