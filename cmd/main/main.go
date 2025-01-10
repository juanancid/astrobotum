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

	// Create an entity and add components
	player := world.AddEntity()
	world.AddComponent(player, &components.Position{X: 0, Y: 0})
	world.AddComponent(player, &components.Velocity{DX: 10, DY: 5})

	// Add the Movement System
	movementSystem := &systems.MovementSystem{}
	world.AddSystem(movementSystem)

	// Simple game loop
	lastTime := time.Now()
	for i := 0; i < 10; i++ { // Simulate 10 frames
		now := time.Now()
		dt := now.Sub(lastTime).Seconds()
		lastTime = now

		world.UpdateSystems(dt)

		// Display the player's position
		position := world.GetComponent(player, reflect.TypeOf(&components.Position{})).(*components.Position)
		fmt.Printf("Frame %d: Position X=%.2f, Y=%.2f\n", i, position.X, position.Y)

		time.Sleep(100 * time.Millisecond) // Simulate ~10 FPS
	}
}
