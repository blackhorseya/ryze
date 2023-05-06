package restful

import (
	"net/http"
	"time"

	_ "github.com/blackhorseya/ryze/api/docs" // import swagger spec
	v1 "github.com/blackhorseya/ryze/internal/adapter/restful/v1"
	"github.com/blackhorseya/ryze/pkg/adapter"
	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/blackhorseya/ryze/pkg/cors"
	bb "github.com/blackhorseya/ryze/pkg/entity/domain/block/biz"
	"github.com/blackhorseya/ryze/pkg/er"
	"github.com/blackhorseya/ryze/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	router *gin.Engine
	biz    bb.IBiz
}

// NewImpl returns a new restful adapter.
func NewImpl(logger *zap.Logger, router *gin.Engine, biz bb.IBiz) adapter.Restful {
	router.Use(cors.AddAllowAll())
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
	}))
	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	instance := &impl{
		router: router,
		biz:    biz,
	}

	return instance
}

func (i *impl) InitRouting() error {
	api := i.router.Group("api")
	{
		api.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.GET("health", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.OK)
		})

		v1.Handle(api.Group("v1"), i.biz)
	}

	return nil
}

// ProviderSet is provider set of restful.
var ProviderSet = wire.NewSet(NewImpl)
