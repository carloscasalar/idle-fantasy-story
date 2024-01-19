package inmemory_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/storage/inmemory"
)

func TestRepository_GetWorlds(t *testing.T) {
	repo, err := inmemory.NewRepository(context.Background(), "./testfiles/worlds.yml")
	require.NoError(t, err)

	t.Run("should return properly load all worlds", func(t *testing.T) {
		worlds, err := repo.GetWorlds(context.Background())

		require.NoError(t, err)
		assert.Len(t, worlds, 2)
	})

	tcExpectedWorlds := []struct {
		ID           string
		ExpectedName string
	}{
		{"krynn", "Krynn"},
		{"aebrynis", "Aebrynis"},
	}
	for _, tc := range tcExpectedWorlds {
		t.Run(fmt.Sprintf("should return properly load world '%v'", tc.ID), func(t *testing.T) {
			worlds, err := repo.GetWorlds(context.Background())
			require.NoError(t, err)

			world, found := getWorld(worlds, tc.ID)
			require.True(t, found)
			assert.Equal(t, tc.ExpectedName, world.Name())
		})
	}
}

func getWorld(worlds []domain.World, id string) (world domain.World, found bool) {
	for _, w := range worlds {
		if w.ID() == domain.WorldID(id) {
			return w, true
		}
	}

	return domain.World{}, false
}
