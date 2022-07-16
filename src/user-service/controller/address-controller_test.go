package controller

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	service "chilindo/src/user-service/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateTestAddress(t *testing.T) (*service.MockIAddressService, *AddressController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvr := service.NewMockIAddressService(ctr)
	userCtr := NewAddressControllerDefault(mockSvr)
	return mockSvr, userCtr
}
func TestAddressController_CreateAddress(t *testing.T) {

}
func TestAddressController_GetAddress(t *testing.T) {
	mockSvr, userCtr := CreateTestAddress(t)
	mockSvr.EXPECT().GetAddress(gomock.Any()).Return(&[]entity.Address{{
		Model:       gorm.Model{},
		Firstname:   "",
		Lastname:    "",
		Phone:       "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		UserId:      0,
		User:        entity.User{},
	}}, nil).Times(1)

	req, err := http.NewRequest("GET", "chilindo/user/address/address", nil)

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set(config.UserId, uint(1))

	userCtr.GetAddress(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
} //done
