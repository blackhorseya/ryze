package scan

import (
	"fmt"

	"github.com/blackhorseya/ryze/adapter/block/wirex"
	_ "github.com/blackhorseya/ryze/api/block/scan" // import swagger
	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
	"github.com/blackhorseya/ryze/pkg/adapterx"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/responsex"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Ryze Block Scan API
// @version 0.1.0
// @description Ryze Block Scan API document.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
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

	// TODO: 2024/7/28|sean|add block scan logic here

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

// Healthz is used to check the health of the service.
// @Summary Check the health of the service.
// @Description Check the health of the service.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /healthz [get]
func (i *restful) healthz(c *gin.Context) {
	responsex.OK(c, nil)
}
