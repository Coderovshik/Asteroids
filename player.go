package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = MustLoadImage("assets/playerShip1_blue.png")

type Player struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerSprite

	width := sprite.Bounds().Dx()
	height := sprite.Bounds().Dy()

	halfW := float64(width) / 2
	halfH := float64(height) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	speed = 5.0

	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.X = speed * math.Cos(math.Pi/2+p.rotation)
		delta.Y = speed * math.Sin(math.Pi/2+p.rotation)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.X = -speed * math.Cos(math.Pi/2+p.rotation)
		delta.Y = -speed * math.Sin(math.Pi/2+p.rotation)
	}

	p.position.X += delta.X
	p.position.Y += delta.Y
}

func (p *Player) Draw(dest *ebiten.Image) {
	width := p.sprite.Bounds().Dx()
	height := p.sprite.Bounds().Dy()

	halfW := float64(width) / 2
	halfH := float64(height) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	dest.DrawImage(p.sprite, op)
}

func (p *Player) Collider() Rect {
	return Rect{
		X:      p.position.X,
		Y:      p.position.Y,
		Width:  float64(p.sprite.Bounds().Dx()),
		Height: float64(p.sprite.Bounds().Dy()),
	}
}
