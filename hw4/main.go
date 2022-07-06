package main

import "fmt"

func main() {

	// create animals
	Bobik := createDog("Bobik", 5, 3)
	Murzik := createCat("Murzik", 5, 5)
	Milka := createCow("Milka", 400, 15)

	// create farm by adding animals
	farm := createFarm(&Bobik, &Murzik, &Milka)
	totalFood := farm.getFarmInfo()

	fmt.Println(totalFood)

}
