package namegenerator

import "github.com/carloscasalar/idle-fantasy-story/internal/domain"

type RandomNameGenerator struct {
}

func (r *RandomNameGenerator) GenerateCharacterName(_ domain.Species) string {
	return "a name"
}

func New() *RandomNameGenerator {
	return &RandomNameGenerator{}
}
