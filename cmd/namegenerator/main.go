package main

import (
	"fmt"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/namegenerator"
)

// Just a simple example of how to use the name generator package
func main() {
	generator := namegenerator.New()

	name := generator.GenerateCharacterName(domain.SpeciesHuman)

	println(fmt.Sprintf("Generated name: %s", name))
}
