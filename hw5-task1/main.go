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
	for _, i := range n {

		go func(i []int) {
			defer wgSum.Done()
			sum(i)
		}(i)
	}

	wgSum.Wait()
}

func sum(s []int) {
	// Ваша реалізація
	sum := 0
	for _, i := range s {
		sum += i
	}
	fmt.Println(sum)
}
