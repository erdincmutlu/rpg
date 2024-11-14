package main

import (
	"image"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img           *ebiten.Image
	X             float64
	Y             float64
	FollowsPlayer bool
}

type Game struct {
	player  *Sprite
	sprites []*Sprite
}

func (g *Game) Update() error {
	// react to key presses
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.Y += 2
	}

	for _, sprite := range g.sprites {
		if sprite.FollowsPlayer {
			if sprite.X < g.player.X {
				sprite.X += 1
			}
			if sprite.X > g.player.X {
				sprite.X -= 1
			}
			if sprite.Y < g.player.Y {
				sprite.Y += 1
			}
			if sprite.Y > g.player.Y {
				sprite.Y -= 1
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	// draw player
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, sprite := range g.sprites {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Sprite{
			Img: playerImg,
			X:   50,
			Y:   50,
		},
		sprites: []*Sprite{
			{
				Img:           skeletonImg,
				X:             100,
				Y:             100,
				FollowsPlayer: true,
			},
			{
				Img:           skeletonImg,
				X:             150,
				Y:             150,
				FollowsPlayer: true,
			},
			{
				Img:           skeletonImg,
				X:             75,
				Y:             75,
				FollowsPlayer: false,
			},
		},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
