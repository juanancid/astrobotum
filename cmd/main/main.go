package main

import (
	"fmt"
	"reflect"
	"time"

	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
	"astrobotum/internal/ecs/systems"
)

func main() {
	world := ecs.NewWorld()

	// Create a player entity and add components
	player := world.AddEntity()
	world.AddComponent(player, &components.Position{X: 160, Y: 120}) // Centered position
	world.AddComponent(player, &components.Velocity{DX: 0, DY: 0})   // Initial velocity

	// Add systems
	inputSystem := &systems.InputSystem{}
	movementSystem := &systems.MovementSystem{}
	world.AddSystem(inputSystem)
	world.AddSystem(movementSystem)

	// Simple game loop
	lastTime := time.Now()
	for i := 0; i < 100; i++ { // Simulate 100 frames
		now := time.Now()
		dt := now.Sub(lastTime).Seconds()
		lastTime = now

		world.UpdateSystems(dt)

		// Display the player's position
		position := world.GetComponent(player, reflect.TypeOf(&components.Position{})).(*components.Position)
		fmt.Printf("Frame %d: Position X=%.2f, Y=%.2f\n", i, position.X, position.Y)

		time.Sleep(16 * time.Millisecond) // Simulate ~60 FPS
	}
}
