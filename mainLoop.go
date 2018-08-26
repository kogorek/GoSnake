package main

import "github.com/veandco/go-sdl2/sdl"
import "github.com/EngoEngine/glm"
import "errors"
// import "fmt"

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
	this.player.Update()
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

	this.renderer.SetDrawColor(60, 168, 0, 255)
	var bodyparts = this.player.bodyparts
	for i := bodyparts.Front(); i != nil; i = i.Next() {
		var part = i.Value.(glm.Vec2)
		var rect = sdl.Rect{int32(part.X())*10, int32(part.Y())*10, 10, 10}
		this.renderer.DrawRect(&rect)
	}
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
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()
	player.AddBodypart()


	mainLoop := mainLoop{true, window, renderer, player}

	return &mainLoop, nil
}