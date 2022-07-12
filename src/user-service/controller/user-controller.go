package controller

import (
	"chilindo/src/admin-service/helper"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	//if u.UserService.IsDuplicateEmail(newUser.Email) {
	//	ctx.JSON(http.StatusConflict, gin.H{
	//		"error": "email existed",
	//	})
	//	log.Println("SignUp: email existed", errDTO)
	//	ctx.Abort()
	//	return
	//}

	createdUser, errCreate := u.UserService.CreateUser(newUser)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreate.Error(),
		})
		log.Println("SignUp: email existed", errDTO)
		ctx.Abort()
		return
	}

	tokenString, errGenToken := u.jwtService.GenerateToken(createdUser.ID)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	createdUser.Token = tokenString
	ctx.JSON(http.StatusCreated, gin.H{"token": createdUser.Token})
}

func (u *UserController) SignIn(ctx *gin.Context) {
	var loginDTO *dto.UserLoginDTO

	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	user, errVerify := u.UserService.VerifyCredential(loginDTO)
	if errVerify != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": errVerify.Error(),
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		ctx.Abort()
		return
	}

	tokenString, errGenToken := u.jwtService.GenerateToken(user.ID)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
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
