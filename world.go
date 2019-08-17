package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// NewWorld creates a new world entity.
func NewWorld(renderer *sdl.Renderer) *Entity {
	world := &Entity{}

	transformComponent := NewTransformComponent(world, sdl.Rect{X: 0, Y: 0, W: screenWidth, H: screenHeight}, 0)
	world.AddComponent(transformComponent)

	spriteComponent := NewSpriteComponent(world, renderer, "assets/world.png", sdl.Point{X: 0, Y: 0})
	world.AddComponent(spriteComponent)

	return world
}
