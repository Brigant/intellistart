package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FindTrains_successful_price(t *testing.T) {
	depStation := "1902"
	arrStation := "1929"
	criteria := "price"

	exp := Trains{
		{TrainID: 1177, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 164.65, ArrivalTime: time.Date(0, time.January, 1, 10, 25, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 16, 36, 0, 0, time.UTC)},
		{TrainID: 1178, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 164.65, ArrivalTime: time.Date(0, time.January, 1, 10, 25, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 16, 36, 0, 0, time.UTC)},
		{TrainID: 1141, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 176.77, ArrivalTime: time.Date(0, time.January, 1, 12, 15, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 16, 48, 0, 0, time.UTC)},
	}

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.NoError(t, err)
}

func Test_FindTrains_successful_arrival(t *testing.T) {
	depStation := "1902"
	arrStation := "1929"
	criteria := "arrival-time"

	exp := Trains{
		{TrainID: 978, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 258.53, ArrivalTime: time.Date(0, time.January, 1, 4, 15, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 13, 10, 0, 0, time.UTC)},
		{TrainID: 1316, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 209.73, ArrivalTime: time.Date(0, time.January, 1, 5, 55, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 13, 52, 0, 0, time.UTC)},
		{TrainID: 2201, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 280, ArrivalTime: time.Date(0, time.January, 1, 6, 15, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 14, 55, 0, 0, time.UTC)},
	}

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.NoError(t, err)
}

func Test_FindTrains_successful_departure(t *testing.T) {
	depStation := "1902"
	arrStation := "1929"
	criteria := "departure-time"

	exp := Trains{
		{TrainID: 1386, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 220.49, ArrivalTime: time.Date(0, time.January, 1, 8, 30, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 13, 3, 0, 0, time.UTC)},
		{TrainID: 978, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 258.53, ArrivalTime: time.Date(0, time.January, 1, 4, 15, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 13, 10, 0, 0, time.UTC)},
		{TrainID: 1316, DepartureStationID: 1902, ArrivalStationID: 1929, Price: 209.73, ArrivalTime: time.Date(0, time.January, 1, 5, 55, 0, 0, time.UTC), DepartureTime: time.Date(0, time.January, 1, 13, 52, 0, 0, time.UTC)},
	}

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.NoError(t, err)
}

func Test_FindTrains_wrong_criteria(t *testing.T) {
	depStation := "1902"
	arrStation := "1929"
	criteria := "awef"
	var exp Trains
	expErr := errors.New("unsupported criteria")

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_absent_depStationId(t *testing.T) {
	depStation := ""
	arrStation := "1929"
	criteria := "departure"
	var exp Trains
	expErr := errors.New("empty departure station")

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_absent_arrStation(t *testing.T) {
	depStation := "1902"
	arrStation := ""
	criteria := "departure"
	var exp Trains
	expErr := errors.New("empty arrival station")

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_wrong_depStation(t *testing.T) {
	depStation := "12"
	arrStation := "1929"
	criteria := "price"
	var exp Trains
	var expErr error

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_wrong_arrStation(t *testing.T) {
	depStation := "1902"
	arrStation := "11"
	criteria := "price"
	var exp Trains
	var expErr error

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_bad_arrStation_input(t *testing.T) {
	depStation := "1902"
	arrStation := "serg"
	criteria := "price"
	var exp Trains
	var expErr = errors.New("bad arrival station input")

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}

func Test_FindTrains_bad_depStation_input(t *testing.T) {
	depStation := "serg"
	arrStation := "1922"
	criteria := "price"
	var exp Trains
	var expErr = errors.New("bad departure station input")

	act, err := FindTrains(depStation, arrStation, criteria)
	assert.Equal(t, exp, act)
	assert.Equal(t, err, expErr)
}
