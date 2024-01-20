package story_test

import (
	"context"
	"errors"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/story"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	createStory, _ := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("non-existing-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrWorldDoesNotExist)
}

func TestCreateStory_when_unexpected_error_happens_querying_world_repository_it_should_return_internal_error(t *testing.T) {
	// Given
	createStory, _ := newCreateStoryUseCase(t)

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
	assert.Equal(t, "a-world-id", *repository.PersistedStoryWorldID())
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
}

func (m *mockRepository) GetWorldByID(_ context.Context, worldID domain.WorldID) (*domain.World, error) {
	switch worldID {
	case "unexpected-error":
		return nil, errors.New("unexpected error")
	case "non-existing-world-id":
		return nil, domain.ErrWorldDoesNotExist
	default:
		return new(domain.WorldBuilder).
			WithID(worldID).
			WithName("a world name").
			Build(), nil
	}
}

func (m *mockRepository) SaveStory(_ context.Context, story *domain.Story) error {
	m.persistedStory = story
	return nil
}

func (m *mockRepository) PersistedStoryWorldID() *string {
	if m.persistedStory == nil {
		return nil
	}
	return pointerTo(string(m.persistedStory.WorldID()))
}

func pointerTo[T any](value T) *T {
	return &value
}
