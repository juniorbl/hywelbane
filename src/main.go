package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var (
	running  bool
	window   *sdl.Window
	renderer *sdl.Renderer
)

func main() {
	running = true
	Setup()
	defer Cleanup()
	for running {
		Input()
		Update()
		Render()
	}
}

var s = 500

func Setup() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %s", err)
	}

	var err error
	window, err = sdl.CreateWindow("Hywelbane", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1024, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Failed to create window: %s", err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s", err)
	}
}

func Cleanup() {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

func Update() {

}

func Input() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			running = false
		}
	}
}

func Render() {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	s = s + 1
	DrawCircle(renderer, s, 384, 10)
	renderer.DrawLine(100, 100, 40, 90)
	renderer.Present()
}

func DrawCircle(renderer *sdl.Renderer, x, y, radius int) {
	renderer.SetDrawColor(255, 0, 0, 255)
	for w := 0; w < radius*2; w++ {
		for h := 0; h < radius*2; h++ {
			dx := radius - w
			dy := radius - h
			if (dx*dx + dy*dy) <= (radius * radius) {
				renderer.DrawPoint(int32(x+dx), int32(y+dy))
			}
		}
	}
}
