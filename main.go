package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func main() {
	fmt.Println("Perlin noise")

	fmt.Println(get_location_in_grid(0, 0))
	fmt.Println(get_gradient(0, 123))

	width := 1600
	height := 1600

	rl.InitWindow(int32(width), int32(height), "perlin noise")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	tot_time := 0.0
	speed := 0.5

	particles := make([]Particle, 0)

	for range 10000 {
		p := NewParticle()
		particles = append(particles, p)
	}

	rl.ClearBackground(rl.RayWhite)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		//		rl.DrawRectangle(0, 0, 1600, 1600, color.RGBA{255, 255, 255, 2})
		//		rl.ClearBackground(rl.RayWhite)
		tot_time += speed

		Process_Particles(particles, tot_time)

		rl.EndDrawing()
	}
}

func draw_noise(scaled_width, scaled_height, scale int) {
	min_val := 99999.0
	max_val := 0.0

	for ix := range scaled_width {
		for iy := range scaled_height {
			val := perlin_octaves3d(float64(ix*scale), float64(iy*scale), 0.0, 1)

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

}
