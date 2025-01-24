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

		for groundEntity := range grounds {
			groundPos := positions[groundEntity].(*components.Position)
			groundSize := sizes[groundEntity].(*components.Size)

			if isColliding(entityPos, entitySize, groundPos, groundSize) {
				// Calculate overlap amounts
				overlapX := calculateOverlap(entityPos.X, entitySize.Width, groundPos.X, groundSize.Width)
				overlapY := calculateOverlap(entityPos.Y, entitySize.Height, groundPos.Y, groundSize.Height)

				// Resolve the smaller overlap first (prioritize vertical collisions)
				if abs(overlapY) < abs(overlapX) {
					if velocity.DY > 0 { // Falling
						entityPos.Y = groundPos.Y - entitySize.Height // Align with ground
						velocity.DY = 0
						grounded.IsGrounded = true
					} else if velocity.DY < 0 { // Jumping and hitting ceiling
						entityPos.Y = groundPos.Y + groundSize.Height
						velocity.DY = 0
					}
				} else { // Horizontal collision
					if velocity.DX > 0 { // Moving right
						entityPos.X = groundPos.X - entitySize.Width
					} else if velocity.DX < 0 { // Moving left
						entityPos.X = groundPos.X + groundSize.Width
					}
					velocity.DX = 0 // Stop horizontal motion
				}

				break // Exit loop after handling one collision
			}
		}
	}
}

// calculateOverlap calculates the overlap distance between two intervals.
func calculateOverlap(pos1, size1, pos2, size2 float64) float64 {
	if pos1 < pos2 {
		return pos1 + size1 - pos2 // Overlap on the right
	}
	return pos1 - (pos2 + size2) // Overlap on the left
}

// abs returns the absolute value of a float64.
func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

func isColliding(pos1 *components.Position, size1 *components.Size, pos2 *components.Position, size2 *components.Size) bool {
	return pos1.X < pos2.X+size2.Width &&
		pos1.X+size1.Width > pos2.X &&
		pos1.Y < pos2.Y+size2.Height &&
		pos1.Y+size1.Height > pos2.Y
}
