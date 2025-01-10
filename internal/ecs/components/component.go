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
