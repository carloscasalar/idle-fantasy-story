package handlers

import (
	"context"

	"github.com/carloscasalar/idle-fantasy-story/internal/application/world"

	"connectrpc.com/connect"
	idlefantasystoryv1 "github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1"
	log "github.com/sirupsen/logrus"
)

func (r *Routes) GetWorlds(
	ctx context.Context,
	_ *connect.Request[idlefantasystoryv1.GetWorldsRequest],
) (*connect.Response[idlefantasystoryv1.GetWorldsResponse], error) {
	log.WithContext(ctx).Info("GetWorlds request received for world")
	worlds, err := r.getWorlds.Execute(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Error getting worlds")
		return nil, err
	}

	return connect.NewResponse(&idlefantasystoryv1.GetWorldsResponse{
		Worlds: toGRPCWorlds(worlds),
	}), nil
}

func toGRPCWorlds(worlds []world.WorldDTO) []*idlefantasystoryv1.World {
	var grpcWorlds []*idlefantasystoryv1.World
	for _, w := range worlds {
		grpcWorlds = append(grpcWorlds, &idlefantasystoryv1.World{
			WorldId: w.ID,
			Name:    w.Name,
		})
	}
	return grpcWorlds
}
