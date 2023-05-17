package repository

import "github.com/yomafleet/opensearch-client-test/pkg/domain/model"

type ProductRepositoryInterface interface {
	Create(*model.ProductCreateModel) error
	Search(string) ([]model.ProductGetModel, error)
}
