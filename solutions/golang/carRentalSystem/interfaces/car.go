package interfaces

import "github.com/ashishps1/awesome-low-level-design/carRentalSystem/enums"

type Car interface {
	Registration() string
	PricePerDay() int
	Year() int
	Model() enums.Model
	Company() enums.Company
}
