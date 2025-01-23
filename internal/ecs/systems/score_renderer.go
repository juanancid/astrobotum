package systems

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// ScoreRenderer renders the player's score on the screen.
type ScoreRenderer struct {
	PlayerEntity ecs.Entity // The player entity
}

func (sr *ScoreRenderer) Render(w *ecs.World, screen *ebiten.Image) {
	// Get the player's score component
	score := w.GetComponent(sr.PlayerEntity, reflect.TypeOf(&components.Score{})).(*components.Score)

	// Display the score in the top-left corner
	scoreText := fmt.Sprintf("Score: %d", score.Points)
	ebitenutil.DebugPrintAt(screen, scoreText, 10, 10)
}
