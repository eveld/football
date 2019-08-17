package main

import "github.com/veandco/go-sdl2/sdl"

// ComponentID holds the ID of the component
type ComponentID string

// Component defines a component for an entity.
type Component interface {
	GetID() ComponentID
	Draw(renderer *sdl.Renderer) error
	Update() error
}
