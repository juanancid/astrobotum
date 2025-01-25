package levels

import "github.com/juanancid/astrobotum/internal/ecs/components"

type Level struct {
	ID          int                 // Level identifier
	PlayerStart components.Position // Starting position of the player
}
