package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/middleware"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/system"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const readHeaderSecondsTimeout = 3 * time.Second

type Instance struct {
	config *Config
	srv    *http.Server
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

	router := gin.New()
	router.Use(
		middleware.StructuredLogger(),
		gin.Recovery(),
	)

	rootPath := router.Group("/")

	appRoutes := new(routes)
	appRoutes.Register(rootPath)

	systemRoutes := system.NewRoutes()
	systemRoutes.Register(rootPath)

	i.srv = &http.Server{
		Addr:              fmt.Sprintf(":%v", i.config.Port),
		Handler:           h2c.NewHandler(router, &http2.Server{}),
		ReadHeaderTimeout: readHeaderSecondsTimeout,
	}
	log.WithContext(ctx).Infof("Server started listening on port %v", i.config.Port)
	if err := i.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %v", err)
	}
	return nil
}

func (i *Instance) Stop(ctx context.Context) error {
	if err := i.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("error stopping REST server: %w", err)
	}
	return nil
}
