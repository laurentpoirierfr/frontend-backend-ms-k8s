package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/laurentpoirierfr/bff/internal/controler"
	"github.com/laurentpoirierfr/bff/util"

	docs "github.com/laurentpoirierfr/bff/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(context context.Context, config util.Config, app *gin.Engine) {
	app.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.WrapHandler(swaggerfiles.Handler)(context)
	})

	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	indexController := controller.NewIndexController(context, config)
	api := app.Group("/api")
	{
		api.GET("/menus", indexController.Menus)
	}

	opsCtrl := controller.NewOpsController(context, config)

	ops := app.Group("/ops")
	{
		ops.GET("/info", opsCtrl.Info)
	}

}
