package systems

import (
	"fmt"
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// DynamicObstacleCollisionSystem handles collisions between the player and dynamic obstacles.
type DynamicObstacleCollisionSystem struct {
	PlayerEntity ecs.Entity // The player entity
}

func (docs *DynamicObstacleCollisionSystem) Update(w *ecs.World, dt float64) {
	playerPos := w.GetComponent(docs.PlayerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	playerSize := w.GetComponent(docs.PlayerEntity, reflect.TypeOf(&components.Size{})).(*components.Size)

	dynamicObstacles := w.GetComponents(reflect.TypeOf(&components.DynamicObstacle{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

	for entity := range dynamicObstacles {
		obstaclePos := positions[entity].(*components.Position)
		obstacleSize := sizes[entity].(*components.Size)

		if isColliding(playerPos, playerSize, obstaclePos, obstacleSize) {
			// Handle collision response
			dynamic := dynamicObstacles[entity].(*components.DynamicObstacle)
			fmt.Printf("Player collided with obstacle! Damage: %d\n", dynamic.Damage)

			// Reduce player health
			health := w.GetComponent(docs.PlayerEntity, reflect.TypeOf(&components.Health{})).(*components.Health)
			health.CurrentHealth -= dynamic.Damage
			fmt.Printf("Player health: %d/%d\n", health.CurrentHealth, health.MaxHealth)
		}
	}
}
