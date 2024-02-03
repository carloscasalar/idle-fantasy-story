package namegenerator

import (
	"fmt"

	"golang.org/x/text/language"

	"golang.org/x/text/cases"

	"github.com/s0rg/fantasyname"
)

type genericNameGenerator struct {
	upperCaser       *cases.Caser
	nameGenerator    fmt.Stringer
	surnameGenerator fmt.Stringer
}

func newGenericNameGenerator() (*genericNameGenerator, error) {
	const namePattern = "(bil|bal|ban|hil|ham|hal|hol|hob|wil|me|or|ol|od|gor|for|fos|tol|ar|fin|ere|leo|vi|bi|bren|thor)" +
		"(|go|orbis|apol|adur|mos|ri|i|na|ole|n)" +
		"(|tur|axia|and|bo|gil|bin|bras|las|mac|grim|wise|l|lo|fo|co|ra|via|" +
		"da|ne|ta|y|wen|thiel|phin|dir|dor|tor|rod|on|rdo|dis)"
	nameGenerator, err := fantasyname.Compile(namePattern, fantasyname.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("error compiling name generator: %w", err)
	}

	const surnamePattern = "(Dragon|Wolf|Bear|Lion|Eagle|Shadow|Fire|Ice|Stone|Storm|Wind|River|Lake|Sky|Light|Night|Gold|Silver|Iron|Bronze|Sword|Shield|Bow|Spear|Horn|Thorn|Crown|Cloak|Star|Moon|Sun|Forest|Mountain|Valley|Sea|Flame|Frost|Dusk|Dawn)(heart|fang|claw|song|breeze|runner|walker|singer|weaver|smith|reaper|watcher|warden|whisper|rider|feather|spear|blade|eye|spell|gaze|bane|caller|cloak|dancer|maul|breaker|cutter|slayer|stalker|wielder)"
	surnameGenerator, err := fantasyname.Compile(surnamePattern, fantasyname.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("error compiling surname generator: %w", err)
	}
	upperCaser := cases.Title(language.English)

	return &genericNameGenerator{
		upperCaser:       &upperCaser,
		nameGenerator:    nameGenerator,
		surnameGenerator: surnameGenerator,
	}, nil
}

func (g *genericNameGenerator) GenerateName() string {
	name := g.toUpperCase(g.nameGenerator.String())
	surname := g.toUpperCase(g.surnameGenerator.String())
	return fmt.Sprintf("%s %s", name, surname)
}

func (g *genericNameGenerator) toUpperCase(value string) string {
	return g.upperCaser.String(value)
}
