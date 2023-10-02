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

	"github.com/penglongli/gin-metrics/ginmetrics"
)

func NewRouter(context context.Context, config util.Config, app *gin.Engine) {

	// ==========================================================================================
	// Monitoring Prometheus
	// ==========================================================================================
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/ops/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(app)
	// ==========================================================================================
	// Swagger
	// ==========================================================================================
	app.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.WrapHandler(swaggerfiles.Handler)(context)
	})

	app.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	indexController := controller.NewIndexController(context, config)

	// ==========================================================================================
	// API
	// ==========================================================================================
	api := app.Group("/api")
	{
		api.GET("/menus", indexController.Menus)
	}

	// ==========================================================================================
	// OPS
	// ==========================================================================================
	opsCtrl := controller.NewOpsController(context, config)

	ops := app.Group("/ops")
	{
		ops.GET("/info", opsCtrl.Info)
	}

}
