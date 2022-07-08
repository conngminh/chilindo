package controller

import (
	"chilindo/src/user/entity"
	token "chilindo/src/user/jwt"
	"chilindo/src/user/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IUserController interface {
	//SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type UserController struct {
	UserService service.IUserService
	Token       *token.JWTClaim
}

func (u UserController) SignUp(c *gin.Context) {
	var userBody *entity.User
	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error to sign up",
		})
		log.Println("SignUp: Error ShouldBindJSON in package controller", err.Error())
		return
	}
	user, err := u.UserService.SignUp(userBody)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSONP(http.StatusOK, user)
}

func NewUserController(userController service.IUserService) *UserController {
	return &UserController{UserService: userController}
}
