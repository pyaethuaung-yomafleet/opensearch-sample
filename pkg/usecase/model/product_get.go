package model

import "github.com/yomafleet/opensearch-client-test/pkg/domain/model"

type ProductGetModel struct {
	ID       string
	Brand    string
	Category string
	Model    string
	Score    float64
}

func ConvertToDomainModel(input *model.ProductGetModel) *ProductGetModel {

	return &ProductGetModel{

		ID:       input.ID(),
		Brand:    input.Brand(),
		Category: input.Category(),
		Model:    input.Model(),
		Score:    input.Score(),
	}
}
