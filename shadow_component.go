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
}

// NewShadowComponent creates a new shadow.
func NewShadowComponent(parent *Entity, renderer *sdl.Renderer, file string) *ShadowComponent {
	texture, err := TextureFromPNG(renderer, file)
	if err != nil {
		panic(err)
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	texture.SetAlphaMod(190)

	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(err)
	}

	shadowComponent := &ShadowComponent{
		ID:      "shadow",
		Parent:  parent,
		Texture: texture,
		Rect:    sdl.Rect{X: 0, Y: 0, W: width, H: height},
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

	// If the entity is on the floor, don't show a shadow.
	if transform.Position.Z == 0 {
		return nil
	}

	x := transform.Position.X - transform.Origin.X + c.Rect.W/2
	y := transform.Position.Y + (transform.Size.Z - c.Rect.H) - transform.Origin.Z
	angle := transform.Rotation

	renderer.CopyEx(
		c.Texture,
		&c.Rect,
		&sdl.Rect{X: x, Y: y, W: c.Rect.W, H: c.Rect.H},
		angle,
		&sdl.Point{X: x, Y: y},
		sdl.FLIP_NONE,
	)

	return nil
}

// Update the scene.
func (c *ShadowComponent) Update() error {
	return nil
}
