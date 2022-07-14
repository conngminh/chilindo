package dto

import "chilindo/src/product-service/entity"

type CreateOptionDTO struct {
	Option *entity.ProductOption
}

func NewCreateOptionDTO(option *entity.ProductOption) *CreateOptionDTO {
	return &CreateOptionDTO{Option: option}
}

type UpdateOptionDTO struct {
	Option *entity.ProductOption
}

func NewUpdateOptionDTO(option *entity.ProductOption) *UpdateOptionDTO {
	return &UpdateOptionDTO{Option: option}
}

type OptionIdDTO struct {
	OptionId int
}

type OptionByIdDTO struct {
	OptionId  int
	ProductId string
}
type ProductIdDTO struct {
	ProductId string
}
