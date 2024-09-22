package src

import "github.com/ashishps1/awesome-low-level-design/carRentalSystem/enums"

type ConcreteCar struct {
	registration string
	company      enums.Company
	model        enums.Model
	year         int
	price        int
}

func (c *ConcreteCar) Registration() string {
	return c.registration
}

func (c *ConcreteCar) Company() enums.Company {
	return c.company
}

func (c *ConcreteCar) Model() enums.Model {
	return c.model
}

func (c *ConcreteCar) Year() int {
	return c.year
}

func (c *ConcreteCar) PricePerDay() int {
	return c.price
}

func NewConcreteCar(reg string, company enums.Company, model enums.Model, year, price int) *ConcreteCar {
	return &ConcreteCar{
		registration: reg,
		company:      company,
		model:        model,
		year:         year,
		price:        price,
	}
}
