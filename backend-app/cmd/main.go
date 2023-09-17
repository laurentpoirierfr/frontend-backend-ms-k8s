package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/laurentpoirierfr/bff/internal/router"
	"github.com/laurentpoirierfr/bff/util"
)

// @title bff
// @version 1.0
// @description This is a bff server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host homezone.swagger.io:8080
// @BasePath /
func main() {
	ctx := context.Background()

	config, err := util.LoadConfig(".")
	if err != nil {
		slog.Fatal("cannot load config")
	}
	app := gin.Default()

	router.NewRouter(ctx, config, app)

	// Listen and Serve
	if err := app.Run(":" + config.Port); err != nil {
		log.Fatal(err.Error())
	}
}
