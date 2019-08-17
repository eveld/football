package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

// Vector3 holds 3 positions x, y, z.
type Vector3 struct {
	X, Y, Z int32
}

// Entity represents the most basic element.
type Entity struct {
	Active     bool
	Components []Component
}

// Draw draws all of the entities components.
func (e *Entity) Draw(renderer *sdl.Renderer) error {
	for _, component := range e.Components {
		err := component.Draw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

// Update updates all of the entities components.
func (e *Entity) Update() error {
	for _, component := range e.Components {
		err := component.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

// AddComponent adds a component to the entity.
func (e *Entity) AddComponent(new Component) {
	for _, existing := range e.Components {
		// if existing.GetID() == new.GetID() {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			fmt.Printf("Entity already has component of type %s", reflect.TypeOf(new))
			return
		}
	}

	e.Components = append(e.Components, new)
}

// GetComponent gets a component of an entity.
func (e *Entity) GetComponent(c Component) Component {
	componentType := reflect.TypeOf(c)
	for _, existing := range e.Components {
		if reflect.TypeOf(existing) == componentType {
			return existing
		}
	}

	return nil
	// panic(fmt.Errorf("Entity does not have a component of type %s", ID))
}
