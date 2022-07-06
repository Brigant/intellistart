package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	// мапа для визначення повторного елемента
	temp := make(map[int]bool)

	// перебірається вхідний слайс, і звіряється кожний єлемент на присутніст  в додатковій мапі
	for i := range arr {
		if !temp[arr[i]] {
			temp[arr[i]] = true
			result = append(result, arr[i])
		}
	}

	fmt.Println(result)
}
