package main

import "github.com/veandco/go-sdl2/sdl"
import "errors"

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
		this.processEvents();
		this.update();
		this.render();
	}
}

func (this *MainLoop) update() {
}

func (this *MainLoop) processEvents() {
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