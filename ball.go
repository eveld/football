package main

import "github.com/veandco/go-sdl2/sdl"

// NewBall creates a new ball entity.
func NewBall(renderer *sdl.Renderer) *Entity {
	ball := &Entity{}

	transformComponent := NewTransformComponent(ball, Vector3{X: 400, Y: screenHeight / 2, Z: 10}, Vector3{X: 28, Y: 0, Z: 28}, Vector3{X: 14, Y: 14, Z: 28}, 0)
	ball.AddComponent(transformComponent)

	shadowComponent := NewShadowComponent(ball, renderer, "assets/shadow.png")
	ball.AddComponent(shadowComponent)

	spriteComponent := NewSpriteComponent(ball, renderer, "assets/ball.png")
	ball.AddComponent(spriteComponent)

	return ball
}
