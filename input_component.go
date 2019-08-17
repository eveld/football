package main

import "github.com/veandco/go-sdl2/sdl"

// InputComponent allows the user to control an entity.
type InputComponent struct {
	ID       ComponentID
	Parent   *Entity
	Position Vector3
	Size     Vector3
	Origin   Vector3
	Rotation float64
}

// NewInputComponent creates a new input component.
func NewInputComponent(parent *Entity) *InputComponent {
	inputComponent := &InputComponent{
		ID:     "input",
		Parent: parent,
	}
	return inputComponent
}

// GetID gets the ID of the component.
func (c *InputComponent) GetID() ComponentID {
	return c.ID
}

// Draw the component.
func (c *InputComponent) Draw(renderer *sdl.Renderer) error {
	return nil
}

// Update the component.
func (c *InputComponent) Update() error {
	keys := sdl.GetKeyboardState()

	transform := c.Parent.GetComponent(&TransformComponent{}).(*TransformComponent)

	if keys[sdl.SCANCODE_LEFT] == 1 {
		transform.Position.X -= 1
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		transform.Position.X += 1
	} else if keys[sdl.SCANCODE_UP] == 1 {
		transform.Position.Y -= 1
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		transform.Position.Y += 1
	}

	return nil
}
