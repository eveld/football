package main

import "github.com/veandco/go-sdl2/sdl"

// NewPlayer creates a new player entity.
func NewPlayer(renderer *sdl.Renderer, position Vector3) *Entity {
	player := &Entity{}

	transformComponent := NewTransformComponent(player, position, Vector3{X: 34, Y: 0, Z: 64}, Vector3{X: 17, Y: 17, Z: 64}, 0)
	player.AddComponent(transformComponent)

	shadowComponent := NewShadowComponent(player, renderer, "assets/shadow.png")
	player.AddComponent(shadowComponent)

	spriteComponent := NewSpriteComponent(player, renderer, "assets/player.png")
	player.AddComponent(spriteComponent)

	inputComponent := NewInputComponent(player)
	player.AddComponent(inputComponent)

	return player
}
