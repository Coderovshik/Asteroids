package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var MeteorSprites = MustLoadImages("assets/Meteors/*.png")

type Meteor struct {
	position      Vector
	movement      Vector
	rotation      float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[r.Intn(len(MeteorSprites))]

	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	radius := ScreenWidth / 2.0
	angle := rand.Float64() * 2 * math.Pi

	x := radius * math.Cos(angle)
	y := radius * math.Sin(angle)

	pos := Vector{
		X: target.X + x,
		Y: target.Y + y,
	}

	velocity := 0.25 + rand.Float64()*1.5

	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	rotationSpeed := -0.02 + rand.Float64()*0.04

	return &Meteor{
		position:      pos,
		movement:      movement,
		rotationSpeed: rotationSpeed,
		sprite:        sprite,
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(dest *ebiten.Image) {
	width := m.sprite.Bounds().Dx()
	height := m.sprite.Bounds().Dy()

	halfW := float64(width) / 2
	halfH := float64(height) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	dest.DrawImage(m.sprite, op)
}

func (m *Meteor) Collider() Rect {
	return Rect{
		X:      m.position.X,
		Y:      m.position.Y,
		Width:  float64(m.sprite.Bounds().Dx()),
		Height: float64(m.sprite.Bounds().Dy()),
	}
}
