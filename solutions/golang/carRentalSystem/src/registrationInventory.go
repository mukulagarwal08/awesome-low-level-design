package src

import (
	"time"
)

type RegistrationInventory struct {
	VehicleAvailability map[string]map[string]struct{}
	Registrations       map[string]*Registration
}

func NewRegistrationInventory() *RegistrationInventory {
	return &RegistrationInventory{
		VehicleAvailability: make(map[string]map[string]struct{}),
		Registrations:       make(map[string]*Registration),
	}
}

func (r *RegistrationInventory) IsCarAvailable(reg string, startDate, endDate time.Time) bool {
	for regId, _ := range r.VehicleAvailability[reg] {
		regObj := r.Registrations[regId]
		if startDate.Before(regObj.EndDate) && endDate.After(regObj.StartDate) {
			return false
		}
	}
	return true
}

func (r *RegistrationInventory) AddRegistration(reg *Registration) {
	r.Registrations[reg.Id] = reg

	regIdsForACar, isPresent := r.VehicleAvailability[reg.Vehicle.Registration()]

	if !isPresent {
		regIdsForACar = make(map[string]struct{})
	}

	regIdsForACar[reg.Id] = struct{}{}
	r.VehicleAvailability[reg.Vehicle.Registration()] = regIdsForACar
}

func (r *RegistrationInventory) DeleteRegistration(reg *Registration) {
	delete(r.VehicleAvailability[reg.Vehicle.Registration()], reg.Id)
	delete(r.Registrations, reg.Id)
}
