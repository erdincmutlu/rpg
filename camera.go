package main

import "math"

type Camera struct {
	X float64
	Y float64
}

func NewCamera(x float64, y float64) *Camera {
	return &Camera{
		X: x,
		Y: y,
	}
}

func (c *Camera) FollowTarget(targetX float64, targetY float64, screenWidth float64,
	screenHeight float64) {

	c.X = -targetX + screenWidth/2.0
	c.Y = -targetY + screenHeight/2.0
}

func (c *Camera) Constrain(tilemapWidthPixels, tilemapHeightPixels, screenWidth, screenHeight float64) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)
}
