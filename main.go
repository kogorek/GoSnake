package main

import "github.com/veandco/go-sdl2/sdl"
// import "github.com/EngoEngine/glm"

import "errors"
import "os"
import "fmt"

type MainLoop struct{
	running bool;
	window *sdl.Window;
	renderer *sdl.Renderer;
}

func (this *MainLoop) exit() {
	sdl.Quit();
	if this.window != nil{
		this.window.Destroy();
	}
	if this.renderer != nil {
		this.renderer.Destroy();
	}
}

func (this *MainLoop) run() {
	this.running = true;
	for this.running {
		this.update();
		this.draw();
	}
}

func (this *MainLoop) update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			this.running = false
			break
		}
	}
}

func (this *MainLoop) render() {
	this.renderer.SetDrawColor(0, 0, 0, 255);
	this.renderer.Clear();
	this.renderer.SetDrawColor(255, 0, 0, 255);
	this.renderer.DrawLine(0,0,100,100);
	this.renderer.Present();
}


func NewMainLoop() (*MainLoop, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, errors.New("Cannot init SDL2");
	}
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 window");
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, errors.New("Cannot create sdl2 renderer");
	}

	mainLoop := MainLoop{true, window, renderer};

	return &mainLoop, nil;
}

func main() {
	app, err := NewMainLoop();
	if err != nil {
		fmt.Println(err);
		os.Exit(-1);
	}

	app.run();
	app.exit();
}