package main

import (
	"fmt"

	"github.com/aniket0951.com/product-service/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.SetTrustedProxies(nil)

	//defer config.CloseClientDB()

	router.Static("static", "static")

	routers.UserRouter(router)
	routers.ProductRouter(router)
	routers.CategoryRouter(router)
	fmt.Println("Application run success...")
	runErr := router.Run(":5050")
	fmt.Println("RunErr : ", runErr)
}
