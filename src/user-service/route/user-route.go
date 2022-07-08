package route

import (
	"chilindo/src/user-service/controller"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	SetRouter()
}
type UserRoute struct {
	UserController controller.IUserController
	Router         *gin.Engine
}

func (u UserRoute) SetRouter() {
	api := u.Router.Group("/api/auth")
	{
		api.POST("/sign-up", u.UserController.SignUp)
		//api.POST("/sign-in", u.UserController.SignIn)
	}

}

func NewUserRoute(userController controller.IUserController, router *gin.Engine) *UserRoute {
	return &UserRoute{UserController: userController, Router: router}
}
