package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/namegenerator"
)

// Just a simple example of how to use the name generator package
func main() {
	generator, err := namegenerator.New()
	if err != nil {
		panic(fmt.Errorf("unable to instantiate name generator: %w", err))
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please, choose a species to generate 10 names for it:")
	fmt.Println("1. Human")
	fmt.Println("2. Elf")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	var species domain.Species
	switch input {
	case "1\n":
		species = domain.SpeciesHuman
	case "2\n":
		species = domain.SpeciesElf
	default:
		fmt.Println("Invalid option")
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		name := generator.GenerateCharacterName(species)
		println(fmt.Sprintf("Generated name (%d): %s", i+1, name))

	}
}
