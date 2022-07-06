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

// to do func (f *Farm) addAnimal(animal Animals)
// to do func (f *Farm) removeAnimal(animal Animals)
// to do func (f *Farm) killAnimal(animal Animals) // = nil

// Animal entity
type Animal struct {
	Name          string
	weight        float64
	FoodPerweight float64
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
	return a.weight * a.FoodPerweight
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

func createDog(name string, weight float64, biteForce int) Dog {
	const (
		foodPerweight = 10
		minWeight     = 0.1
		typeAnimal    = "Dog"
	)

	if name == "" {
		name = typeAnimal
	}
	if weight < minWeight {
		weight = minWeight
	}
	return Dog{
		Animal: Animal{
			Name:          name,
			weight:        weight,
			FoodPerweight: foodPerweight,
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

func createCat(name string, weight float64, agility int) Cat {
	const (
		foodPerweight = 7
		minWeight     = 0.1
		typeAnimal    = "Cat"
	)

	if name == "" {
		name = typeAnimal
	}
	if weight < minWeight {
		weight = minWeight
	}

	return Cat{
		Animal: Animal{
			Name:          name,
			weight:        weight,
			FoodPerweight: foodPerweight,
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

func createCow(name string, weight float64, amountMilk float64) Cow {
	const (
		foodPerweight = 25
		minWeight     = 20
		typeAnimal    = "Cow"
	)

	if name == "" {
		name = typeAnimal
	}
	if weight < minWeight {
		weight = minWeight
	}

	return Cow{
		Animal: Animal{
			Name:          name,
			weight:        weight,
			FoodPerweight: foodPerweight,
		},
		AmountMilk: amountMilk,
	}
}

func (c *Cow) info() {
	c.showInfo()
	fmt.Printf("The cow gives %v liter of milk\n", c.AmountMilk)
}
