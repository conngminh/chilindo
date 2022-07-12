package controller

import (
	"chilindo/src/admin-service/helper"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strconv"
)

type IUserController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	Update(c *gin.Context)
}

type UserController struct {
	UserService service.IUserService
	jwtService  service.JWTService
}

func NewUserControllerDefault(userService service.IUserService, jwtService service.JWTService) *UserController {
	return &UserController{UserService: userService, jwtService: jwtService}
}

func (u UserController) SignUp(ctx *gin.Context) {
	var newUser *entity.User
	errDTO := ctx.ShouldBindJSON(&newUser)

	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if u.UserService.IsDuplicateEmail(newUser.Email) {
		response := utils.BuildErrorResponse("Failed to process request", "email already existed", utils.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		ctx.Abort()
		return
	}
	createdUser := u.UserService.CreateUser(newUser)
	tokenString, errGenToken := u.jwtService.GenerateToken(createdUser.ID)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error SignIn",
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	createdUser.Token = tokenString
	response := utils.BuildResponse(true, "OK!", createdUser)
	ctx.JSON(http.StatusCreated, response)

}
func (u *UserController) SignIn(ctx *gin.Context) {
	var loginDTO *dto.UserLoginDTO
	errDTO := ctx.ShouldBindJSON(&loginDTO)

	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to sign in",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	user, err := u.UserService.VerifyCredential(loginDTO)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": "Error SignIn",
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		ctx.Abort()
		return
	}

	tokenString, errGenToken := u.jwtService.GenerateToken(user.ID)

	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error SignIn",
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}

func (u *UserController) Update(context *gin.Context) {
	var userUpdateDTO *dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := u.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	user := u.UserService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
