package scan

import (
	"fmt"

	"github.com/blackhorseya/ryze/adpater/block/wirex"
	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/responsex"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type restful struct {
	injector *wirex.Injector
	server   *httpx.Server
}

// NewRestful is used to create a new adapterx.Restful
func NewRestful(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	return &restful{
		injector: injector,
		server:   server,
	}
}

func (i *restful) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		ctx.Error("Failed to init routing", zap.Error(err))
		return err
	}

	err = i.server.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start server", zap.Error(err))
		return err
	}

	swaggerURL := fmt.Sprintf("http://localhost:%d/api/docs/index.html", i.injector.A.HTTP.Port)
	ctx.Info("start restful server", zap.String("swagger_url", swaggerURL))

	return nil
}

func (i *restful) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

func (i *restful) InitRouting() error {
	router := i.server.Router

	// api
	api := router.Group("/api")
	{
		api.GET("/healthz", i.healthz)
		api.GET("/docs/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName("block_scan"),
		))
	}

	return nil
}

func (i *restful) GetRouter() *gin.Engine {
	return i.server.Router
}

// healthz is used to handle health check
func (i *restful) healthz(c *gin.Context) {
	responsex.OK(c, nil)
}
