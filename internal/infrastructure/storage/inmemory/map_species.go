package inmemory

import (
	"context"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	log "github.com/sirupsen/logrus"
)

func mapSpecies(ctx context.Context, species []string) ([]domain.Species, error) {
	speciesByStringID := map[string]domain.Species{
		"human":   domain.SpeciesHuman,
		"elf":     domain.SpeciesElf,
		"dwarf":   domain.SpeciesDwarf,
		"halfing": domain.SpeciesHalfing,
		"kender":  domain.SpeciesKender,
		"gnome":   domain.SpeciesGnome,
	}
	var result []domain.Species
	for _, s := range species {
		speciesFound, found := speciesByStringID[s]
		if !found {
			log.WithContext(ctx).
				WithField("species", species).
				Errorf("unknown species '%v'", s)
			return nil, ErrInvalidSpecies
		}
		result = append(result, speciesFound)
	}
	return result, nil
}
