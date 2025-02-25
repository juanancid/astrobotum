package components

// Position defines the 2D coordinates of an entity.
type Position struct {
	X, Y float64
}

// Velocity defines the speed of an entity along X and Y axes.
type Velocity struct {
	DX, DY float64
}

// Size defines the width and height of an entity.
type Size struct {
	Width, Height float64
}

// PlayerControlled marks an entity as being controlled by the player.
type PlayerControlled struct{}

// Collectible marks an entity as an item the player can collect.
type Collectible struct {
	Value int // Points or value associated with the collectible
}

// StaticObstacle marks an entity as an immovable object.
type StaticObstacle struct{}

// DynamicObstacle marks an entity as a moving obstacle or hazard.
type DynamicObstacle struct {
	Damage int // Damage inflicted on the player upon collision
}

// OnGround indicates whether an entity is currently grounded.
type OnGround struct {
	IsGrounded bool
}

// Health tracks an entity's current and maximum health.
type Health struct {
	CurrentHealth int
	MaxHealth     int
}

// HealingCollectible marks an entity as a collectible that restores health.
type HealingCollectible struct {
	HealAmount int // Amount of health restored upon collection
}

// Score tracks the player's current score.
type Score struct {
	Points int // Total points accumulated
}
