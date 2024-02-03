package main

import (
	"fmt"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/namegenerator"
)

// Just a simple example of how to use the name generator package
func main() {
	generator, err := namegenerator.New()
	if err != nil {
		panic(fmt.Errorf("unable to instantiate name generator: %w", err))
	}

	for i := 0; i < 10; i++ {
		name := generator.GenerateCharacterName(domain.SpeciesHuman)
		println(fmt.Sprintf("Generated name (%d): %s", i+1, name))

	}
}
