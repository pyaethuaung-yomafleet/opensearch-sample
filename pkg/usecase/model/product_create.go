package model

import "github.com/yomafleet/opensearch-client-test/pkg/domain/model"

type ProductCreateModel struct {
	ID       string
	Brand    string
	Category string
	Model    string
}

func (m ProductCreateModel) ConvertToDomainModel() *model.ProductCreateModel {

	return model.NewProductCreateModel(m.ID, m.Brand, m.Category, m.Model)
}
