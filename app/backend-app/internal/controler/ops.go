package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laurentpoirierfr/bff/internal/domain"
	"github.com/laurentpoirierfr/bff/util"
)

// IndexController is the default controller
type OpsController struct {
	context context.Context
	config  util.Config
}

func NewOpsController(context context.Context, config util.Config) *OpsController {
	return &OpsController{
		config:  config,
		context: context,
	}
}

// Info micro service
// @Summary Info
// @Schemes
// @Description Informations sur le service
// @Tags ops
// @Accept json
// @Produce json
// @Success 200 {object} domain.Info
// @Router /ops/info [get]
func (ctrl *OpsController) Info(c *gin.Context) {
	c.JSON(200, domain.Info{
		Version:     "0.1.0",
		Name:        "backend-app",
		Description: "Service backend-app.",
	})
}
