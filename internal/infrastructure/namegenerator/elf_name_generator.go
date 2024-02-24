package namegenerator

import (
	"fmt"

	"golang.org/x/text/language"

	"golang.org/x/text/cases"

	"github.com/s0rg/fantasyname"
)

type elfNameGenerator struct {
	upperCaser       *cases.Caser
	nameGenerator    fmt.Stringer
	surnameGenerator fmt.Stringer
}

func newElfNameGenerator() (*elfNameGenerator, error) {
	const namePattern = `
		(<B|v>)(a|e|i|o|u|ö|ü)(<B|v>)(a|e|i|o|u|ö|ü|l|n|r|t|s|h|d)(|(an|as|on|in|el|ain|ien|ien|ion|ath|orn|ith|ath|or|alar|ar|os|us|ath|iel|lar|lär|lir|lür|lur|lor|lär|lir|lür))(<v|B|C|>)
	`
	nameGenerator, err := fantasyname.Compile(namePattern, fantasyname.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("error compiling name generator: %w", err)
	}

	surnameGenerator, err := fantasyname.Compile(namePattern, fantasyname.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("error compiling surname generator: %w", err)
	}
	upperCaser := cases.Title(language.English)

	return &elfNameGenerator{
		upperCaser:       &upperCaser,
		nameGenerator:    nameGenerator,
		surnameGenerator: surnameGenerator,
	}, nil
}

func (g *elfNameGenerator) GenerateName() string {
	name := g.toUpperCase(g.nameGenerator.String())
	surname := g.toUpperCase(g.surnameGenerator.String())
	return fmt.Sprintf("%s %s", name, surname)
}

func (g *elfNameGenerator) toUpperCase(value string) string {
	return g.upperCaser.String(value)
}
