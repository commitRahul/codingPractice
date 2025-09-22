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

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 100
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 50
	e.battery -= 10.0
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 5.0
	return nil
}

// processTruck handles loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck %+v\n", truck)
	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "N-1"}
	et := &ElectricTruck{id: "E-1", battery: 100.0}
	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error processing normal truck: %v\n", err)
	}

	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error processing electric truck: %v\n", err)
	}

	log.Printf("%+v\n", nt)
	log.Printf("%+v\n", et)
}
