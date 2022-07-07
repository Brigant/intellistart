package main

import (
	"fmt"
	"strconv"
)

/*
Створити структури та реалізувати (імплементувати) методи:
a) метод для типу Rectangle, який намалює в консолі прямокутник літерами “O” (як на прикладі в кутку),
метод повинен давати вибір - малювати вертикально, або горизонтально.
b) метод, що збільшить розмір Прямокутника в задану кількість разів
c) на Прямокутнику зробити метод, що порівнює його площу із площею іншого Прямокутника
d) створити тип Квадрат, зробити на Прямокутнику метод, що приймає Квадрат і відповідає скільки =
таких Квадратів можна вписати у цей Прямокутник
e) Створити слайс Користувачів (структура як із лекції), вивести повне імʼя (fullName) і
нік користувачів зі слайса, які старші за 20 років, і імʼя (firstName) яких “John”.
*/

type Rectangle struct {
	Height int
	Width  int
}

type Square struct {
	Side int
}

type canDraw interface {
	drawIt()
}

func creatRectangle(heigt int, width int) Rectangle {
	return Rectangle{
		Height: heigt,
		Width:  width,
	}
}

func (r *Rectangle) showInfo() {
	fmt.Println("___________________")
	fmt.Printf("heigt: %v\nwidth: %v\n", r.Height, r.Width)
	fmt.Println("-------------------\n")
}

func (r *Rectangle) increaseRectangle(n int) {
	r.Height = r.Height * n
	r.Width = r.Width * n
}

func (r *Rectangle) drowIt(s string) {
	hline := ""
	vline := ""

	if s == "v" {
		for i := 0; i < r.Width; i++ {
			hline = hline + "o "
			if i == 0 || i == r.Width-1 {
				vline = vline + "o "
			} else {
				vline = vline + "  "
			}
		}

		for i := 0; i < r.Height; i++ {
			// fmt.Println(i)
			if i == 0 || i == r.Height-1 {
				fmt.Println(hline)
			} else {
				fmt.Println(vline)
			}
		}
	} else {
		for i := 0; i < r.Height; i++ {
			hline = hline + "o "
			if i == 0 || i == r.Height-1 {
				vline = vline + "o "
			} else {
				vline = vline + "  "
			}
		}

		for i := 0; i < r.Width; i++ {
			// fmt.Println(i)
			if i == 0 || i == r.Width-1 {
				fmt.Println(hline)
			} else {
				fmt.Println(vline)
			}
		}
	}
}

func (r *Rectangle) compare(other Rectangle) {
	sr := float64(r.Height) * float64(r.Width)
	or := float64(other.Height) * float64(other.Width)

	if sr < or {
		fmt.Printf("This rectangle area %v is less then other in %v times\n", int(sr), or/sr)
	}
	if sr > or {
		fmt.Printf("This rectangle area %v is greater then other in %v times\n", int(sr), sr/or)
	}
	if sr == or {
		fmt.Println("These rectangle  areas are equal")
	}
}

func creatSquare(s int) Square {
	return Square{
		Side: s,
	}
}

func (s *Square) drowIt() {
	hline := ""
	vline := ""

	for i := 0; i < s.Side; i++ {
		hline = hline + "o "
		if i == 0 || i == s.Side-1 {
			vline = vline + "o "
		} else {
			vline = vline + "  "
		}
	}

	for i := 0; i < s.Side; i++ {
		// fmt.Println(i)
		if i == 0 || i == s.Side-1 {
			fmt.Println(hline)
		} else {
			fmt.Println(vline)
		}
	}
}

func (r *Rectangle) howManySquares(s Square) {

	ch := r.Height / s.Side
	cv := r.Width / s.Side

	fmt.Printf("The rectanle containts %v squars\n", ch*cv)
}

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	nick      string
	age       int
}

func createNewUser(id, email, firstName, lastName string, age int) User {
	return User{
		id:        id,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func main() {
	/* 	r := creatRectangle(6, 6)
	   	q := creatSquare(3)

	   	r.howManySquares(q)

	   	r.drowIt("v")
	   	q.drowIt() */
	const (
		minAge    = 18
		maxAmount = 15
	)

	var UserList []User
	for i := 0; i < maxAmount; i++ {
		email := "example_" + strconv.FormatInt(int64(i), 10) + "@mail.com"
		firstName := ""
		lastName := ""

		firstName = string(70+i) + "oh" + string(106+(i))
		lastName = string(76+i) + "no" + string(100+(i))

		age := minAge + i

		user := createNewUser(strconv.FormatInt(int64(i), 10), email, firstName, lastName, age)
		UserList = append(UserList, user)
	}

	for _, user := range UserList {
		if user.firstName == "John" && user.age > 20 {
			fmt.Println(user)
		}
	}
}
