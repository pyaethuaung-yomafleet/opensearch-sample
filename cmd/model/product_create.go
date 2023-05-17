package model

import (
	usecaseModel "github.com/yomafleet/opensearch-client-test/pkg/usecase/model"
)

type ProductCreateModel struct {
	ID       string `json:"id"`
	Brand    string `json:"brand"`
	Category string `json:"category"`
	Model    string `json:"model"`
}

func (m ProductCreateModel) ConvertToUsecaseModel() *usecaseModel.ProductCreateModel {

	return &usecaseModel.ProductCreateModel{

		ID:       m.ID,
		Brand:    m.Brand,
		Category: m.Category,
		Model:    m.Model,
	}
}
