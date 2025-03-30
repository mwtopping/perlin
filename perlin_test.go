package main

import (
	"testing"
)

func Test_grid_location(t *testing.T) {
	gx, gy, rx, ry := get_location_in_grid(0, 0)
	if (gx != 0) || (gy != 0) || (rx != 0) || (ry != 0) {
		t.Errorf("Expected grid location of 0,0, got %v,%v", gx, gy)
	}

	gx, gy, _, _ = get_location_in_grid(99, 200)
	if (gx != 3) || (gy != 6) {
		t.Errorf("Expected grid location of 2,6, got %v,%v", gx, gy)
	}

}
