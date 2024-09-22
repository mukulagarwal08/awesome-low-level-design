package main

import (
	"fmt"
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/enums"
	"github.com/ashishps1/awesome-low-level-design/carRentalSystem/src"
	"time"
)

func main() {
	carRental := src.NewCarRental()

	alice := src.NewCustomer("Alice", 12345678, "UEFUIV372UFJ")
	bob := src.NewCustomer("Bob", 23456778, "DJUVNIWER8B")

	cars := carRental.FindAvailableCars(enums.MAHINDRA, enums.THAR, time.Now().Add(time.Hour*24*10), time.Now().Add(time.Hour*24*12))

	if len(cars) > 0 {
		car := cars[0] //by default choosing first car
		reg := carRental.BookCar(car, alice, time.Now().Add(time.Hour*24*10), time.Now().Add(time.Hour*24*12))
		if reg == nil {
			fmt.Println("Payment failed")
		} else {
			reg.Print()
		}
	}

	// again requesting mahindra thar with overlapping period, different car should come
	cars = carRental.FindAvailableCars(enums.MAHINDRA, enums.THAR, time.Now().Add(time.Hour*24*9), time.Now().Add(time.Hour*24*11))

	if len(cars) > 0 {
		car := cars[0] //by default choosing first car
		reg := carRental.BookCar(car, bob, time.Now().Add(time.Hour*24*9), time.Now().Add(time.Hour*24*13))
		if reg == nil {
			fmt.Println("Payment failed")
		} else {
			reg.Print()
		}
	}

	// again requesting mahindra thar with overlapping period, no car should come as 2 thars were added
	cars = carRental.FindAvailableCars(enums.MAHINDRA, enums.THAR, time.Now().Add(time.Hour*24*8), time.Now().Add(time.Hour*24*12))

	if len(cars) > 0 {
		car := cars[0] //by default choosing first car
		reg := carRental.BookCar(car, bob, time.Now().Add(time.Hour*24*10), time.Now().Add(time.Hour*24*12))
		if reg == nil {
			fmt.Println("Payment failed")
		} else {
			reg.Print()
		}
	} else {
		fmt.Println("No car found for provided filters")
	}
}
