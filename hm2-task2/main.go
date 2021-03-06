package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* На вхід подано стрінг з цілими числа, котри розділені пробілами.
Треба повернути найбільше та найменше число.
Наприклад:
input := "1 2 3 4 5" // повертає "5 1"
input := "1 9 3 4 -5" // повертає "9 -5"
Уточнення:
1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
2. В стрінгі завжди буде принаймні одне число.
3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
число). Найбільше число має бути першим.
4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
*/
func main() {
	input := "1 9 3 4 -5"
	var result string

	var (
		max  int32
		min  int32
		temp []int32
	)
	_, _, _, _ = max, min, temp, result

	// перетворює стринг в слайс
	for _, val := range strings.Split(input, " ") {
		val, err := strconv.ParseInt(val, 0, 32)
		if err != nil {
			fmt.Printf("Не можу конвертувати в int32")
			break
		}
		temp = append(temp, int32(val))
	}

	// змінюе дефолтні значення для max min щоб уникнути ситуації коли відʼємне більше чим 0
	max = temp[0]
	min = temp[0]

	// перебірає тимчасовий масив  и візначає максимальне і мінімальне знчення
	// і записує в відповідні змінні
	for _, val := range temp {
		fmt.Println(val)
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}

	}

	result = strconv.FormatInt(int64(max), 10) + " " + strconv.FormatInt(int64(min), 10)

	fmt.Println(result)
}
