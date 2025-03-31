package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math"
	"math/rand"
)

type Particle struct {
	position rl.Vector2
	velocity rl.Vector2
	color    color.RGBA
}

func NewParticle() Particle {
	pos := rl.Vector2{X: 1600 * rand.Float32(), Y: 1600 * rand.Float32()}
	vel := rl.Vector2{X: 0.0, Y: 0.0}
	col := color.RGBA{0, 0, 0, 5}

	new_particle := Particle{position: pos, velocity: vel, color: col}

	return new_particle
}

func (p *Particle) Reset() {
	pos := rl.Vector2{X: 1600 * rand.Float32(), Y: 1600 * rand.Float32()}
	vel := rl.Vector2{X: 0.0, Y: 0.0}
	p.position = pos
	p.velocity = vel
}

func (p *Particle) Draw() {
	rl.DrawCircleV(p.position, 1.5, p.color)
}

func (p *Particle) Update_Velocity(angle, dt float64) {

	force := rl.Vector2{X: float32(32 * math.Sin(angle)), Y: float32(-32 * math.Cos(angle))}
	p.velocity = rl.Vector2Scale(p.velocity, 0.9)
	p.velocity = rl.Vector2Add(p.velocity, rl.Vector2Scale(force, float32(dt)))
	//p.velocity = rl.Vector2Scale(force, float32(dt))
}

func (p *Particle) Update_Position(dt float64) {
	p.position = rl.Vector2Add(p.position, rl.Vector2Scale(p.velocity, float32(dt)))

	if (p.position.X < 0) || (p.position.Y < 0) {
		p.Reset()
	}
	if (p.position.X > 1600) || (p.position.Y > 1600) {
		p.Reset()
	}

}

func Process_Particles(ps []Particle, z float64) {
	for i := range ps {

		val := perlin_octaves3d(float64(ps[i].position.X), float64(ps[i].position.Y), z, 4)
		angle := 2 * math.Pi * val / 255
		ps[i].Update_Velocity(angle, 0.08)
		ps[i].Update_Position(0.08)
		ps[i].Draw()
	}
}
