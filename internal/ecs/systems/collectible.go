package systems

import (
	"astrobotum/internal/ecs"
	"astrobotum/internal/ecs/components"
	"fmt"
	"reflect"
)

// CollectibleSystem handles interactions between the player and collectibles.
type CollectibleSystem struct {
	PlayerEntity ecs.Entity // The entity controlled by the player
	Score        int        // Tracks the player's score
}

func (cs *CollectibleSystem) Update(w *ecs.World, dt float64) {
	playerPos := w.GetComponent(cs.PlayerEntity, reflect.TypeOf(&components.Position{})).(*components.Position)
	playerSize := w.GetComponent(cs.PlayerEntity, reflect.TypeOf(&components.Size{})).(*components.Size)

	collectibles := w.GetComponents(reflect.TypeOf(&components.Collectible{}))
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))

	for entity := range collectibles {
		collectiblePos := positions[entity].(*components.Position)
		collectibleSize := sizes[entity].(*components.Size)

		if isColliding(playerPos, playerSize, collectiblePos, collectibleSize) {
			// Increment the score
			collectible := collectibles[entity].(*components.Collectible)
			cs.Score += collectible.Value
			fmt.Printf("Collected item! New score: %d\n", cs.Score)

			// Remove the collectible entity
			w.RemoveEntity(entity)
		}
	}
}
