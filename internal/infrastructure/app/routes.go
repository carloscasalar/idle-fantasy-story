package app

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	idlefantasystoryv1 "github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1"
	"github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1/idlefantasystoryv1connect"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type routes struct {
	idlefantasystoryv1connect.UnimplementedStoryServiceHandler
}

func (s *routes) Register(g *gin.RouterGroup) {
	path, handler := idlefantasystoryv1connect.NewStoryServiceHandler(s, connect.WithInterceptors(
	// TODO add interceptors
	))
	g.Any(path+"/*path", gin.WrapH(handler))

	// Enable reflection
	reflector := grpcreflect.NewStaticReflector(idlefantasystoryv1connect.StoryServiceName)
	pathV1, handlerV1 := grpcreflect.NewHandlerV1(reflector)
	g.POST(fmt.Sprintf("%vServerReflectionInfo", pathV1), gin.WrapH(handlerV1))
	pathV1Alpha, handlerV1Alpha := grpcreflect.NewHandlerV1Alpha(reflector)
	g.POST(fmt.Sprintf("%v/ServerReflectionInfo", pathV1Alpha), gin.WrapH(handlerV1Alpha))
}

func (s *routes) GetWorldState(
	ctx context.Context,
	req *connect.Request[idlefantasystoryv1.GetWorldStateRequest],
) (*connect.Response[idlefantasystoryv1.GetWorldStateResponse], error) {
	id := req.Msg.GetWorldId()
	log.WithContext(ctx).Infof("GetWorldState request received for world %v", id)
	return connect.NewResponse(&idlefantasystoryv1.GetWorldStateResponse{}), nil
}
