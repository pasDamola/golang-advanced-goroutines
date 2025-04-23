package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Always rely on abstractions and not concrete implementations, this is why you use interfaces

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.battery = -1
	e.cargo += 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.battery += -1
	e.cargo = 0
	return nil
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error truck did not load correctly %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error truck did not unload correctly %w", err)
	}
	return nil
}

func main() {
	truckID := 42
	anotherTruckID := &truckID

	log.Println(anotherTruckID)

	truckID = 0
	log.Println(*anotherTruckID)

	// until values are assigned, the memory address will be nil
	var userID *int
	var userName *string
	var userAge *float64
	log.Println(userID, userName, userAge)
}
