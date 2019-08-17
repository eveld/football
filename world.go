package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// NewWorld creates a new world entity.
func NewWorld(renderer *sdl.Renderer) *Entity {
	world := &Entity{}

	transformComponent := NewTransformComponent(world, Vector3{X: 0, Y: 0, Z: 0}, Vector3{X: screenWidth, Y: 0, Z: screenHeight}, 0)
	world.AddComponent(transformComponent)

	spriteComponent := NewSpriteComponent(world, renderer, "assets/world.png", sdl.Point{X: 0, Y: 0})
	world.AddComponent(spriteComponent)

	return world
}
