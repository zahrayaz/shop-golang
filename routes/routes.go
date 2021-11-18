package routes

import (
	"myproject/controllers/api"
	"os"

	"github.com/labstack/echo"
)

func Routes() {
	e := echo.New()
	//user
	e.GET("/user/index", api.Index)
	e.GET("/user/show/:id", api.Show)
	e.POST("/user/store", api.Store)
	//	e.PUT("/user/update/:id", api.Update)
	e.DELETE("/user/delete/:id", api.Delete)
	//product
	/* e.GET("/product/index", api.Index)
	e.GET("/product/show/:id", api.Show)
	e.POST("/product/store", api.Store)
	e.PUT("/product/update/:id", api.Update)
	e.DELETE("/product/delete/:id", api.Delete)
	*/
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
