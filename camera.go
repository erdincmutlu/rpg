package main

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

func (c *Camera) FollowCamera(targetX float64, targetY float64, screenWidth float64,
	screenHeight float64) {

	c.X = -targetX + screenWidth/2.0
	c.Y = -targetY + screenHeight/2.0
}
