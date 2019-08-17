package main

import "github.com/veandco/go-sdl2/sdl"

// NewBall creates a new ball entity.
func NewBall(renderer *sdl.Renderer) *Entity {
	ball := &Entity{}

	transformComponent := NewTransformComponent(ball, sdl.Rect{X: screenWidth / 2, Y: screenHeight / 2, W: 32, H: 32}, 0)
	ball.AddComponent(transformComponent)

	spriteComponent := NewSpriteComponent(ball, renderer, "assets/ball.png", sdl.Point{X: 0, Y: 0})
	ball.AddComponent(spriteComponent)

	return ball
}
