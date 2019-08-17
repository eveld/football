package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// ShadowComponent holds the details of a Shadow.
type ShadowComponent struct {
	ID      ComponentID
	Parent  *Entity
	Texture *sdl.Texture
	Rect    sdl.Rect
	Origin  sdl.Point
}

// NewShadowComponent creates a new shadow.
func NewShadowComponent(parent *Entity, renderer *sdl.Renderer, file string, origin sdl.Point) *ShadowComponent {
	texture, err := TextureFromPNG(renderer, file)
	if err != nil {
		panic(err)
	}

	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(err)
	}

	shadowComponent := &ShadowComponent{
		ID:      "shadow",
		Parent:  parent,
		Texture: texture,
		Rect:    sdl.Rect{X: 0, Y: 0, W: width, H: height},
		Origin:  origin,
	}

	return shadowComponent
}

// GetID gets the ID of the component.
func (c *ShadowComponent) GetID() ComponentID {
	return c.ID
}

// Draw the scene.
func (c *ShadowComponent) Draw(renderer *sdl.Renderer) error {
	transform := c.Parent.GetComponent(&TransformComponent{}).(*TransformComponent)
	x := transform.Position.X + (transform.Size.X-c.Rect.W)/2
	y := transform.Position.Y + (transform.Size.Z - c.Rect.H)
	angle := transform.Rotation

	renderer.CopyEx(
		c.Texture,
		&c.Rect,
		&sdl.Rect{X: x, Y: y, W: c.Rect.W, H: c.Rect.H},
		angle,
		&sdl.Point{X: c.Origin.X, Y: c.Origin.Y},
		sdl.FLIP_NONE,
	)

	return nil
}

// Update the scene.
func (c *ShadowComponent) Update() error {
	return nil
}
