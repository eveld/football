package main

import "github.com/veandco/go-sdl2/sdl"

// NewPlayer creates a new player entity.
func NewPlayer(renderer *sdl.Renderer) *Entity {
	player := &Entity{}

	transformComponent := NewTransformComponent(player, Vector3{X: screenWidth / 2, Y: screenHeight / 2, Z: 50}, Vector3{X: 32, Y: 0, Z: 32}, 0)
	player.AddComponent(transformComponent)

	shadowComponent := NewShadowComponent(player, renderer, "assets/shadow.png", sdl.Point{X: 0, Y: 0})
	player.AddComponent(shadowComponent)

	spriteComponent := NewSpriteComponent(player, renderer, "assets/player.png", sdl.Point{X: 0, Y: 0})
	player.AddComponent(spriteComponent)

	return player
}
