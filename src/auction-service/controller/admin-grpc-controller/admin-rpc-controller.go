package admin_grpc_controller

import (
	pb "chilindo/src/pkg/pb/admin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAdminServiceController interface {
	CheckIsAuth(adminClient pb.AdminServiceClient) gin.HandlerFunc
}

type AdminServiceController struct {
}

func NewAdminServiceController() *AdminServiceController {
	return &AdminServiceController{}
}

func (a AdminServiceController) CheckIsAuth(adminClient pb.AdminServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			c.Abort()
			return
		}
		res, err := adminClient.CheckIsAuth(c, &pb.CheckIsAuthRequest{
			Token: token,
		})
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error call admin service",
			})
			c.Abort()
			return
		}
		if !(res.IsAuth) {
			c.JSON(http.StatusForbidden, gin.H{
				"Message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
