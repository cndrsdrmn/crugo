package routes

import (
	"github.com/cndrsdrmn/crugo/facades"
	"github.com/cndrsdrmn/crugo/users"
	"github.com/gin-gonic/gin"
)

func RouteRegistar(route *gin.Engine) {
	repo := users.NewRepository(facades.DB)
	srvs := users.NewService(repo)
	ctrl := users.NewController(srvs)

	route.GET("/users", ctrl.Index)
	route.POST("/users", ctrl.Store)
	route.GET("/users/:id", ctrl.Show)
	route.PATCH("/users/:id", ctrl.Update)
	route.DELETE("/users/:id", ctrl.Destroy)
}
