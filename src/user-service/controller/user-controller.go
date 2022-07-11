package controller

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUserController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type UserController struct {
	UserService service.IUserService
	jwtService  service.JWTService
}

func NewUserControllerDefault(userService service.IUserService, jwtService service.JWTService) *UserController {
	return &UserController{UserService: userService, jwtService: jwtService}
}

func (u *UserController) SignUp(ctx *gin.Context) {
	var newUser dto.LoginDTO
	errDTO := ctx.ShouldBind(&newUser)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(newUser.Email)
	if !u.UserService.IsDuplicateEmail(newUser.Email) {
		response := utils.BuildErrorResponse("Failed to process request", "email already existed", utils.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := u.UserService.CreateUser(newUser)
		token := u.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		response := utils.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
func (u *UserController) SignIn(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := utils.BuildErrorResponse("Failed to process request", errDTO.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := u.UserService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if authenticatedUser, ok := authResult.(entity.User); ok {
		generatedToken := u.jwtService.GenerateToken(strconv.FormatUint(uint64(authenticatedUser.ID), 10))
		authenticatedUser.Token = generatedToken
		response := utils.BuildResponse(true, "OK!", authenticatedUser)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := utils.BuildErrorResponse("Please check again your credential", "wrong email or password", utils.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
