package src

import (
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/enums"
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/interfaces"
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/src/paymentHandlers"
	"time"
)

type CarRental struct {
	Cars           []interfaces.Car
	Registrations  *RegistrationInventory
	PaymentHandler interfaces.PaymentHandler
}

// this can be done through singleton, please refer other packages like parking-lot for reference
func NewCarRental() *CarRental {
	return &CarRental{
		Cars: []interfaces.Car{
			NewConcreteCar("HR123AB", enums.MAHINDRA, enums.THAR, 2020, 1000),
			NewConcreteCar("DL123AB", enums.MAHINDRA, enums.THAR, 2024, 2000),
			//NewConcreteCar("HR234RT", enums.MAHINDRA, enums.THAR, 2020, 1000),
			//NewConcreteCar("DL567HJ", enums.MAHINDRA, enums.THAR, 2024, 4000),

			NewConcreteCar("CH123AB", enums.HYUNDAI, enums.CRETA, 2018, 500),
			NewConcreteCar("DL134AB", enums.HYUNDAI, enums.CRETA, 2022, 1000),
			NewConcreteCar("PB234RT", enums.MARUTI, enums.CIAZ, 2016, 200),
			NewConcreteCar("DL578HJ", enums.MARUTI, enums.CIAZ, 2019, 1500),
		},
		Registrations:  NewRegistrationInventory(),
		PaymentHandler: &paymentHandlers.UpiHandler{},
	}
}

func (c *CarRental) FindAvailableCars(company enums.Company, model enums.Model, startDate, endDate time.Time) []interfaces.Car {
	var result []interfaces.Car
	for _, car := range c.Cars {
		if car.Company() == company && car.Model() == model {
			if c.Registrations.IsCarAvailable(car.Registration(), startDate, endDate) {
				result = append(result, car)
			}
		}
	}
	return result
}

func (c *CarRental) BookCar(car interfaces.Car, customer *Customer, startDate, endDate time.Time) *Registration {
	reg := NewRegistration(customer.license, car, startDate, endDate)
	reg.TotalPrice = reg.CalculateTotalPrice()

	if !c.PaymentHandler.HandlePayment() {
		return nil
	}
	c.Registrations.AddRegistration(reg)
	return reg
}

func (c *CarRental) ChangePaymentHandler(handler interfaces.PaymentHandler) {
	c.PaymentHandler = handler
}
