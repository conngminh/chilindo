package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
)

type ProductOptionService interface {
	CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error)
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionByIdDTO) (*entity.ProductOption, error)
	DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error)
}
