package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var BulletSprite = MustLoadImage("assets/laserBlue04.png")

type Bullet struct {
	position Vector
	movement Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewBullet(player *Player) *Bullet {
	playerSprite := player.sprite

	width := playerSprite.Bounds().Dx()
	height := playerSprite.Bounds().Dy()

	halfW := float64(width) / 2
	halfH := float64(height) / 2

	offset := 50.0

	bulletSprite := BulletSprite

	pos := Vector{
		player.position.X + halfW -
			float64(bulletSprite.Bounds().Dx())/2 + math.Sin(player.rotation)*offset,
		player.position.Y + halfH -
			float64(bulletSprite.Bounds().Dy())/2 + math.Cos(player.rotation)*(-offset),
	}

	speed := 7.0

	movement := Vector{
		X: -speed * math.Cos(math.Pi/2+player.rotation),
		Y: -speed * math.Sin(math.Pi/2+player.rotation),
	}

	return &Bullet{
		position: pos,
		movement: movement,
		rotation: player.rotation,
		sprite:   bulletSprite,
	}
}

func (b *Bullet) Update() {
	b.position.X += b.movement.X
	b.position.Y += b.movement.Y
}

func (b *Bullet) Draw(dest *ebiten.Image) {
	width := b.sprite.Bounds().Dx()
	height := b.sprite.Bounds().Dy()

	halfW := float64(width) / 2
	halfH := float64(height) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	dest.DrawImage(b.sprite, op)
}

func (b *Bullet) Collider() Rect {
	return Rect{
		X:      b.position.X,
		Y:      b.position.Y,
		Width:  float64(b.sprite.Bounds().Dx()),
		Height: float64(b.sprite.Bounds().Dy()),
	}
}
