package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}
func (f Fan) Oscillate() {
	fmt.Println("Rotating on base")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering up")
}
func (c CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func Use(appliance Appliance) {

	// YOUR CODE HERE: If the appliance is a
	// Fan, call its Oscillate method.
	// If the appliance is a CoffeePot, call
	// its Brew method.

	
	fmt.Println(appliance)
	appliance.TurnOn()
	cf, exists := appliance.(Fan)
	if exists {
		cf.Oscillate()
	} else {
		appliance.(CoffeePot).Brew() 
	}
}

func main() {
	Use(Fan("Windco Breeze"))
	Use(CoffeePot("LuxBrew"))
}
