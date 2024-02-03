package namegenerator

import "github.com/carloscasalar/idle-fantasy-story/internal/domain"

type RandomNameGenerator struct {
}

func (r *RandomNameGenerator) GenerateCharacterName(_ domain.Species) string {
	name, err := generateGenericName()
	if err != nil {
		return err.Error()
	}

	return name
}

func New() *RandomNameGenerator {
	return &RandomNameGenerator{}
}
