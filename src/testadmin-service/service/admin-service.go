package service

import (
	"chilindo/src/admin-service/dto"
	"chilindo/src/admin-service/entity"
	"chilindo/src/admin-service/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AdminService interface {
	VerifyCredential(email string, password string) interface{}
	CreateAdmin(admin dto.RegisterDTO) entity.Administrator
	IsDuplicateEmail(email string) bool
	Update(admin dto.AdminUpdateDTO) entity.Administrator
}
type adminService struct {
	adminRepository repository.AdminRepository
}

func (service adminService) Update(admin dto.AdminUpdateDTO) entity.Administrator {
	//TODO implement me
	adminToUpdate := entity.Administrator{}
	err := smapping.FillStruct(&adminToUpdate, smapping.MapFields(&admin))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.adminRepository.UpdateAdmin(adminToUpdate)
	return updatedUser
}

func (service adminService) IsDuplicateEmail(email string) bool {
	//TODO implement me
	res := service.adminRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service adminService) VerifyCredential(email string, password string) interface{} {
	//TODO implement me
	res := service.adminRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.Administrator); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service adminService) CreateAdmin(admin dto.RegisterDTO) entity.Administrator {
	//TODO implement me
	userToCreate := entity.Administrator{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&admin))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.adminRepository.InsertAdmin(userToCreate)
	return res
}

func NewAdminService(adminRep repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository: adminRep,
	}
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
