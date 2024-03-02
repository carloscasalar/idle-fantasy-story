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
	const namePattern = "(sil|sal|san|sil|sam|sal|sol|sob|sil|me|or|ol|od|gor|for|fos|tol|ar|fin|ere|leo|vi|bi|bren|thor|theo|thai|eow|aow|ewi)" +
	"(|go|orbis|apol|adur|mos|ri|i|na|ole|n)" +
	"(|tur|axia|and|bo|gil|bin|bras|las|mac|grim|wise|l|lo|fo|co|ra|via|" +
	"da|ne|ta|y|wen|thiel|phin|dir|dor|tor|rod|on|rdo|dis)"

	const titles = "(Speaker|Guardian|Keeper|Warden|Defender|Watcher|Sentinel|Champion|Master|Lord|Lady|Prince|Princess|King|Queen|Emperor|Empress|High|Grand|Arch|Elder|Ancient|Wise|Sage|Great|Noble|Pure|Bright|Golden|Silver|Iron|Bronze|Steel|Copper|Platinum|Diamond|Ruby|Emerald|Sapphire|Amethyst|Topaz|Opal|Pearl|Jade|Crystal|Quartz|Obsidian|Onyx|Amber|Ivory|Ebony|Mahogany|Crimson|Azure|Cerulean|Cobalt|Indigo|Violet|Lilac|Lavender|Magenta|Coral|Cyan|Teal|Turquoise|Aquamarine|Peridot|Malachite|Lapis|Agate|Garnet|Jasper|Moonstone|Sunstone|Starstone|Bloodstone|Firestone|Icestone|Stormstone|Windstone|Earthstone|Waterstone|Air|Earth|Fire|Water|Ice|Storm|Wind|Light|Dark|Night|Day|Dawn|Dusk|Sun|Moon|Star|Sky|Cloud|Rain|Snow|Frost|Hail|Thunder|Lightning|Aurora|Rainbow|Flame|Blaze|Inferno|Pyre|Furnace|Forge|Anvil|Hammer|Sword|Axe|Spear|Bow|Arrow|Shield|Armor|Helm|Crown|Cloak|Robe|Gown|Dress|Tunic|Vest|Shirt|Pants|Boots|Gloves|Gauntlets|Bracers|Belt|Cape|Pendant|Necklace|Ring|Bracelet|Earring|Brooch|Tiara|Circlet|Diadem|Crown|Scepter|Staff|Wand|Rod|Tome|Book|Scroll|Parchment|Ink|Quill|Pen|Sword|Blade|Dagger|Knife|Axe|Hammer|Mace|Flail|Morningstar|Spear|Lance|Halberd|Pike|Bow|Crossbow|Sling)"
	const articles = "( of the | of |'s |' |, the |, | and the | and | or the | or | nor the | nor | but the | but | yet the | yet | so the | so | for the | for | to the | to | in the | in | on the | on | at the | at | by the | by | with the | with | without the | without | within the | within | among the | among | between the | between | through the | through | over the | over | under the | under | behind the | behind | in front of the | in front of | near the | near | far from the | far from | next to the | next to | beside the | beside | above the | above | below the | below | inside the | inside | outside the | outside | around the | around | about the | about | before the | before | after the | after | during the | during | while the | while | until the | until | since the | since | for the | for | ago the | ago | from the | from | to the | to | into the | into | onto the | onto | upon the | upon | out of the | out of | off the | off | up the | up | down the | down | across the | across | through the | through | around the | around | over the | over | under the | under | behind the | behind | in front of the | in front of | near the | near | far from the | far from | next to the | next to | beside the | beside | above the | above | below the | below | inside the | inside | outside the | outside | around the | around | about the | about | before the | before | after the | after | during the | during | while the | while | until the | until | since the | since | for the | for | ago the | ago | from the | from | to the | to | into the | into | onto the | onto | upon the | upon | out of the | out of | off the | off | up the | up | down the | down | across the | across | through the | through | around the | around | over the | over | under the | under | behind the | behind | in front of the | in front of | near the | near | far from the | far from | next to)"
	const nouns = "(Sun|Moon|Star|Sky|Cloud|Rain|Snow|Frost|Hail|Thunder|Lightning|Aurora|Rainbow|Flame|Blaze|Inferno|Pyre|Furnace|Forge|Anvil|Hammer|Sword|Axe|Spear|Bow|Arrow|Shield|Armor|Helm|Crown|Cloak|Robe|Gown|Dress|Tunic|Vest|Shirt|Pants|Boots|Gloves|Gauntlets|Bracers|Belt|Cape|Pendant|Necklace|Ring|Bracelet|Earring|Brooch|Tiara|Circlet|Diadem|Crown|Scepter|Staff|Wand|Rod|Tome|Book|Scroll|Parchment|Ink|Quill|Pen|Sword|Blade|Dagger|Knife|Axe|Hammer|Mace|Flail|Morningstar|Spear|Lance|Halberd|Pike|Bow|Crossbow|Sling)"

	const surnamePattern = titles + articles + nouns

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
