package systems

import (
	"fmt"
	"reflect"

	"github.com/juanancid/astrobotum/internal/ecs"
	"github.com/juanancid/astrobotum/internal/ecs/components"
)

// ScoreSystem manages the player's score and updates it based on events.
type ScoreSystem struct {
	PlayerEntity ecs.Entity // The player entity
	TimeSurvived float64    // Tracks time for survival scoring
	TargetScore  int        // Score required to win the game
	Victory      bool       // Indicates if the player has won
}

func (ss *ScoreSystem) Update(w *ecs.World, dt float64) {
	// Increment score for time survived
	ss.TimeSurvived += dt
	if ss.TimeSurvived >= 1.0 { // Award points every second
		ss.TimeSurvived = 0
		score := w.GetComponent(ss.PlayerEntity, reflect.TypeOf(&components.Score{})).(*components.Score)
		score.Points += 10 // Award 10 points per second
		fmt.Printf("Score updated: %d points\n", score.Points)

		// Check for victory condition
		if score.Points >= ss.TargetScore {
			ss.Victory = true
			fmt.Println("Victory!")
		}
	}
}
