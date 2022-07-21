package main

import (
	"fmt"
)

func main() {

	// create animals
	Bobik := createDog("Bobik", 5, 3)
	Strelka := createDog("Strelka", 6, 3)
	Murzik := createCat("Murzik", 5, 5)
	Milka := createCow("Milka", 400, 15)

	// Тест на тип
	// Bobik.typeAnimal = "cake"

	// Тест на вагу
	// Milka.weight = 1

	// create farm by adding animals
	farm := createFarm(&Bobik, &Murzik, &Milka, &Strelka)
	totalFood, err := farm.getFarmInfo()

	if err != nil {
		fmt.Printf("не можу порахувати тому що: %v\n", err)
	}

	fmt.Println(totalFood)

}
