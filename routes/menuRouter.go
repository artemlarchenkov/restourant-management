package routes

import (
	controller "golang-restourant-management/controlles"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.GetMenus())
	incomingRoutes.PATCH("menus/:menu_id", controller.UpdateMenu())
}
