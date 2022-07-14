package dto

import "chilindo/src/product-service/entity"

type CreateOptionDTO struct {
	Option *entity.ProductOption
}

func NewCreateOptionDTO(option *entity.ProductOption) *CreateOptionDTO {
	return &CreateOptionDTO{Option: option}
}

type OptionIdDTO struct {
	OptionId string
}

type ProductIdDTO struct {
	ProductId string
}
