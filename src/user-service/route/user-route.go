package route

import (
	"chilindo/src/user-service/controller"
	"chilindo/src/user-service/middleware"
	"chilindo/src/user-service/service"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	GetRouter()
}

type UserRouteDefault struct {
	UserController controller.IUserController
	Router         *gin.Engine
}

func (u *UserRouteDefault) GetRouter() {
	newUserRoute(u.UserController, u.Router)
}

func newUserRoute(controller controller.IUserController, group *gin.Engine) {
	userRoute := group.Group("/chilindo/user")
	{
		userRoute.POST("/sign-up", controller.SignUp)
		userRoute.POST("/sign-in", controller.SignIn)
	}
	userAuthRoute := group.Group("/chilindo/user", middleware.AuthorizeJWT(service.NewJWTService()))
	{
		userAuthRoute.PUT("/update", controller.Update)
	}
}

func NewUserRouteDefault(userController controller.IUserController, router *gin.Engine) *UserRouteDefault {
	return &UserRouteDefault{UserController: userController, Router: router}
}
