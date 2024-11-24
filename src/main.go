package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"hywelbane/physics"
	"log"
)

const (
	fps               = 60
	millisecsPerFrame = 1000 / fps
	pixelsPerMeter    = 50

	window_height = 1024
	window_width  = 768
)

var (
	running           bool
	window            *sdl.Window
	renderer          *sdl.Renderer
	gravity           = physics.NewVec2(0.0, 9.8*pixelsPerMeter)
	smallParticle     = physics.NewParticle(100.0, 100.0, 1)
	previousFrameTime uint64
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

func Setup() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %s", err)
	}

	var err error
	window, err = sdl.CreateWindow("Hywelbane", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, window_height, window_width, sdl.WINDOW_SHOWN)
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
	currentFrameTime := sdl.GetTicks64()
	if previousFrameTime == 0 {
		previousFrameTime = currentFrameTime
	}

	waitTimeToReachTargetFrameTime := millisecsPerFrame - uint32(currentFrameTime-previousFrameTime)
	if waitTimeToReachTargetFrameTime > 0 && waitTimeToReachTargetFrameTime <= millisecsPerFrame {
		sdl.Delay(waitTimeToReachTargetFrameTime)
	}
	deltaInSecs := float64(currentFrameTime-previousFrameTime) / 1000.0
	previousFrameTime = currentFrameTime

	wind := physics.NewVec2(2.0*pixelsPerMeter, 0.0)
	weight := physics.NewVec2(0.0, gravity.Multiply(smallParticle.Mass).Y)
	smallParticle.ApplyForce(wind)
	smallParticle.ApplyForce(gravity)
	smallParticle.ApplyForce(weight)

	smallParticle.Integrate(deltaInSecs)

	// temporary hack to keep smallParticle inside window - horizontal boudaries
	if smallParticle.Position.X <= 0 {
		smallParticle.Position.X = 0.0
		smallParticle.Velocity.X *= -1.0
	} else if smallParticle.Position.X >= window_width {
		smallParticle.Position.X = window_width
		smallParticle.Velocity.X *= -1.0
	}
	// vertical boundaries
	if smallParticle.Position.Y <= 0 {
		smallParticle.Position.Y = 0.0
		smallParticle.Velocity.Y *= -1.0
	} else if smallParticle.Position.Y >= window_height {
		smallParticle.Position.Y = window_height
		smallParticle.Velocity.Y *= -1.0
	}
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
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	DrawCircle(renderer, int(smallParticle.Position.X), int(smallParticle.Position.Y), 10)

	// Arrow to see the resulting force being applied to the particle
	//	renderer.SetDrawColor(0, 0, 255, 255)
	//	startX := int32(smallParticle.Position.X)
	//	startY := int32(smallParticle.Position.Y)
	//	endX := int32(smallParticle.Position.X + smallParticle.Velocity.X)
	//	endY := int32(smallParticle.Position.Y + smallParticle.Velocity.Y)
	//	renderer.DrawLine(startX, startY, endX, endY)

	renderer.Present()
}

func DrawCircle(renderer *sdl.Renderer, x, y, radius int) {
	renderer.SetDrawColor(100, 0, 0, 100)
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
