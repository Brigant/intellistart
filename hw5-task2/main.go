package main

import (
	"fmt"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація

	var total int
	c := make(chan int)
	for _, s := range n {
		go sum(s, c)
		total += <-c
	}
	close(c)
	fmt.Println(total)
}

func sum(s []int, c chan int) {
	var sum int
	for _, i := range s {
		sum += i
	}
	c <- sum
}
