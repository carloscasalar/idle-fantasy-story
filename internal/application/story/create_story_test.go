package story_test

import (
	"context"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/story"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateStory_should_require_a_world(t *testing.T) {
	// Given
	createStory, err := story.NewCreateStory()
	require.NoError(t, err)

	// When
	err = createStory.Execute(context.Background(), newStoryRequestWithoutWorld())

	// Then
	assert.ErrorIs(t, err, story.ErrWorldIDIsRequired)
}

func TestCreateStory_when_specified_number_of_characters(t *testing.T) {
	testCases := map[string]struct {
		numberOfCharacters uint8
	}{
		"should not be less than 3":    {2},
		"should not be greater than 6": {7},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// Given
			createStory, err := story.NewCreateStory()
			require.NoError(t, err)

			// When
			err = createStory.Execute(context.Background(), newStoryRequestWithNumberOfCharacters(tc.numberOfCharacters))

			// Then
			assert.ErrorIs(t, err, story.ErrInvalidPartySize)
		})
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
