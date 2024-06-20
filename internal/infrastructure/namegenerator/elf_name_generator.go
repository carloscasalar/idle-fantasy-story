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
	const namePattern = "<V|V|B|(Qu)<v>><V|C|C><V|C>(<v>|ü|ë|ä|û|î|ê|â|we|wi|wo|ue)('el|an|or|is|iel)"
	const surnamePattern = "BvCC(an|or|is|iel|ion|wen)"

	nameGenerator, err := fantasyname.Compile(namePattern, fantasyname.Collapse(true))
	if err != nil {
		return nil, fmt.Errorf("error compiling name generator: %w", err)
	}

	surnameGenerator, err := fantasyname.Compile(surnamePattern, fantasyname.Collapse(true))
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
