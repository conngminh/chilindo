package controller

import (
	"chilindo/src/user-service/entity"
	service "chilindo/src/user-service/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTestUser(t *testing.T) (*service.MockIUserService, *UserController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSrv := service.NewMockIUserService(ctr)
	userCtr := NewUserControllerDefault(mockSrv)
	return mockSrv, userCtr
}

func TestUserController_SignUp(t *testing.T) {
	mockSrv, userCtr := createTestUser(t)

	//mock service
	mockSrv.EXPECT().CreateUser(gomock.Any()).Return(&entity.User{
		Model:     gorm.Model{},
		Id:        0,
		Firstname: "",
		Lastname:  "",
		Password:  "",
		Birthday:  "",
		Phone:     "",
		Email:     "",
		Gender:    false,
		Country:   "",
		Language:  "",
		Token:     "",
	}, nil).Times(1)

	mockSrv.EXPECT().IsDuplicateEmail(gomock.Any()).Return(false)

	bodyRequest := `{"email":""}`

	req, err := http.NewRequest("POST", "chilindo/user/sign-up", strings.NewReader(bodyRequest))
	if err != nil {
		t.Fatalf("error %v", err)
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	userCtr.SignUp(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status is 201 but got %v", w.Code)
	}

} //done
