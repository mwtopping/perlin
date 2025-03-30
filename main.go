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

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		my_string := fmt.Sprintf("FPS: %v", rl.GetFPS())

		//		for ix := range 1600 / 32 {
		//			for iy := range 1600 / 32 {
		//				here := rl.Vector2{X: float32(32 * ix), Y: float32(32 * iy)}
		//				vec := rl.Vector2Scale(get_gradient(float64(32*ix), float64(32*iy)), 12)
		//				vec = rl.Vector2Add(here, vec)
		//
		//				rl.DrawLineV(here, vec, rl.DarkGray)
		//				rl.DrawCircleV(here, 2.0, rl.DarkGray)
		//			}
		//		}
		for ix := range scaled_width {
			for iy := range scaled_height {
				val := 250 * (perlin(float64(ix*scale), float64(iy*scale)) + 2) / 4
				//				fmt.Println(val)
				col := color.RGBA{uint8(val), uint8(val), uint8(val), 255}
				//				rl.DrawPixel(int32(ix), int32(iy), col)
				rl.DrawRectangle(int32(ix*scale), int32(iy*scale),
					int32(scale), int32(scale), col)

			}
		}
		rl.DrawText(my_string, 40, 40, 32, rl.Black)
		rl.EndDrawing()
	}
}
