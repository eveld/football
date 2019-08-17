package main

import "github.com/veandco/go-sdl2/sdl"

// TransformComponent holds the transform details.
type TransformComponent struct {
	ID       ComponentID
	Parent   *Entity
	Bounds   sdl.Rect
	Rotation float64
}

// NewTransformComponent creates a new transform component.
func NewTransformComponent(parent *Entity, bounds sdl.Rect, rotation float64) *TransformComponent {
	transformComponent := &TransformComponent{
		ID:       "transform",
		Parent:   parent,
		Bounds:   bounds,
		Rotation: rotation,
	}
	return transformComponent
}

// GetID gets the ID of the component.
func (c *TransformComponent) GetID() ComponentID {
	return c.ID
}

// Draw the component.
func (c *TransformComponent) Draw(renderer *sdl.Renderer) error {
	return nil
}

// Update the component.
func (c *TransformComponent) Update() error {
	return nil
}
