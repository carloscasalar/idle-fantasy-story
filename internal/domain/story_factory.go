package domain

import (
	"fmt"

	"github.com/carloscasalar/idle-fantasy-story/pkg/utils"
)

const defaultPartySize = 4

type StoryFactory struct {
	nameGenerator NameGenerator
	world         *World
	partySize     *uint8
}

func (sb *StoryFactory) WithNameGenerator(nameGenerator NameGenerator) *StoryFactory {
	sb.nameGenerator = nameGenerator
	return sb
}

func (sb *StoryFactory) WithWorld(world *World) *StoryFactory {
	sb.world = world
	return sb
}

func (sb *StoryFactory) WithPartySize(partySize *uint8) *StoryFactory {
	sb.partySize = partySize
	return sb
}

func (sb *StoryFactory) Build() (*Story, error) {
	if sb.world == nil {
		return nil, NewUnexpectedError("world is required to build a story")
	}
	partySize := utils.NoNilValue(sb.partySize, defaultPartySize)
	characters, err := sb.generateCharacters(partySize)
	if err != nil {
		return nil, err
	}
	party, err := NewParty(characters)
	if err != nil {
		return nil, err
	}
	return new(StoryBuilder).
		WithWorld(sb.world).
		WithParty(party).
		Build()
}

func (sb *StoryFactory) generateCharacters(size uint8) ([]Character, error) {
	characters := make([]Character, size)
	for i := range characters {
		id := CharacterID(fmt.Sprintf("character-%d", i))
		character, err := new(CharacterBuilder).
			WithID(id).
			WithName(sb.nameGenerator.GenerateCharacterName(SpeciesHuman)).
			Build()
		if err != nil {
			return nil, err
		}
		characters[i] = *character
	}
	return characters, nil
}
