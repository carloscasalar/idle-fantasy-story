package namegenerator

import "github.com/carloscasalar/idle-fantasy-story/internal/domain"

type RandomNameGenerator struct {
	generator *genericNameGenerator
}

func (r *RandomNameGenerator) GenerateCharacterName(_ domain.Species) string {
	return r.generator.GenerateName()
}

func New() (*RandomNameGenerator, error) {
	generator, err := newGenericNameGenerator()
	if err != nil {
		return nil, err
	}
	return &RandomNameGenerator{
		generator: generator,
	}, nil
}
