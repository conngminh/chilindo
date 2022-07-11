package controller

import (
	"chilindo/src/admin-service/dto"
	"chilindo/src/admin-service/helper"
	"chilindo/src/admin-service/service"
	"chilindo/src/user-service/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
)

type AdminController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(context *gin.Context)
}
type adminController struct {
	adminService service.AdminService
	jwtService   service.JWTService
}

func (a adminController) Login(ctx *gin.Context) {
	//TODO implement me
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := a.adminService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := a.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (a adminController) Register(ctx *gin.Context) {
	//TODO implement me
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !a.adminService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdAdmin := a.adminService.CreateAdmin(registerDTO)
		token := a.jwtService.GenerateToken(strconv.FormatUint(uint64(createdAdmin.Id), 10))
		createdAdmin.Token = token
		response := helper.BuildResponse(true, "OK!", createdAdmin)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (a adminController) Update(context *gin.Context) {
	//TODO implement me
	var adminUpdateDTO dto.AdminUpdateDTO
	errDTO := context.ShouldBind(&adminUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := a.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	adminUpdateDTO.ID = id
	u := a.adminService.Update(adminUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func NewAdminController(adminService service.AdminService, jwtService service.JWTService) AdminController {
	return &adminController{
		adminService: adminService,
		jwtService:   jwtService,
	}
}
