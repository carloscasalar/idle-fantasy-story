package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/carloscasalar/idle-fantasy-story/internal/application/generate"
	"os"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
)

// Just a simple example of how to use the name generator package
func main() {
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

	generateNames, err := generate.NewNames()
	if err != nil {
		println(fmt.Errorf("unable to instantiate name generator: %w", err))
		os.Exit(1)
	}
	names := generateNames.Execute(context.Background(), species, 10)
	for i, name := range names {
		println(fmt.Sprintf("Generated name (%d): %s", i+1, name))
	}

}
