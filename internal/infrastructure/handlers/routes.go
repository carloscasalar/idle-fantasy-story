package handlers

import (
	"fmt"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/carloscasalar/idle-fantasy-story/internal/application/world"
	"github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1/idlefantasystoryv1connect"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	idlefantasystoryv1connect.UnimplementedStoryServiceHandler

	getWorlds *world.GetWorlds
}

func NewRoutes(getWorlds *world.GetWorlds) *Routes {
	return &Routes{
		getWorlds: getWorlds,
	}
}

func (r *Routes) Register(g *gin.RouterGroup) {
	path, handler := idlefantasystoryv1connect.NewStoryServiceHandler(r, connect.WithInterceptors(
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
