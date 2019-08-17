package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// SpriteComponent holds the details of a sprite.
type SpriteComponent struct {
	ID      ComponentID
	Parent  *Entity
	Texture *sdl.Texture
	Rect    sdl.Rect
}

// NewSpriteComponent creates a new sprite renderer.
func NewSpriteComponent(parent *Entity, renderer *sdl.Renderer, file string) *SpriteComponent {
	texture, err := TextureFromPNG(renderer, file)
	if err != nil {
		panic(err)
	}

	_, _, width, height, err := texture.Query()
	if err != nil {
		panic(err)
	}

	spriteComponent := &SpriteComponent{
		ID:      "sprite",
		Parent:  parent,
		Texture: texture,
		Rect:    sdl.Rect{X: 0, Y: 0, W: width, H: height},
	}

	return spriteComponent
}

// GetID gets the ID of the component.
func (c *SpriteComponent) GetID() ComponentID {
	return c.ID
}

// Draw the scene.
func (c *SpriteComponent) Draw(renderer *sdl.Renderer) error {
	transform := c.Parent.GetComponent(&TransformComponent{}).(*TransformComponent)
	x := transform.Position.X - transform.Origin.X
	y := transform.Position.Y - transform.Position.Z - transform.Origin.Z
	angle := transform.Rotation
	center := &sdl.Point{X: transform.Origin.X, Y: transform.Origin.Z}

	renderer.CopyEx(
		c.Texture,
		&c.Rect,
		&sdl.Rect{X: x, Y: y, W: c.Rect.W, H: c.Rect.H},
		angle,
		center,
		sdl.FLIP_NONE,
	)

	return nil
}

// Update the scene.
func (c *SpriteComponent) Update() error {
	return nil
}

// TextureFromPNG creates a texture from an image file.
func TextureFromPNG(renderer *sdl.Renderer, file string) (*sdl.Texture, error) {
	image, err := img.Load(file)
	if err != nil {
		return nil, err
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, err
	}

	return texture, nil
}
