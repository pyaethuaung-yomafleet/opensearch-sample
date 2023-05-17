package model

type ProductCreateModel struct {
	id       string
	brand    string
	category string
	model    string
}

func (m ProductCreateModel) ID() string {

	return m.id
}

func (m ProductCreateModel) Brand() string {

	return m.brand
}

func (m ProductCreateModel) Category() string {

	return m.category
}

func (m ProductCreateModel) Model() string {

	return m.model
}

func NewProductCreateModel(id, brand, category, model string) *ProductCreateModel {

	return &ProductCreateModel{

		id:       id,
		brand:    brand,
		category: category,
		model:    model,
	}
}
