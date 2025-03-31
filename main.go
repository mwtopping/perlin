package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	// "math"
)

func main() {
	fmt.Println("Perlin noise")

	fmt.Println(get_location_in_grid(0, 0))
	fmt.Println(get_gradient(0, 123))

	width := 1600
	height := 1600

	rl.InitWindow(int32(width), int32(height), "perlin noise")
	defer rl.CloseWindow()

	scale := 8
	scaled_width := width / scale
	scaled_height := height / scale

	rl.SetTargetFPS(15)
	min_val := 99999.0
	max_val := 0.0

	tot_time := 0.0
	speed := 2.5

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		my_string := fmt.Sprintf("FPS: %v, TIME: %v", rl.GetFPS(), tot_time)
		tot_time += speed
		for ix := range scaled_width {
			for iy := range scaled_height {
				val := perlin_octaves3d(float64(ix*scale), float64(iy*scale), tot_time, 4)

				if val < min_val {
					min_val = val
				}
				if val > max_val {
					max_val = val
				}
				col := color.RGBA{uint8(val), uint8(val), uint8(val), 255}
				rl.DrawRectangle(int32(ix*scale), int32(iy*scale),
					int32(scale), int32(scale), col)

			}
		}
		rl.DrawText(my_string, 40, 40, 32, rl.Black)
		rl.EndDrawing()
	}
	fmt.Println(min_val, max_val)
}
