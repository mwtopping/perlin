package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func get_location_in_grid(x, y float64) (a, b, c, d float64) {
	gridsize := 400.0

	// get the grid index
	gridx := math.Floor(x / gridsize)
	gridy := math.Floor(y / gridsize)
	x_rem := math.Mod(x, gridsize) / gridsize
	y_rem := math.Mod(y, gridsize) / gridsize

	return gridx, gridy, x_rem, y_rem
}

func perlin(x, y float64) float64 {

	gx, gy, rx, ry := get_location_in_grid(x, y)
	location := rl.Vector2{float32(rx), float32(ry)}

	g00 := get_gradient(gx, gy)
	g10 := get_gradient(gx+1, gy)
	g01 := get_gradient(gx, gy+1)
	g11 := get_gradient(gx+1, gy+1)

	// distance within the grid
	d00 := rl.Vector2{-location.X, -location.Y}
	d10 := rl.Vector2{1 - location.X, -location.Y}
	d01 := rl.Vector2{-location.X, 1 - location.Y}
	d11 := rl.Vector2{1 - location.X, 1 - location.Y}

	// dot products
	dot00 := rl.Vector2DotProduct(g00, d00)
	dot10 := rl.Vector2DotProduct(g10, d10)
	dot01 := rl.Vector2DotProduct(g01, d01)
	dot11 := rl.Vector2DotProduct(g11, d11)

	u := smooth(rx)
	v := smooth(ry)

	x1 := lerp(dot00, dot10, u)
	x2 := lerp(dot01, dot11, u)

	val := lerp(x1, x2, v)

	return float64(val)
}

func lerp(x1, x2 float32, y float64) float32 {
	return float32(1-y)*x1 + (float32(y))*x2
}

func smooth(val float64) float64 {
	return 6*math.Pow(val, 5) - 15*math.Pow(val, 4) + 10*math.Pow(val, 3)
}

func get_gradient(x, y float64) rl.Vector2 {

	// ensure values are correct
	ix := int32(x)
	iy := int32(y)

	var hash int32 = ix
	hash ^= iy * 1609587929
	hash = hash*9650029 + 7
	hash = (hash >> 13) ^ hash

	angle := float64(hash&7) * math.Pi / 4.0
	return rl.Vector2{
		X: float32(math.Cos(angle)),
		Y: float32(math.Sin(angle)),
	}

	// xdir := 0
	// ydir := 0
	//
	//	if hash&1 == 0 {
	//		xdir = 1
	//	} else {
	//
	//		xdir = -1
	//	}
	//
	//	if hash&2 == 0 {
	//		ydir = 1
	//	} else {
	//
	//		ydir = -1
	//	}
	//
	// return rl.Vector2{X: float32(xdir), Y: float32(ydir)}
}

func perlin_octaves(x, y float64, n int) float64 {
	val := 0.0
	freq := 1.0
	amplitude := 1.0
	tot_amp := 0.0
	for range n {
		//val += (250 * (perlin(x*math.Pow(float64(1+i), 2),
		//y*math.Pow(float64(1+i), 2)) + 2) / 4) * (1 / math.Pow(float64(1+i), 2))
		val += amplitude * (0.5 + 0.5*(perlin(x*freq, y*freq)/0.60722))
		tot_amp += amplitude
		freq *= 2
		amplitude /= 2
	}

	return 200 * val / tot_amp
}
