package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wgSum sync.WaitGroup
	wgSum.Add(len(n))
	for k, i := range n {

		go func(i []int, k int) {
			defer wgSum.Done()
			sum(i, k)
		}(i, k)
	}

	wgSum.Wait()
}

func sum(s []int, k int) {
	// Ваша реалізація
	var sum int
	for _, i := range s {
		sum += i
	}
	fmt.Printf("Slice %v: %v\n", k, sum)
}
