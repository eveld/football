package main

import "github.com/veandco/go-sdl2/sdl"

// TransformComponent holds the transform details.
type TransformComponent struct {
	ID       ComponentID
	Parent   *Entity
	Position Vector3
	Size     Vector3
	Origin   Vector3
	Rotation float64
}

// NewTransformComponent creates a new transform component.
func NewTransformComponent(parent *Entity, position Vector3, size Vector3, origin Vector3, rotation float64) *TransformComponent {
	transformComponent := &TransformComponent{
		ID:       "transform",
		Parent:   parent,
		Position: position,
		Size:     size,
		Origin:   origin,
		Rotation: rotation,
	}
	return transformComponent
}

// GetID gets the ID of the component.
func (c *TransformComponent) GetID() ComponentID {
	return c.ID
}

// GetBounds returns the bounds of the entity.
func (c *TransformComponent) GetBounds() sdl.Rect {
	return sdl.Rect{
		X: c.Position.X,
		Y: c.Position.Y,
		W: c.Size.X,
		H: c.Size.Z,
	}
}

// Draw the component.
func (c *TransformComponent) Draw(renderer *sdl.Renderer) error {
	return nil
}

// Update the component.
func (c *TransformComponent) Update() error {
	return nil
}
