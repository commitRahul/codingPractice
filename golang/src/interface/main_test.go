package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload cargo for a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "N-1", cargo: 400}
			et := &ElectricTruck{id: "E-1", battery: 100.0}
			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error processing normal truck: %v\n", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing electric truck: %v\n", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatalf("Expected normal truck cargo to be 0, got %d", nt.cargo)
			}
			if et.cargo != 0 {
				t.Fatalf("Expected electric truck cargo to be 0, got %d", et.cargo)
			}
			if et.battery != 85.0 {
				t.Fatalf("Expected electric truck battery to be 85.0, got %f", et.battery)
			}
		})
	})
}
