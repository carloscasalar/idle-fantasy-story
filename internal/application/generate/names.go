package generate

import (
	"context"
	"fmt"
	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/namegenerator"
)

type Names struct {
	generator *namegenerator.RandomNameGenerator
}

func NewNames() (*Names, error) {
	generator, err := namegenerator.New()
	if err != nil {
		return nil, fmt.Errorf("error instantiating namegenerator: %w", err)
	}

	return &Names{
		generator: generator,
	}, nil

}

func (g Names) Execute(_ context.Context, species domain.Species, numberOfNamesToGenerate uint8) []string {
	names := make([]string, numberOfNamesToGenerate)
	for i := 0; i < int(numberOfNamesToGenerate); i++ {
		names[i] = g.generator.GenerateCharacterName(species)
	}
	return names
}
