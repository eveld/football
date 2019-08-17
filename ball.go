package main

import "github.com/veandco/go-sdl2/sdl"

// NewBall creates a new ball entity.
func NewBall(renderer *sdl.Renderer) *Entity {
	ball := &Entity{}

	transformComponent := NewTransformComponent(ball, Vector3{X: screenWidth / 2, Y: screenHeight / 2, Z: 50}, Vector3{X: 32, Y: 0, Z: 32}, 0)
	ball.AddComponent(transformComponent)

	shadowComponent := NewShadowComponent(ball, renderer, "assets/shadow.png", sdl.Point{X: 0, Y: 0})
	ball.AddComponent(shadowComponent)

	spriteComponent := NewSpriteComponent(ball, renderer, "assets/ball.png", sdl.Point{X: 0, Y: 0})
	ball.AddComponent(spriteComponent)

	return ball
}
