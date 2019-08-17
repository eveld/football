package main

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	screenWidth  = int32(1280)
	screenHeight = int32(720)
	entities     = []*Entity{}
	running      = true
)

func setup() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}

	// desktop, err := sdl.GetDesktopDisplayMode(0)
	// if err != nil {
	// 	return nil, nil, err
	// }

	window, err := sdl.CreateWindow(
		"Football",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(screenWidth), int32(screenHeight),
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		return nil, nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, nil, err
	}

	return window, renderer, nil
}

func handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
			return
		case *sdl.WindowEvent:
			we := event.(*sdl.WindowEvent)
			if we.Event == sdl.WINDOWEVENT_RESIZED {
				screenWidth = we.Data1
				screenHeight = we.Data2
			}
			return
		}
	}
}

func update() error {
	for _, entity := range entities {
		err := entity.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func draw(renderer *sdl.Renderer) error {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()

	// TODO: add z-index ordering of entities and draw then back to front.

	for _, entity := range entities {
		err := entity.Draw(renderer)
		if err != nil {
			return err
		}
	}

	renderer.Present()

	return nil
}

func main() {
	window, renderer, err := setup()
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	world := NewWorld(renderer)
	ball := NewBall(renderer)
	entities = append(entities, world, ball)

	for i := 0; i < 1; i++ {
		rand.Seed(int64(i))
		x := int32(rand.Intn(int(screenWidth)))
		y := int32(rand.Intn(int(screenHeight)))
		fmt.Printf("x: %d / y: %d", x, y)

		player := NewPlayer(renderer, Vector3{X: x, Y: y, Z: 0})
		entities = append(entities, player)
	}

	for running == true {
		handleEvents()
		err = update()
		if err != nil {
			fmt.Printf("Could not update: %v", err)
		}
		err = draw(renderer)
		if err != nil {
			fmt.Printf("Could not draw: %v", err)
		}
	}
}
