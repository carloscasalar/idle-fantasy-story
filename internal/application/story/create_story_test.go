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

func newStoryRequestWithoutWorld() story.CreateStoryRequest {
	return story.CreateStoryRequest{}
}
