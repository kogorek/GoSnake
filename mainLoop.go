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
	food []glm.Vec2
	foodCount int
	fieldSize glm.Vec2
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
	this.spawnFood()
	this.foodColision()
	this.player.Update()
	if this.player.checkSelfCollision(){
		this.running = false
	}
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

		case sdl.SCANCODE_Q:
			this.player.AddBodypart()
			break
	}
}

func (this *mainLoop) render() {
	this.renderer.SetDrawColor(0, 0, 0, 255)
	this.renderer.Clear()

	this.renderer.SetDrawColor(60, 168, 0, 255)
	for _, part := range this.player.bodyparts {
		var rect = sdl.Rect{int32(part.X())*30, int32(part.Y())*30, 30, 30}
		this.renderer.DrawRect(&rect)
	}

	this.renderer.SetDrawColor(168, 60, 0, 255)
	for _, part := range this.food {
		var rect = sdl.Rect{int32(part.X())*30, int32(part.Y())*30, 30, 30}
		this.renderer.DrawRect(&rect)
	}
	this.renderer.Present()
}

func (this *mainLoop) spawnFood() {
	if this.foodCount > len(this.food){
		this.food = append(this.food, glm.Vec2{2.0, 2.0})
	}
}

func (this *mainLoop) foodColision() {
	for i := 0; i < len(this.food); i++ {
		println(i)
		if this.player.checkPointCollision(this.food[i]) {
			this.player.AddBodypart()
		}
	}
}

func NewMainLoop() (*mainLoop, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.New("Cannot init SDL2")
	}
	window, err := sdl.CreateWindow("Frekin' snake!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 window")
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 renderer")
	}

	player := NewSnake(glm.Vec2{2, 2})
	var food []glm.Vec2 = make([]glm.Vec2, 0)
	var fieldSize glm.Vec2 = glm.Vec2{10.0, 10.0}
	var foodCount int = 1
	mainLoop := mainLoop{true, window, renderer, player, food, foodCount, fieldSize}

	return &mainLoop, nil
}