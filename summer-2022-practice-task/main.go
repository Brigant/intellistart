package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"time"
)

type Trains []Train

type Train struct {
	TrainID            int
	DepartureStationID int
	ArrivalStationID   int
	Price              float32
	ArrivalTime        time.Time
	DepartureTime      time.Time
}

// умови завдання
var (
	errUnsupportedCriteria = errors.New("unsupported criteria")
	errEmptyDepStation     = errors.New("empty departure station")
	errEmptyArrStation     = errors.New("empty arrival station")
	errBadArrStation       = errors.New("bad arrival station input")
	errBadDepStation       = errors.New("bad departure station input")
)

const (
	criteriaPrice     = "price"
	criteriaArrTime   = "arrival-time"
	criteriaDepTime   = "departure-time"
	fileName          = "data.json"
	timeFormat        = "15:04:05"
	firsNaturalNumber = 1
	resultLimit       = 3
)

func main() {
	//	... запит даних від користувача
	var departureStation, arrivalStation, criteria string
	fmt.Println("Введіть номер станції відправлення")
	fmt.Scan(&departureStation)
	fmt.Println("Введіть номер станції прибуття")
	fmt.Scan(&arrivalStation)
	fmt.Println("Введіть критерій сортування (price, arrival-time, departure-time)")
	fmt.Scan(&criteria)

	result, err := FindTrains(departureStation, arrivalStation, criteria)

	//	... обробка помилки
	if err != nil {
		fmt.Println(err)
	}

	//	... друк result
	fmt.Println(result)

}

//
func FindTrains(departureStation, arrivalStation, criteria string) (Trains, error) {
	// ... код

	// валідації введенних даних
	err := validate(departureStation, arrivalStation, criteria)
	if err != nil {
		return nil, err
	}
	arrStation, err := strconv.Atoi(arrivalStation)
	if err != nil || arrStation < 1 {
		return nil, errBadArrStation
	}
	depStation, err := strconv.Atoi(departureStation)
	if err != nil || depStation < 1 {
		return nil, errBadDepStation
	}

	trains, _ := getData()
	result := Trains{}

	// пошук по станціям
	/* for _, v := range trains {
		if v.ArrivalStationID == arrStation && v.DepartureStationID == depStation {
			result = append(result, v)
		}
	} */

	for i := 0; i < len(trains); i++ {
		if trains[i].ArrivalStationID == arrStation && trains[i].DepartureStationID == depStation {
			result = append(result, trains[i])
		}
	}

	/* sort.Slice(result, func(i, j int) bool {
		return result[i].TrainID < result[j].TrainID
	}) */

	// сортування по критерію
	if criteria == criteriaPrice {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].Price < result[j].Price
		})
	}
	if criteria == criteriaDepTime {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].DepartureTime.Before(result[j].DepartureTime)
		})
	}
	if criteria == criteriaArrTime {
		sort.SliceStable(result, func(i, j int) bool {
			return result[i].ArrivalTime.Before(result[j].ArrivalTime)
		})
	}

	if len(result) == 0 {
		return nil, nil
	}
	if len(result) < resultLimit {
		return result, nil
	}
	return result[:resultLimit], nil // маєте повернути правильні значення
}

// перевірка введенних данних від користувача
func validate(departureStation, arrivalStation, criteria string) error {
	if departureStation == "" {
		return errEmptyDepStation
	}
	if arrivalStation == "" {
		return errEmptyArrStation
	}
	if criteria != criteriaPrice && criteria != criteriaDepTime && criteria != criteriaArrTime {
		return errUnsupportedCriteria
	}

	return nil
}

// зчитує данні з вказаного файлу и парсить їх в структуру Trains
func getData() (Trains, error) {
	type rout struct {
		TrainID            int     `json:"trainId"`
		DepartureStationID int     `json:"departureStationId"`
		ArrivalStationID   int     `json:"arrivalStationId"`
		Price              float32 `json:"price"`
		ArrivalTime        string  `json:"arrivalTime"`
		DepartureTime      string  `json:"departureTime"`
	}
	routs := []rout{}

	rawData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawData, &routs)
	if err != nil {
		return nil, err
	}

	t := Trains{}

	for _, i := range routs {
		arrTime, err := time.Parse(timeFormat, i.ArrivalTime)
		if err != nil {
			return nil, err
		}
		depTime, err := time.Parse(timeFormat, i.DepartureTime)
		if err != nil {
			return nil, err
		}
		train := Train{
			TrainID:            i.TrainID,
			DepartureStationID: i.DepartureStationID,
			ArrivalStationID:   i.ArrivalStationID,
			Price:              i.Price,
			ArrivalTime:        arrTime,
			DepartureTime:      depTime,
		}
		t = append(t, train)
	}
	return t, nil
}

// простодля зручності  огляду структури змінено Stringer
func (t Trains) String() string {
	var result string
	for i, val := range t {
		if i == len(t)-1 {
			result += fmt.Sprintf("%+v", val)
		} else {
			result += fmt.Sprintf("%+v\n", val)
		}
	}
	return result
}
