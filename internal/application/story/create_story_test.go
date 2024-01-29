package story_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/pkg/utils"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/story"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateStory_regarding_party_size_in_persisted_story(t *testing.T) {
	numberOfCharactersTestCases := map[string]struct {
		partySize                  *int
		expectedNumberOfCharacters int
	}{
		"should be 3 when party size is 3":             {utils.PointerTo(3), 3},
		"should be 4 when party size is not specified": {nil, 4},
		"should be 4 when party size is 4":             {utils.PointerTo(4), 4},
		"should be 6 when party size is 5":             {utils.PointerTo(5), 5},
		"should be 6 when party size is 6":             {utils.PointerTo(6), 6},
	}

	for name, tc := range numberOfCharactersTestCases {
		t.Run(name, func(t *testing.T) {
			// Given
			createStory, repository, _ := newCreateStoryUseCase(t)

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(uint8(tc.expectedNumberOfCharacters)))

			// Then
			require.NoError(t, err)
			require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
			characters := repository.persistedStory.Characters()
			assert.Len(t, characters, tc.expectedNumberOfCharacters)
		})
	}
}

func TestCreateStory_every_character_should_have_a_unique_character_id(t *testing.T) {
	// Given
	createStory, repository, _ := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(6))

	// Then
	require.NoError(t, err)
	require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
	characters := repository.persistedStory.Characters()
	characterIDs := make(map[domain.CharacterID]bool)
	for _, character := range characters {
		characterIDs[character.ID()] = true
	}
	assert.Len(t, characterIDs, 6)
}

func TestCreateStory_for_a_party_of_6_character(t *testing.T) {
	namesToGenerate := []string{"name 1", "name 2", "name 3", "name 4", "name 5", "name 6"}
	for i, expectedName := range namesToGenerate {
		t.Run(fmt.Sprintf("character at %d position should have been generated with '%v' name", i+1, expectedName), func(t *testing.T) {
			// Given
			createStory, repository, _ := newCreateStoryUseCase(t, withNamesToGenerate(namesToGenerate))

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(6))

			// Then
			require.NoError(t, err)
			persistedStory := repository.persistedStory
			require.NotNil(t, persistedStory, "persisted story should not be nil")
			require.Equal(t, 6, persistedStory.PartySize())
			character := persistedStory.Characters()[i]
			assert.Equal(t, expectedName, character.Name())
		})
	}
}

func TestCreateStory_all_characters_should_have_a_species_from_the_world(t *testing.T) {
	// Given
	worldSpecies := []domain.Species{domain.SpeciesElf, domain.SpeciesDwarf}
	createStory, repository, _ := newCreateStoryUseCase(t, withWorldSpecies(worldSpecies))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(6))

	// Then
	require.NoError(t, err)
	require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
	characters := repository.persistedStory.Characters()
	require.Len(t, characters, 6)
	for i, character := range characters {
		assert.NotEmpty(t, character.Species(), fmt.Sprintf("character at %d position should have a species", i+1))
		assert.NoError(t, verifyCharacterAtPosHasAnyOfExpectedSpecies(i+1, character.Species(), domain.SpeciesElf, domain.SpeciesDwarf))
	}
}

func verifyCharacterAtPosHasAnyOfExpectedSpecies(position int, current domain.Species, expectedSpecies ...domain.Species) error {
	for _, expected := range expectedSpecies {
		if current == expected {
			return nil
		}
	}
	return fmt.Errorf("expected species to be one of %v but was %v for character at position %d", expectedSpecies, current, position)
}

func TestCreateStory_all_character_names_should_be_generated_using_the_character_species(t *testing.T) {
	// Given
	worldSpecies := []domain.Species{domain.SpeciesGnome}
	createStory, _, nameGenerator := newCreateStoryUseCase(t, withWorldSpecies(worldSpecies))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(4))

	// Then
	require.NoError(t, err)
	require.NotNil(t, nameGenerator, "name generator should not be nil")
	nameGenerator.assertGenerateCharacterNameNthCallCalledWithSpecies(t, 0, domain.SpeciesGnome)
	nameGenerator.assertGenerateCharacterNameNthCallCalledWithSpecies(t, 1, domain.SpeciesGnome)
	nameGenerator.assertGenerateCharacterNameNthCallCalledWithSpecies(t, 2, domain.SpeciesGnome)
	nameGenerator.assertGenerateCharacterNameNthCallCalledWithSpecies(t, 3, domain.SpeciesGnome)
}

func TestCreateStory_should_properly_map_world_id_of_the_persisted_story(t *testing.T) {
	// Given
	createStory, _, _ := newCreateStoryUseCase(t)

	// When
	dto, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("a-world-id"))

	// Then
	require.NoError(t, err)
	assert.Equal(t, "a-world-id", dto.WorldID)
}

func TestCreateStory_should_require_a_world(t *testing.T) {
	// Given
	createStory, _, _ := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithoutWorld())

	// Then
	assert.ErrorIs(t, err, story.ErrWorldIDIsRequired)
}

func TestCreateStory_when_specified_party_size(t *testing.T) {
	testCases := map[string]struct {
		partySize uint8
	}{
		"should require party size not be less than 3":    {2},
		"should require party size not be greater than 6": {7},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Given
			createStory, _, _ := newCreateStoryUseCase(t)

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(tc.partySize))

			// Then
			assert.ErrorIs(t, err, story.ErrInvalidPartySize)
		})
	}
}

