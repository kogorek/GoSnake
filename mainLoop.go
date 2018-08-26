package main

import "github.com/veandco/go-sdl2/sdl"
import "errors"
import "fmt"

type mainLoop struct{
	running bool
	window *sdl.Window
	renderer *sdl.Renderer
	player *snake
}

func (this *mainLoop) exit() {
	sdl.Quit();
	if this.window != nil{
		this.window.Destroy()
	}
	if this.renderer != nil {
		this.renderer.Destroy()
	}
}

func (this *mainLoop) run() {
	this.running = true
	for this.running {
		this.processEvents()
		this.update()
		this.render()
	}
}

func (this *mainLoop) update() {
}

func (this *mainLoop) processEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			this.running = false
			break
		case *sdl.KeyboardEvent:
			this.processKeyboardEvent(t)
		}
	}
}

func (this *mainLoop) processKeyboardEvent(event *sdl.KeyboardEvent) {
fmt.Println(event.Keysym.Scancode == sdl.SCANCODE_W)
	switch event.Keysym.Scancode{
		case sdl.SCANCODE_W:
			this.player.MoveUp()
			break

		case sdl.SCANCODE_S:
			this.player.MoveDown()
			break

		case sdl.SCANCODE_A:
			this.player.MoveLeft()
			break

		case sdl.SCANCODE_D:
			this.player.MoveRight()
			break
	}
}

func (this *mainLoop) render() {
	this.renderer.SetDrawColor(0, 0, 0, 255)
	this.renderer.Clear()

	this.renderer.SetDrawColor(255, 0, 0, 255)
	this.renderer.DrawLine(0,0,100,100)

	this.renderer.Present()
}


func NewMainLoop() (*mainLoop, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.New("Cannot init SDL2")
	}
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 window")
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 renderer")
	}

	player := NewSnake()

	mainLoop := mainLoop{true, window, renderer, player}

	return &mainLoop, nil
}