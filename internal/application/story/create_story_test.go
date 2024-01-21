package story_test

import (
	"context"
	"errors"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/pkg/utils"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/story"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateStory_regarding_party_in_persisted_story(t *testing.T) {
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
			createStory, repository := newCreateStoryUseCase(t)

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(uint8(tc.expectedNumberOfCharacters)))

			// Then
			require.NoError(t, err)
			require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
			party := repository.persistedStory.Party()
			assert.Len(t, party.Characters(), tc.expectedNumberOfCharacters)
		})
	}

	t.Run("every character should have a unique ID", func(t *testing.T) {
		// Given
		createStory, repository := newCreateStoryUseCase(t)

		// When
		_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(6))

		// Then
		require.NoError(t, err)
		require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
		party := repository.persistedStory.Party()
		characters := party.Characters()
		characterIDs := make(map[domain.CharacterID]bool)
		for _, character := range characters {
			characterIDs[character.ID()] = true
		}
		assert.Len(t, characterIDs, len(characters))
	})

	t.Run("every character should have a unique name", func(t *testing.T) {
		// Given
		createStory, repository := newCreateStoryUseCase(t)

		// When
		_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(6))

		// Then
		require.NoError(t, err)
		require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
		party := repository.persistedStory.Party()
		characters := party.Characters()
		characterNames := make(map[string]bool)
		for _, character := range characters {
			characterNames[character.Name()] = true
		}
		assert.Len(t, characterNames, len(characters))
	})
}

func TestCreateStory_should_require_a_world(t *testing.T) {
	// Given
	createStory, _ := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithoutWorld())

	// Then
	assert.ErrorIs(t, err, story.ErrWorldIDIsRequired)
}

func TestCreateStory_when_specified_party_size(t *testing.T) {
	testCases := map[string]struct {
		partySize uint8
	}{
		"should not be less than 3":    {2},
		"should not be greater than 6": {7},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Given
			createStory, _ := newCreateStoryUseCase(t)

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(tc.partySize))

			// Then
			assert.ErrorIs(t, err, story.ErrInvalidPartySize)
		})
	}
}

func TestCreateStory_when_world_does_not_exist_it_should_return_error(t *testing.T) {
	// Given
	createStory, repo := newCreateStoryUseCase(t)
	repo.FailOnQueryWith(domain.ErrWorldDoesNotExist)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("non-existing-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrWorldDoesNotExist)
}

func TestCreateStory_when_unexpected_error_happens_querying_world_repository_it_should_return_internal_error(t *testing.T) {
	// Given
	createStory, repo := newCreateStoryUseCase(t)
	repo.FailOnQueryWith(errors.New("unexpected error"))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("unexpected-error"))

	// Then
	assert.ErrorIs(t, err, application.ErrInternalServer)
}

func TestCreateStory_should_persist_a_new_story_with_the_specified_world(t *testing.T) {
	// Given
	createStory, repository := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("a-world-id"))

	// Then
	require.NoError(t, err)
	require.NotNil(t, repository.persistedStory, "persisted story should not be nil")
	assert.Equal(t, domain.WorldID("a-world-id"), repository.persistedStory.WorldID())
}

func TestCreateStory_when_error_happens_persisting_story_it_should_return_internal_error(t *testing.T) {
	// Given
	createStory, repository := newCreateStoryUseCase(t)
	repository.FailOnSaveWith(errors.New("unexpected error"))

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("a-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrInternalServer)
}

func newCreateStoryUseCase(t *testing.T) (*story.CreateStory, *mockRepository) {
	repository := new(mockRepository)
	createStory, err := story.NewCreateStory(repository)
	require.NoError(t, err)
	return createStory, repository
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

type mockRepository struct {
	persistedStory *domain.Story
	errorOnSave    error
	errorOnQuery   error
}

func (m *mockRepository) GetWorldByID(_ context.Context, worldID domain.WorldID) (*domain.World, error) {
	if m.errorOnQuery != nil {
		return nil, m.errorOnQuery
	}
	return new(domain.WorldBuilder).
		WithID(worldID).
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

func (m *mockRepository) FailOnSaveWith(err error) {
	m.errorOnSave = err
}

func (m *mockRepository) FailOnQueryWith(err error) {
	m.errorOnQuery = err
}
