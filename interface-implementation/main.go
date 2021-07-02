package main

import "fmt"

type producer interface {
	drive()
	stop()
}

type car struct {
	modelName string
	speed     int
	fuelType  string
}

func (c car) drive() {
	fmt.Printf(" I am driveing %s at a speed of %d \n", c.modelName, c.speed)
	fmt.Printf("drrr....drrr \n")
	fmt.Printf("%s car fuel type is %s \n", c.modelName, c.fuelType)
}

func (c car) stop() {
	fmt.Printf("%s car is stopped....\n", c.modelName)
}

func testCar(p producer) {
	p.drive()
	p.stop()
}

func main() {
	maruthi := car{
		modelName: "breza",
		speed:     120,
		fuelType:  "petrol",
	}

	benz := car{
		modelName: "c-class",
		speed:     150,
		fuelType:  "petrol",
	}
	testCar(maruthi)
	testCar(benz)
	
}
