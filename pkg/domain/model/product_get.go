package model

type ProductGetModel struct {
	id       string
	brand    string
	category string
	model    string
	score    float64
}

func (m ProductGetModel) ID() string {

	return m.id
}

func (m ProductGetModel) Brand() string {

	return m.brand
}

func (m ProductGetModel) Category() string {

	return m.category
}

func (m ProductGetModel) Model() string {

	return m.model
}

func (m ProductGetModel) Score() float64 {

	return m.score
}

func NewProductGetModel(
	id, brand, category, model string, score float64) *ProductGetModel {

	return &ProductGetModel{

		id:       id,
		brand:    brand,
		category: category,
		model:    model,
		score:    score,
	}
}
