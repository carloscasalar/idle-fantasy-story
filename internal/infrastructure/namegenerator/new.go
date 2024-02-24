package namegenerator

import "github.com/carloscasalar/idle-fantasy-story/internal/domain"

type NameGenerator interface {
	GenerateName() string
}

type RandomNameGenerator struct {
	generatorBySpecies map[domain.Species]NameGenerator
	defaultGenerator   *genericNameGenerator
}

func (r *RandomNameGenerator) GenerateCharacterName(species domain.Species) string {
	generator, found := r.generatorBySpecies[species]
	if !found {
		generator = r.defaultGenerator
	}

	return generator.GenerateName()
}

func New() (*RandomNameGenerator, error) {
	generator, err := newGenericNameGenerator()
	if err != nil {
		return nil, err
	}

	elfGenerator, err := newElfNameGenerator()
	if err != nil {
		return nil, err
	}

	generatorBySpecies := map[domain.Species]NameGenerator{
		domain.SpeciesElf: elfGenerator,
	}

	return &RandomNameGenerator{
		generatorBySpecies: generatorBySpecies,
		defaultGenerator:   generator,
	}, nil
}
