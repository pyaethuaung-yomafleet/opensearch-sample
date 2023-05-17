package usecase

import (
	"github.com/yomafleet/opensearch-client-test/pkg/domain/repository"
	"github.com/yomafleet/opensearch-client-test/pkg/usecase/model"
)

type ProductUsecase struct {
	productRepo repository.ProductRepositoryInterface
}

func (m ProductUsecase) Create(input *model.ProductCreateModel) error {

	if input == nil {

		return nil
	}

	return m.productRepo.Create(input.ConvertToDomainModel())
}

func (m ProductUsecase) Search(query string) ([]model.ProductGetModel, error) {

	products, err := m.productRepo.Search(query)
	if err != nil {

		return nil, err
	}

	var result []model.ProductGetModel
	for _, v := range products {

		if m := model.ConvertToDomainModel(&v); m != nil {

			result = append(result, *m)
		}
	}

	return result, nil
}

func NewProductUsecase(
	productRepo repository.ProductRepositoryInterface) *ProductUsecase {

	return &ProductUsecase{

		productRepo: productRepo,
	}
}
