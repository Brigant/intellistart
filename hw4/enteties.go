package main

import (
	"fmt"
)

// Farm entity
type Farm struct {
	Animals []Animals
}

func createFarm(animals ...Animals) Farm {
	farm := new(Farm)
	farm.Animals = append(farm.Animals, animals...)
	return *farm
}

func (f *Farm) getFarmInfo() (tootalFood float64) {
	for _, animal := range f.Animals {
		animal.info()
		tootalFood = tootalFood + animal.howMuchFood()
	}

	return tootalFood
}

// Animal entity
type Animal struct {
	Name           string
	Wheight        float64
	FoodPerWheight float64
}

type Animals interface {
	canShowInfo
	canEat
}

type canEat interface {
	howMuchFood() float64
}

type canShowInfo interface {
	info()
}

func (a *Animal) howMuchFood() float64 {
	return a.Wheight * a.FoodPerWheight
}

func (a *Animal) showInfo() {
	fmt.Printf("\nThe pet name is \"%v\"\n", a.Name)
	fmt.Printf("%v - is needed in %v foods per month\n", a.Name, a.howMuchFood())
}

// variety of Animal entities
// Dog entity
type Dog struct {
	Animal
	BiteForce int
}

func createDog(name string, wheight float64, biteForce int) Dog {
	const foodPerWheight = 10
	return Dog{
		Animal: Animal{
			Name:           name,
			Wheight:        wheight,
			FoodPerWheight: foodPerWheight,
		},
		BiteForce: biteForce,
	}
}

func (d *Dog) info() {
	d.showInfo()
	fmt.Printf("The dog's bite force is %v\n", d.BiteForce)
}

// Cat entity
type Cat struct {
	Animal
	Agilaty int
}

func createCat(name string, wheight float64, agility int) Cat {
	const foodPerWheight = 7
	return Cat{
		Animal: Animal{
			Name:           name,
			Wheight:        wheight,
			FoodPerWheight: foodPerWheight,
		},
		Agilaty: agility,
	}
}

func (c *Cat) info() {
	c.showInfo()
	fmt.Printf("Agility of the cat is  %v \n", c.Agilaty)
}

// Cow entety
type Cow struct {
	Animal
	AmountMilk float64
}

func createCow(name string, wheight float64, amountMilk float64) Cow {
	const foodPerWheight = 25
	return Cow{
		Animal: Animal{
			Name:           name,
			Wheight:        wheight,
			FoodPerWheight: foodPerWheight,
		},
		AmountMilk: amountMilk,
	}
}

func (c *Cow) info() {
	c.showInfo()
	fmt.Printf("The cow gives %v liter of milk\n", c.AmountMilk)
}
