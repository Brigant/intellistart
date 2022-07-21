package main

import (
	"errors"
	"fmt"
)

var (
	errEdible     = errors.New("тварина не їстівна")
	errLowWeight  = errors.New("надто низька вага")
	errAnimalType = errors.New("тип тварини не збігаеться")
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

func (f farm) getFarmInfo() (tootalFood float64, err error) {
	for _, animal := range f.animals {
		err := animal.info()
		if errors.Is(err, errEdible) {
			fmt.Printf("!попередження для %v: %v\n", animal, err)
			continue
		}
		if err != nil {
			return .0, fmt.Errorf("info() для %v :%w", animal, err)
		}
		tootalFood = tootalFood + animal.howMuchFood()
	}

	return tootalFood, nil
}

// animal entity
type animal struct {
	name          string
	weight        float64
	foodPerWeight float64
	typeAnimal    string
	edible        bool
	minWeight     float64
}

type animals interface {
	canShowInfo
	canEat
}

type canEat interface {
	howMuchFood() float64
}

type canShowInfo interface {
	info() error
}

func (a animal) howMuchFood() float64 {
	return a.weight * a.foodPerWeight
}

func (a animal) showInfo() error {
	if a.weight < a.minWeight {
		return errLowWeight
	}
	fmt.Printf("\nThe pet name is \"%v\"\n", a.name)
	fmt.Printf("Animal weight is %v\n", a.weight)
	fmt.Printf("%v - is needed in %v foods per month\n", a.name, a.howMuchFood())
	if !a.edible {
		return errEdible
	}

	return nil
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
		edible        = false
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
			typeAnimal:    typeanimal,
			edible:        edible,
			minWeight:     minWeight,
		},
		biteForce: biteForce,
	}
}

func (d dog) info() error {
	if d.typeAnimal != "dog" {
		return errAnimalType
	}
	err := d.showInfo()
	if err != nil {
		return fmt.Errorf("showInfo(): %w", err)
	}
	fmt.Printf("The dog's bite force is %v\n", d.biteForce)
	return nil
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
		edible        = false
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
			typeAnimal:    typeanimal,
			edible:        edible,
			minWeight:     minWeight,
		},
		agilaty: agility,
	}
}

func (c cat) info() error {

	err := c.showInfo()
	fmt.Printf("Agility of the cat is  %v \n", c.agilaty)
	if err != nil {
		return fmt.Errorf("showInfo(): %w", err)
	}

	if c.typeAnimal != "cat" {
		return errAnimalType
	}
	return nil
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
		edible        = true
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
			typeAnimal:    typeanimal,
			edible:        edible,
			minWeight:     minWeight,
		},
		amountMilk: amountMilk,
	}
}

func (c cow) info() error {
	if c.typeAnimal != "cow" {
		return errAnimalType
	}
	err := c.showInfo()
	if err != nil {
		return fmt.Errorf("showInfo(): %w", err)
	}
	fmt.Printf("The cow gives %v liter of milk\n", c.amountMilk)
	return nil
}
