package model

import (
	domainModel "github.com/yomafleet/opensearch-client-test/pkg/domain/model"
)

type ProductCreateModel struct {
	Brand    string `json:"brand"`
	Category string `json:"category"`
	Model    string `json:"model"`
}

func ConvertFromDomainModel(
	input *domainModel.ProductCreateModel) *ProductCreateModel {

	return &ProductCreateModel{

		Brand:    input.Brand(),
		Category: input.Category(),
		Model:    input.Model(),
	}
}
