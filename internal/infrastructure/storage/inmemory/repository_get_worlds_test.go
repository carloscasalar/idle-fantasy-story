package inmemory_test

import (
	"context"
	"testing"

	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/storage/inmemory"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetWorlds(t *testing.T) {
	t.Skip("TODO: enable again when we have a proper parametrized assets folder")
	repo, err := inmemory.NewRepository(context.Background())
	require.NoError(t, err)

	t.Run("should return all worlds", func(t *testing.T) {
		worlds, err := repo.GetWorlds(context.Background())
		require.NoError(t, err)
		require.Len(t, worlds, 2)
	})
}
