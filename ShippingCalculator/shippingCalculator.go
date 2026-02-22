//خیلی خوب و تمیز نوشتی, ممنون امین جان
package main

import "fmt"

type ShippingCalculator interface {
	CalculateCost(weight int64, distance int64) int64
}
type NormalShipping struct {
	name             string
	pricePerKilogram int64
}

func (n NormalShipping) CalculateCost(weight int64, distance int64) int64 {
	return weight * n.pricePerKilogram * distance
}

type ExpressShipping struct {
	name             string
	pricePerKilogram int64
}

func (e ExpressShipping) CalculateCost(weight int64, distance int64) int64 {
	return weight * e.pricePerKilogram * distance
}

func DisplayShippingCost(shippingCalculator ShippingCalculator, weight int64, distance int64) {
	switch shippingCalculator.(type) {
	case NormalShipping:
		fmt.Println(fmt.Sprintf("the normal shipping price is %v", shippingCalculator.CalculateCost(weight, distance)))
	case ExpressShipping:
		fmt.Println(fmt.Sprintf("the express shipping price is %v", shippingCalculator.CalculateCost(weight, distance)))
	}
}

func main() {
	var weight, distance int64 = 5, 2
	shippings := []ShippingCalculator{
		NormalShipping{
			name:             "Normal Shipping",
			pricePerKilogram: 10,
		},
		ExpressShipping{
			name:             "Express Shipping",
			pricePerKilogram: 12,
		},
	}

	for _, shipping := range shippings {
		DisplayShippingCost(shipping, weight, distance)
	}

}
