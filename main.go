package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player           *Player
	shootCooldown    *Timer
	bullets          []*Bullet
	meteorSpawnTimer *Timer
	meteors          []*Meteor
}

func (g *Game) Update() error {
	g.player.Update()

	g.shootCooldown.Update()
	if g.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.shootCooldown.Reset()

		b := NewBullet(g.player)
		g.bullets = append(g.bullets, b)
	}

	for _, b := range g.bullets {
		b.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
			}
		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bullets {
		b.Draw(screen)
	}

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Reset() {
	g.player = NewPlayer()
	g.meteors = nil
	g.bullets = nil
}

func main() {
	g := &Game{
		player:           NewPlayer(),
		shootCooldown:    NewTimer(time.Second),
		meteorSpawnTimer: NewTimer(time.Second * 3),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
