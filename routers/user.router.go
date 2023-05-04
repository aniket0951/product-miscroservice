package routers

import (
	"github.com/aniket0951.com/product-service/controller"
	"github.com/aniket0951.com/product-service/repository"
	"github.com/aniket0951.com/product-service/services"
	"github.com/gin-gonic/gin"
)

var (
	userrepo       = repository.NewUserRepository()
	userservice    = services.NewUserService(userrepo)
	usercontroller = controller.NewUserController(userservice)
)

func UserRouter(router *gin.Engine) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("/create-account", usercontroller.CreateUserAccount)
		userRoutes.GET("/fetch-user", usercontroller.GetUserByID)
	}

	addressRoute := router.Group("/api/user/address")
	{
		addressRoute.POST("/add-user-address", usercontroller.AddUserAddress)
	}

}
