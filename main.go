package main

import (
	"errors"
	"fmt"
	"sync"
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

	fmt.Printf("Started processing %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error truck did not load correctly %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error truck did not unload correctly %w", err)
	}

	fmt.Printf("Finished processing %+v\n", truck)
	return nil
}

// processFleet demonstrates concurrent processing of multiple trucks
func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			processTruck(t)
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully")
}
