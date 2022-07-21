package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	var total int
	c := make(chan int)

	wg.Add(len(n))
	for _, s := range n {
		go func(s []int, c chan int) {
			defer wg.Done()
			sum(s, c)

		}(s, c)
		total += <-c
	}
	wg.Wait()

	fmt.Println(total)
}

func sum(s []int, c chan int) {
	var sum int
	for _, i := range s {
		sum += i
	}
	c <- sum
}
