package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/middleware"
	"github.com/carloscasalar/idle-fantasy-story/internal/infrastructure/status"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

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

	var wg sync.WaitGroup
	wg.Add(1)
	go i.startRestService(ctx, &wg)
	wg.Wait()
	return nil
}

func (i *Instance) Stop(ctx context.Context) error {
	return i.srv.Shutdown(ctx)
}

func (i *Instance) startRestService(ctx context.Context, wg *sync.WaitGroup) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		middleware.StructuredLogger(),
		gin.Recovery(),
	)
	v1 := router.Group("/v1")
	v1.GET("/status", status.Handler)

	port := fmt.Sprintf(":%v", i.config.Port)
	const readHeaderSecondsTimeout = 3
	i.srv = &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: readHeaderSecondsTimeout * time.Second,
	}
	log.WithContext(ctx).Infof("Server started listening on port %v", i.config.Port)
	if err := i.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithContext(ctx).Errorf("error starting REST server: %v", err)
	}
	wg.Done()
}
