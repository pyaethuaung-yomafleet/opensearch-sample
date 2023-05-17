package model

import (
	usecaseModel "github.com/yomafleet/opensearch-client-test/pkg/usecase/model"
)

type ProductGetModel struct {
	ID       string  `json:"id"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Model    string  `json:"model"`
	Score    float64 `json:"score"`
}

func ConvertFromUsecaseModel(input *usecaseModel.ProductGetModel) *ProductGetModel {

	return &ProductGetModel{

		ID:       input.ID,
		Brand:    input.Brand,
		Category: input.Category,
		Model:    input.Model,
		Score:    input.Score,
	}
}
