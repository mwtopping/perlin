package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func get_location_in_grid3d(x, y, z float64) (a, b, c, d, e, f float64) {
	gridsize := 400.0

	// get the grid index
	gridx := math.Floor(x / gridsize)
	gridy := math.Floor(y / gridsize)
	gridz := math.Floor(z / gridsize)
	x_rem := math.Mod(x, gridsize) / gridsize
	y_rem := math.Mod(y, gridsize) / gridsize
	z_rem := math.Mod(z, gridsize) / gridsize

	return gridx, gridy, gridz, x_rem, y_rem, z_rem
}

func perlin3d(x, y, z float64) float64 {

	gx, gy, gz, rx, ry, rz := get_location_in_grid3d(x, y, z)
	location := rl.Vector3{float32(rx), float32(ry), float32(rz)}

	g000 := get_gradient3d(gx, gy, gz)
	g100 := get_gradient3d(gx+1, gy, gz)
	g010 := get_gradient3d(gx, gy+1, gz)
	g110 := get_gradient3d(gx+1, gy+1, gz)
	g001 := get_gradient3d(gx, gy, gz+1)
	g101 := get_gradient3d(gx+1, gy, gz+1)
	g011 := get_gradient3d(gx, gy+1, gz+1)
	g111 := get_gradient3d(gx+1, gy+1, gz+1)

	// distance within the grid
	d000 := rl.Vector3{-location.X, -location.Y, -location.Z}
	d100 := rl.Vector3{1 - location.X, -location.Y, -location.Z}
	d010 := rl.Vector3{-location.X, 1 - location.Y, -location.Z}
	d110 := rl.Vector3{1 - location.X, 1 - location.Y, -location.Z}
	d001 := rl.Vector3{-location.X, -location.Y, 1 - location.Z}
	d101 := rl.Vector3{1 - location.X, -location.Y, 1 - location.Z}
	d011 := rl.Vector3{-location.X, 1 - location.Y, 1 - location.Z}
	d111 := rl.Vector3{1 - location.X, 1 - location.Y, 1 - location.Z}

	// dot products
	dot000 := rl.Vector3DotProduct(g000, d000)
	dot100 := rl.Vector3DotProduct(g100, d100)
	dot010 := rl.Vector3DotProduct(g010, d010)
	dot110 := rl.Vector3DotProduct(g110, d110)
	dot001 := rl.Vector3DotProduct(g001, d001)
	dot101 := rl.Vector3DotProduct(g101, d101)
	dot011 := rl.Vector3DotProduct(g011, d011)
	dot111 := rl.Vector3DotProduct(g111, d111)

	u := smooth(rx)
	v := smooth(ry)
	w := smooth(rz)

	x1 := lerp(dot000, dot100, u)
	x2 := lerp(dot010, dot110, u)
	x3 := lerp(dot001, dot101, u)
	x4 := lerp(dot011, dot111, u)

	y1 := lerp(x1, x2, v)
	y2 := lerp(x3, x4, v)

	val := lerp(y1, y2, w)

	return float64(val)
}

func get_gradient3d(x, y, z float64) rl.Vector3 {
	ix := int32(x)
	iy := int32(y)
	iz := int32(z)

	var hash int32 = ix
	hash ^= iy * 1609587929
	hash ^= iz * 2048419325
	hash = hash*9650029 + 7
	hash = (hash >> 13) ^ hash

	theta := float64(hash&15) * math.Pi / 8.0   // azimuth
	phi := float64((hash>>4)&7) * math.Pi / 4.0 // altitude

	sinPhi := math.Sin(phi)
	return rl.Vector3{
		X: float32(math.Cos(theta) * sinPhi),
		Y: float32(math.Sin(theta) * sinPhi),
		Z: float32(math.Cos(phi)),
	}
}

func perlin_octaves3d(x, y, z float64, n int) float64 {
	val := 0.0
	freq := 1.0
	amplitude := 1.0
	tot_amp := 0.0
	for range n {
		val += amplitude * (0.5 + 0.5*(perlin3d(x*freq, y*freq, z*freq)/0.60722))
		tot_amp += amplitude
		freq *= 2
		amplitude /= 2
	}

	return 200 * val / tot_amp
}
