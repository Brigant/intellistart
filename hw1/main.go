package main

import "fmt"

var (
	priceApple     float64 = 5.99
	priсePear      int     = 7
	balance        int     = 24
	costApplePear  float64
	maxAmountPear  int
	maxAmountApple int
	canWeBuy       bool
)

func main() {
	costApplePear = 9*priceApple + 8*float64(priсePear)
	maxAmountPear = balance / priсePear
	maxAmountApple = int(float64(balance) / priceApple)
	if 2*priceApple+2*float64(priсePear) <= float64(balance) {
		canWeBuy = true
	}
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Printf("Потрібно %f грн.\n______________________\n", costApplePear)
	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Printf("Зможемо купи груш - %d шт.\n______________________\n", maxAmountPear)
	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Printf("Зможемо купи яблук - %d шт.\n______________________\n", maxAmountApple)
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Printf(" %v\n", canWeBuy)
}
