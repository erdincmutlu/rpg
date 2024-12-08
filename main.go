package main

import (
	"fmt"
	"image"
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img *ebiten.Image
	X   float64
	Y   float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Potion struct {
	*Sprite
	AmountHeal uint
}
type Game struct {
	player      *Player
	enemies     []*Enemy
	potions     []*Potion
	tilemapJSON *TilemapJSON
	tilemapImg  *ebiten.Image
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

	for _, enemy := range g.enemies {
		if enemy.FollowsPlayer {
			if enemy.X < g.player.X {
				enemy.X += 1
			}
			if enemy.X > g.player.X {
				enemy.X -= 1
			}
			if enemy.Y < g.player.Y {
				enemy.Y += 1
			}
			if enemy.Y > g.player.Y {
				enemy.Y -= 1
			}
		}
	}

	for _, potion := range g.potions {
		if g.player.X > potion.X {
			g.player.Health += potion.AmountHeal
			fmt.Printf("Picked up potion! Health: %d\n", g.player.Health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	for _, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {

			// get the tile position of the tile
			x := index % layer.Width
			y := index / layer.Width
			x *= 16
			y *= 16

			// Get the position on the image where the tile id is
			srcX := (id - 1) % 22
			srcY := (id - 1) / 22
			srcX *= 16
			srcY *= 16

			opts.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(g.tilemapImg.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image), &opts)

			opts.GeoM.Reset()
		}
	}

	// draw player
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, enemy := range g.enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)

		screen.DrawImage(
			enemy.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	for _, potion := range g.potions {
		opts.GeoM.Translate(potion.X, potion.Y)

		screen.DrawImage(
			potion.Img.SubImage(
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

	potionImg, _, err := ebitenutil.NewImageFromFile("assets/images/potion.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			Sprite: &Sprite{
				Img: playerImg,
				X:   50,
				Y:   50,
			},
			Health: 3,
		},
		enemies: []*Enemy{
			{
				&Sprite{Img: skeletonImg,
					X: 100,
					Y: 100,
				},
				true,
			},
			{
				&Sprite{Img: skeletonImg,
					X: 150,
					Y: 50,
				},
				false,
			},
		},
		potions: []*Potion{
			{
				&Sprite{
					Img: potionImg,
					X:   210,
					Y:   100,
				},
				1,
			},
		},
		tilemapJSON: tilemapJSON,
		tilemapImg:  tilemapImg,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
