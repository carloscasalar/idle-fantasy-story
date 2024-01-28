package domain

type NameGenerator interface {
	GenerateCharacterName(species Species) string
}
