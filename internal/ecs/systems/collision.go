package systems

import (
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// CollisionSystem checks for overlaps and handles collision responses.
type CollisionSystem struct {
	previousPositions map[ecs.Entity]*components.Position
}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{
		previousPositions: make(map[ecs.Entity]*components.Position),
	}
}

// SavePreviousPositions stores the positions of all entities before movement.
func (cs *CollisionSystem) SavePreviousPositions(w *ecs.World) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))

	for entity, pos := range positions {
		cs.previousPositions[entity] = &components.Position{
			X: pos.(*components.Position).X,
			Y: pos.(*components.Position).Y,
		}
	}
}

func (cs *CollisionSystem) Update(w *ecs.World, dt float64) {
	positions := w.GetComponents(reflect.TypeOf(&components.Position{}))
	sizes := w.GetComponents(reflect.TypeOf(&components.Size{}))
	grounds := w.GetComponents(reflect.TypeOf(&components.StaticObstacle{}))

	for entity := range positions {
		// Check if the entity has OnGround and Velocity components
		onGround := w.GetComponent(entity, reflect.TypeOf(&components.OnGround{}))
		if onGround == nil {
			continue
		}
		grounded := onGround.(*components.OnGround)

		velocity := w.GetComponent(entity, reflect.TypeOf(&components.Velocity{})).(*components.Velocity)
		entityPos := positions[entity].(*components.Position)
		entitySize := sizes[entity].(*components.Size)

		// Reset grounded state
		grounded.IsGrounded = false

		// Check for collisions with ground entities
		for groundEntity := range grounds {
			groundPos := positions[groundEntity].(*components.Position)
			groundSize := sizes[groundEntity].(*components.Size)

			if isColliding(entityPos, entitySize, groundPos, groundSize) {
				// If colliding and falling, stop downward velocity and mark as grounded
				if velocity.DY > 0 {
					velocity.DY = 0
					entityPos.Y = groundPos.Y - entitySize.Height // Align with the ground
					grounded.IsGrounded = true
					break
				}
			}
		}
	}
}

func isColliding(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	return pos1.X < pos2.X+size2.Width &&
		pos1.X+size1.Width > pos2.X &&
		pos1.Y < pos2.Y+size2.Height &&
		pos1.Y+size1.Height > pos2.Y
}

func isSweptColliding(
	prevPos1, currPos1 *components.Position, size1 *components.Size,
	pos2 *components.Position, size2 *components.Size,
) bool {
	return isColliding(prevPos1, size1, pos2, size2) || isColliding(currPos1, size1, pos2, size2)
}
