package systems

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/juanancid/astrobotum/internal/ecs"
)

type VictoryScreen struct {
	PlayerEntity ecs.Entity
	Active       bool // Indicates if the victory screen is active
	Score        int  // Final score to display
}

func (vs *VictoryScreen) Render(w *ecs.World, screen *ebiten.Image) {
	if !vs.Active {
		return
	}

	// Display the victory message
	text := fmt.Sprintf("Victory!\nFinal Score: %d\nPress R to Restart", vs.Score)
	ebitenutil.DebugPrintAt(screen, text, 50, 100)
}
