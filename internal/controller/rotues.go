package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuminekosan/hikariLibBackend/internal/service"
)

type Routes struct {
	services *service.Service
}

func NewRoutes(services *service.Service) *Routes {
	return &Routes{services: services}
}

func (r *Routes) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("signUp", r.signUp)
		auth.POST("signIn", r.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", r.createList)
			lists.GET("/", r.getAllLists)
			lists.GET("/:id", r.getListById)
			lists.PUT("/:id", r.updateList)
			lists.DELETE("/:id", r.deleteList)

			items := api.Group(":id/items")
			{
				items.POST("/", r.createItem)
				items.GET("/", r.getAllItems)
				items.GET("/:item_id", r.getItemById)
				items.PUT("/:item_id", r.updateItem)
				items.DELETE("/:item_id", r.deleteItem)
			}
		}
	}

	return router
}
