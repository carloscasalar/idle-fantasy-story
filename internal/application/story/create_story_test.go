package story_test

import (
	"context"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/story"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateStory_should_require_a_world(t *testing.T) {
	// Given
	createStory := newCreateStoryUseCase(t)

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
			createStory := newCreateStoryUseCase(t)

			// When
			_, err := createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(tc.partySize))

			// Then
			assert.ErrorIs(t, err, story.ErrInvalidPartySize)
		})
	}
}

func TestCreateStory_when_world_does_not_exist_it_should_return_error(t *testing.T) {
	// Given
	createStory := newCreateStoryUseCase(t)

	// When
	_, err := createStory.Execute(context.Background(), newStoryRequestWithWorldID("non-existing-world-id"))

	// Then
	assert.ErrorIs(t, err, application.ErrWorldDoesNotExist)
}

func newCreateStoryUseCase(t *testing.T) *story.CreateStory {
	createStory, err := story.NewCreateStory(new(mockWorldRepository))
	require.NoError(t, err)
	return createStory
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

type mockWorldRepository struct {
}

func (m mockWorldRepository) GetWorldByID(_ context.Context, worldID domain.WorldID) (*domain.World, error) {
	if worldID == "non-existing-world-id" {
		return nil, domain.ErrWorldDoesNotExist
	}

	return new(domain.WorldBuilder).
		WithID(worldID).
		WithName("a world name").
		Build(), nil
}
