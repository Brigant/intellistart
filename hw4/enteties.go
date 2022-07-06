package main

import (
	"fmt"
)

// farm entity
type farm struct {
	animals []animals
}

func createFarm(animals ...animals) farm {
	farm := farm{}
	farm.animals = append(farm.animals, animals...)
	return farm
}

func (f farm) getFarmInfo() (tootalFood float64) {
	for _, animal := range f.animals {
		animal.info()
		tootalFood = tootalFood + animal.howMuchFood()
	}

	return tootalFood
}

// animal entity
type animal struct {
	name          string
	weight        float64
	foodPerWeight float64
}

type animals interface {
	canShowInfo
	canEat
}

type canEat interface {
	howMuchFood() float64
}

type canShowInfo interface {
	info()
}

func (a animal) howMuchFood() float64 {
	return a.weight * a.foodPerWeight
}

func (a animal) showInfo() {
	fmt.Printf("\nThe pet name is \"%v\"\n", a.name)
	fmt.Printf("%v - is needed in %v foods per month\n", a.name, a.howMuchFood())
}

// variety of animal entities
// dog entity
type dog struct {
	animal
	biteForce int
}

func createDog(name string, weight float64, biteForce int) dog {
	const (
		foodPerWeight = 10
		minWeight     = 0.1
		typeanimal    = "dog"
	)

	if name == "" {
		name = typeanimal
	}
	if weight < minWeight {
		weight = minWeight
	}
	return dog{
		animal: animal{
			name:          name,
			weight:        weight,
			foodPerWeight: foodPerWeight,
		},
		biteForce: biteForce,
	}
}

func (d dog) info() {
	d.showInfo()
	fmt.Printf("The dog's bite force is %v\n", d.biteForce)
}

// cat entity
type cat struct {
	animal
	agilaty int
}

func createCat(name string, weight float64, agility int) cat {
	const (
		foodPerWeight = 7
		minWeight     = 0.1
		typeanimal    = "cat"
	)

	if name == "" {
		name = typeanimal
	}
	if weight < minWeight {
		weight = minWeight
	}

	return cat{
		animal: animal{
			name:          name,
			weight:        weight,
			foodPerWeight: foodPerWeight,
		},
		agilaty: agility,
	}
}

func (c cat) info() {
	c.showInfo()
	fmt.Printf("Agility of the cat is  %v \n", c.agilaty)
}

// cow entety
type cow struct {
	animal
	amountMilk float64
}

func createCow(name string, weight float64, amountMilk float64) cow {
	const (
		foodPerWeight = 25
		minWeight     = 20
		typeanimal    = "cow"
	)

	if name == "" {
		name = typeanimal
	}
	if weight < minWeight {
		weight = minWeight
	}

	return cow{
		animal: animal{
			name:          name,
			weight:        weight,
			foodPerWeight: foodPerWeight,
		},
		amountMilk: amountMilk,
	}
}

func (c cow) info() {
	c.showInfo()
	fmt.Printf("The cow gives %v liter of milk\n", c.amountMilk)
}
