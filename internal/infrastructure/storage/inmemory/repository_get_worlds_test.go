package inmemory_test

import (
	"context"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/storage/inmemory"
)

func TestRepository_GetWorlds(t *testing.T) {
	repo, err := inmemory.NewRepository(context.Background(), "./testfiles/worlds.yml")
	require.NoError(t, err)

	t.Run("should return properly load all worlds", func(t *testing.T) {
		worlds, err := repo.GetWorlds(context.Background())

		require.NoError(t, err)
		require.Len(t, worlds, 2)
		assert.Equal(t, domain.WorldID("aebrynis"), worlds[0].ID())
		assert.Equal(t, "Aebrynis", worlds[0].Name())
		assert.Equal(t, domain.WorldID("krynn"), worlds[1].ID())
		assert.Equal(t, "Krynn", worlds[1].Name())
	})
}
