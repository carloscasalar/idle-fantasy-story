package domain

import (
	"fmt"

	"github.com/carloscasalar/idle-fantasy-story/pkg/utils"
)

const defaultPartySize = 4

type StoryFactory struct {
	world     *World
	partySize *uint8
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
	characters, err := generateCharacters(partySize)
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

func generateCharacters(size uint8) ([]Character, error) {
	characters := make([]Character, size)
	for i := range characters {
		id := CharacterID(fmt.Sprintf("character-%d", i))
		character, err := new(CharacterBuilder).
			WithID(id).
			WithName(fmt.Sprintf("Character %d", i)).
			Build()
		if err != nil {
			return nil, err
		}
		characters[i] = *character
	}
	return characters, nil
}
