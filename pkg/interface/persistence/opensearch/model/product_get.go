package model

import (
	domainModel "github.com/yomafleet/opensearch-client-test/pkg/domain/model"
)

type ProductGetModel struct {
	ID     string            `json:"_id"`
	Score  float64           `json:"_score"`
	Source map[string]string `json:"_source"`
}

func (m ProductGetModel) ConvertToDomainModel() *domainModel.ProductGetModel {

	return domainModel.NewProductGetModel(
		m.ID,
		m.Source["brand"],
		m.Source["category"],
		m.Source["model"],
		m.Score)
}
