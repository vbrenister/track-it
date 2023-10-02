package prompt

import (
	"fmt"
	"math"
	"time"

	"github.com/vbrenister/track-it/internal/tracker"
)

func formatDuration(duration time.Duration) string {
	return fmt.Sprintf("%.0f:%.0f:%.0f",
		math.Floor(duration.Hours()),
		math.Floor(duration.Minutes()),
		math.Floor(duration.Seconds()))
}

func PrintCommands() {
	println("Press 's' to stop tracking")
	println("Press 'p' to pause/unpause tracking")
	println("Press 'q' to quit")
}

func PrintStatistics(t *tracker.Tracker) {
	fmt.Println("Your time statistics:")
	fmt.Printf("Break duration: %s\n", formatDuration(t.BreaksDuration()))
	fmt.Printf("Worked duration: %s\n", formatDuration(t.WorkedDuration()))
	fmt.Printf("Total duration: %s\n", formatDuration(t.TotalDuration()))
}
