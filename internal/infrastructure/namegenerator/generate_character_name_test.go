package namegenerator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/namegenerator"
	"github.com/stretchr/testify/assert"
)

func Test_RandomNameGenerator_GenerateCharacterName(t *testing.T) {
	testCases := map[string]struct {
		species domain.Species
	}{
		"should generate non empty name for humans":   {domain.SpeciesHuman},
		"should generate non empty name for elves":    {domain.SpeciesElf},
		"should generate non empty name for dwarfs":   {domain.SpeciesDwarf},
		"should generate non empty name for Halfings": {domain.SpeciesHalfing},
		"should generate non empty name for Kenders":  {domain.SpeciesKender},
		"should generate non empty name for Gnomes":   {domain.SpeciesGnome},
	}

	for assertion, tc := range testCases {
		t.Run(assertion, func(t *testing.T) {
			generator, err := namegenerator.New()
			require.NoError(t, err)

			name := generator.GenerateCharacterName(tc.species)

			assert.NotEmpty(t, name)
		})
	}
}
