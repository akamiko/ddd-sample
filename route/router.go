package route

import (
	"github.com/akamiko/ddd-sample/controller"

	"github.com/labstack/echo"
)

// Init ルーティング
func Init() *echo.Echo {
	e := echo.New()
	api := e.Group("/api")
	{
		api.GET("/members", controller.GetMembers())
		//api.POST("/members", controller.PostMembers())
		//api.PUT("/members", controller.PutMembers())
		//api.DELETE("/members", controller.DeleteMembers())
	}
	return e
}
