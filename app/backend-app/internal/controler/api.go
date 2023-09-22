package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
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
	menuItem.ID = "010"
	menuItem.Title = "Home"
	menuItem.Url = "/"
	menuItems = append(menuItems, menuItem)
	////////////////////////////////////////////////////////////////////////////////////////////////////
	submenu := domain.MenuItem{}
	submenu.ID = "021"
	submenu.Title = "web design"
	submenu.Url = "web-design"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	submenu.ID = "022"
	submenu.Title = "web development"
	submenu.Url = "web-dev"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	submenu.ID = "023"
	submenu.Title = "backend"
	submenu.Url = "backend"
	menuItem.MenuItems = append(menuItem.MenuItems, submenu)

	menuItem.ID = "020"
	menuItem.Title = "Services"
	menuItem.Url = "/services"
	menuItems = append(menuItems, menuItem)

	menu := domain.Menu{}
	menu.MenuItems = menuItems
	////////////////////////////////////////////////////////////////////////////////////////////////////
	menuItem.ID = "030"
	menuItem.Title = "About"
	menuItem.Url = "/about"
	menuItem.MenuItems = nil
	menuItems = append(menuItems, menuItem)
	////////////////////////////////////////////////////////////////////////////////////////////////////

	c.JSON(http.StatusOK, menuItems)
}

// [
//   {
//     title: 'Home',
//     url: '/',
//   },
//   {
//     title: 'Services',
//     url: '/services',
//     submenu: [
//       {
//         title: 'web design',
//         url: 'web-design',
//       },
//       {
//         title: 'web development',
//         url: 'web-dev',
//         submenu: [
//           {
//             title: 'Frontend',
//             url: 'frontend',
//           },
//           {
//             title: 'Backend',
//             submenu: [
//               {
//                 title: 'NodeJS',
//                 url: 'node',
//               },
//               {
//                 title: 'PHP',
//                 url: 'php',
//               },
//             ],
//           },
//         ],
//       },
//       {
//         title: 'SEO',
//         url: 'seo',
//       },
//     ],
//   },
//   {
//     title: 'About',
//     url: '/about',
//     submenu: [
//       {
//         title: 'Who we are',
//         url: 'who-we-are',
//       },
//       {
//         title: 'Our values',
//         url: 'our-values',
//       },
//     ],
//   },
// ];
