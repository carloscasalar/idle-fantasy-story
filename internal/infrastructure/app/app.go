package app

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"context"
	"errors"
	"fmt"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/middleware"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/status"
	idlefantasystoryv1 "github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1"
	"github.com/carloscasalar/idle-fantasy-story/pkg/idlefantasystory/v1/idlefantasystoryv1connect"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"sync"
	"time"
)

const readHeaderSecondsTimeout = 3 * time.Second

type Instance struct {
	config      *Config
	srv         *http.Server
	grpcService *http.Server
}

func New(config *Config) *Instance {
	return &Instance{
		config: config,
	}
}

func (i *Instance) Start(ctx context.Context) error {
	if err := ConfigureLogger(i.config.Log); err != nil {
		return fmt.Errorf("unable to configure logger: %w", err)
	}
	log.WithContext(ctx).Debugf("Api configuration: %v", i.config)

	gin.SetMode(gin.ReleaseMode)
	var wg sync.WaitGroup
	wg.Add(2)
	go i.startRestService(ctx, &wg)
	go i.startGrpcService(ctx, &wg)
	wg.Wait()
	return nil
}

func (i *Instance) Stop(ctx context.Context) error {
	errCh := make(chan error, 2) // buffer size of 2, one for each service

	go func() {
		if err := i.srv.Shutdown(ctx); err != nil {
			errCh <- fmt.Errorf("error stopping REST server: %w", err)
		}
		errCh <- nil
	}()

	go func() {
		if err := i.grpcService.Shutdown(ctx); err != nil {
			errCh <- fmt.Errorf("error stopping GRPC server: %w", err)
		}
		errCh <- nil
	}()

	firstError := <-errCh
	secondError := <-errCh

	if firstError != nil {
		return firstError
	}
	return secondError
}

func (i *Instance) startRestService(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	router := gin.New()
	router.Use(
		middleware.StructuredLogger(),
		gin.Recovery(),
	)
	v1 := router.Group("/v1")
	v1.GET("/status", status.Handler)

	port := fmt.Sprintf(":%v", i.config.Port)
	i.srv = &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: readHeaderSecondsTimeout,
	}
	log.WithContext(ctx).Infof("Server started listening on port %v", i.config.Port)
	if err := i.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithContext(ctx).Errorf("error starting REST server: %v", err)
	}
}

func (i *Instance) startGrpcService(ctx context.Context, wg *sync.WaitGroup) {
	wg.Done()
	router := gin.New()
	router.Use(
		middleware.StructuredLogger(),
		gin.Recovery(),
	)
	storyService := new(storyServiceServer)
	storyService.RegisterService(router.Group("/"))

	i.grpcService = &http.Server{
		Addr:              fmt.Sprintf(":%v", i.config.Grpc.Port),
		Handler:           h2c.NewHandler(router, &http2.Server{}),
		ReadHeaderTimeout: readHeaderSecondsTimeout,
	}
	log.WithContext(ctx).Infof("GRPC Server started listening on port %v", i.config.Grpc.Port)
	if err := i.grpcService.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithContext(ctx).Errorf("error starting GRPC server: %v", err)
	}
}

// TODO move this to a different file
type storyServiceServer struct {
	idlefantasystoryv1connect.UnimplementedStoryServiceHandler
}

func (s *storyServiceServer) RegisterService(g *gin.RouterGroup) {
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

func (s *storyServiceServer) GetWorldState(
	ctx context.Context,
	req *connect.Request[idlefantasystoryv1.GetWorldStateRequest],
) (*connect.Response[idlefantasystoryv1.GetWorldStateResponse], error) {
	id := req.Msg.GetWorldId()
	log.WithContext(ctx).Infof("GetWorldState request received for world %v", id)
	return connect.NewResponse(&idlefantasystoryv1.GetWorldStateResponse{}), nil
}
