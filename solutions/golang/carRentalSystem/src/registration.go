package src

import (
	"fmt"
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/interfaces"
	"github.com/google/uuid"
	"time"
)

type Registration struct {
	Id              string
	CustomerLicense string
	Vehicle         interfaces.Car
	StartDate       time.Time
	EndDate         time.Time
	TotalPrice      int
}

func NewRegistration(customerLicense string, car interfaces.Car, startDate, endDate time.Time) *Registration {
	return &Registration{
		Id:              generateUniqueId(),
		CustomerLicense: customerLicense,
		Vehicle:         car,
		StartDate:       startDate,
		EndDate:         endDate,
	}
}

func (r *Registration) Print() {
	fmt.Printf("%s is booked by license no %s from %s to %s for Rs.%d\n", r.Vehicle.Registration(), r.CustomerLicense,
		r.StartDate.Format(time.RFC3339), r.EndDate.Format(time.RFC3339), r.TotalPrice)
}

func (r *Registration) CalculateTotalPrice() int {
	return (int)(r.EndDate.Sub(r.StartDate).Hours()) / 24 * r.Vehicle.PricePerDay()
}
func generateUniqueId() string {
	return uuid.New().String()
}
