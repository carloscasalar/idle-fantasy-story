package world_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/carloscasalar/idle-fantasy-story/internal/application"
	"github.com/carloscasalar/idle-fantasy-story/internal/application/world"
	"github.com/carloscasalar/idle-fantasy-story/internal/application/world/mocks"
	"github.com/carloscasalar/idle-fantasy-story/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetWorlds_should_retrieve_the_list_of_worlds(t *testing.T) {
	// Given
	worldRepository := new(mocks.Repository)
	worlds := []domain.World{
		*new(domain.WorldBuilder).
			WithID("world-1").
			WithName("World 1").
			Build(),
	}
	worldRepository.On("GetWorlds", mock.Anything).Return(worlds, nil).Maybe()
	getWorlds := world.NewGetWorlds(worldRepository)

	// When
	result, err := getWorlds.Execute(context.Background())

	// Then
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, "world-1", result[0].ID)
	assert.Equal(t, "World 1", result[0].Name)
}

func TestGetWorlds_when_repository_returns_an_error_should_return_an_error(t *testing.T) {
	// Given
	worldRepository := new(mocks.Repository)
	worldRepository.On("GetWorlds", mock.Anything).Return(nil, errors.New("some error")).Maybe()
	getWorlds := world.NewGetWorlds(worldRepository)

	// When
	result, err := getWorlds.Execute(context.Background())

	// Then
	require.Error(t, err)
	assert.ErrorIs(t, err, application.ErrInternalServer)
	assert.Nil(t, result)
}
