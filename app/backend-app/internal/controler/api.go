package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/laurentpoirierfr/bff/internal/domain"
	"github.com/laurentpoirierfr/bff/util"
)

// IndexController is the default controller
type IndexController struct {
	context context.Context
	config  util.Config
}

func NewIndexController(context context.Context, config util.Config) *IndexController {
	return &IndexController{
		config:  config,
		context: context,
	}
}

// Menu
//
//	@Summary	Get versioni
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	domain.Menu{}	"ok"
//	@Router		/api/menus [get]
func (ctrl *IndexController) Menus(c *gin.Context) {
	var menuItems []domain.MenuItem

	menuItem := domain.MenuItem{}
	menuItem.ID = UUIDGenerate()
	menuItem.Title = "Home"
	menuItem.Url = "/"
	menuItems = append(menuItems, menuItem)
	////////////////////////////////////////////////////////////////////////////////////////////////////
	submenu := domain.MenuItem{}
	submenu.ID = UUIDGenerate()
	submenu.Title = "ms-service-01"
	submenu.Url = "ms-service-01"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	submenu.ID = UUIDGenerate()
	submenu.Title = "ms-service-02"
	submenu.Url = "ms-service-02"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	submenu.ID = UUIDGenerate()
	submenu.Title = "ms-service-03"
	submenu.Url = "ms-service-03"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	menuItem.ID = UUIDGenerate()
	menuItem.Title = "Services"
	menuItem.Url = "/services"
	menuItems = append(menuItems, menuItem)

	menu := domain.Menu{}
	menu.MenuItems = menuItems
	////////////////////////////////////////////////////////////////////////////////////////////////////
	menuItem.ID = UUIDGenerate()
	menuItem.Title = "About"
	menuItem.Url = "/about"
	menuItem.MenuItems = nil
	menuItems = append(menuItems, menuItem)
	////////////////////////////////////////////////////////////////////////////////////////////////////
	menuItem.ID = UUIDGenerate()
	menuItem.Title = "Login"
	menuItem.Url = "/login"
	menuItem.MenuItems = nil
	menuItems = append(menuItems, menuItem)
	////////////////////////////////////////////////////////////////////////////////////////////////////

	c.JSON(http.StatusOK, menuItems)
}

func UUIDGenerate() string {
	id, _ := uuid.NewV4()
	return id.String()
}