func TestCreateStory_when_world_does_not_exist_it_should_return_error(t *testing.T) {
	// Given
	createStory, _, _ := newCreateStoryUseCase(t, withErrorOnQueryStory(domain.ErrWorldDoesNotExist))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("non-existing-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrWorldDoesNotExist)
}

func TestCreateStory_when_unexpected_error_happens_querying_world_repository_it_should_return_internal_error(t *testing.T) {
	// Given
	createStory, _, _ := newCreateStoryUseCase(t, withErrorOnQueryStory(errors.New("unexpected error")))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("unexpected-error"))

	// Then
	assert.ErrorIs(t, err, application.ErrInternalServer)
}

func TestCreateStory_should_persist_a_new_story_with_the_specified_world(t *testing.T) {
	// Given
	createStory, repository, _ := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("a-world-id"))

	// Then
	require.NoError(t, err)
	require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
	assert.Equal(t, domain.WorldID("a-world-id"), repository.persistedStory.WorldID())
}

func TestCreateStory_when_error_happens_persisting_story_it_should_return_internal_error(t *testing.T) {
	// Given
	createStory, _, _ := newCreateStoryUseCase(t, withErrorOnSaveStory(errors.New("unexpected error")))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("a-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrInternalServer)
}

func newCreateStoryUseCase(t *testing.T, opts ...mocksOption) (*story.CreateStory, *mockRepository, *mockNameGenerator) {
	options := mocksSettings{}
	for _, opt := range opts {
		opt(&options)
	}

	defaultSpecies := []domain.Species{domain.SpeciesHuman}

	repository := &mockRepository{
		persistedStory: nil,
		worldSpecies:   utils.NoEmptySlice(options.worldSpecies, defaultSpecies),
		errorOnSave:    options.errorOnSaveStory,
		errorOnQuery:   options.errorOnQueryStory,
	}
	nameGenerator := &mockNameGenerator{names: options.namesToGenerate}
	createStory, err := story.NewCreateStory(repository, nameGenerator)
	require.NoError(t, err)
	return createStory, repository, nameGenerator
}

func newStoryRequestWithWorldID(worldID string) story.CreateStoryRequest {
	return story.CreateStoryRequest{
		WorldID: worldID,
	}
}

func newStoryRequestWithNumberOfCharacters(numberOfCharacters uint8) story.CreateStoryRequest {
	return story.CreateStoryRequest{
		WorldID:   "whatever",
		PartySize: &numberOfCharacters,
	}
}

func newStoryRequestWithoutWorld() story.CreateStoryRequest {
	return story.CreateStoryRequest{}
}

type mocksSettings struct {
	namesToGenerate   []string
	worldSpecies      []domain.Species
	errorOnSaveStory  error
	errorOnQueryStory error
}

type mocksOption func(settings *mocksSettings)

func withNamesToGenerate(names []string) mocksOption {
	return func(settings *mocksSettings) {
		settings.namesToGenerate = names
	}
}

func withWorldSpecies(species []domain.Species) mocksOption {
	return func(settings *mocksSettings) {
		settings.worldSpecies = species
	}
}

func withErrorOnSaveStory(err error) mocksOption {
	return func(settings *mocksSettings) {
		settings.errorOnSaveStory = err
	}
}

func withErrorOnQueryStory(err error) mocksOption {
	return func(settings *mocksSettings) {
		settings.errorOnQueryStory = err
	}
}

type mockRepository struct {
	persistedStory *domain.Story
	worldSpecies   []domain.Species
	errorOnSave    error
	errorOnQuery   error
}

func (m *mockRepository) GetWorldByID(_ context.Context, worldID domain.WorldID) (*domain.World, error) {
	if m.errorOnQuery != nil {
		return nil, m.errorOnQuery
	}
	return new(domain.WorldBuilder).
		WithID(worldID).
		WithSpecies(m.worldSpecies).
		WithName("a world name").
		Build(), nil
}

func (m *mockRepository) SaveStory(_ context.Context, story *domain.Story) error {
	if m.errorOnSave != nil {
		return m.errorOnSave
	}
	m.persistedStory = story
	return nil
}

type mockNameGenerator struct {
	names             []string
	nextNameIndex     *int
	calledWithSpecies []domain.Species
}

func (m *mockNameGenerator) GenerateCharacterName(species domain.Species) string {
	m.calledWithSpecies = append(m.calledWithSpecies, species)

	if len(m.names) == 0 {
		return ""
	}
	nameIndex := utils.NoNilValue(m.nextNameIndex, 0)
	m.nextNameIndex = utils.PointerTo(nameIndex + 1)
	return m.names[nameIndex]
}

func (m *mockNameGenerator) assertGenerateCharacterNameNthCallCalledWithSpecies(t *testing.T, callIndex int, expectedSpeciesCalled domain.Species) {
	require.True(t, len(m.calledWithSpecies) >= callIndex+1, "GenerateCharacterName was not called %d times", callIndex+1)
	assert.Equal(t, expectedSpeciesCalled, m.calledWithSpecies[callIndex])
}
