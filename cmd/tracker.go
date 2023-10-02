package main

import (
	"bufio"
	"flag"
	"os"
	"time"

	"github.com/vbrenister/track-it/internal/prompt"
	"github.com/vbrenister/track-it/internal/tracker"
)

func runTracker(workDir string, workDurationLimit time.Duration) {
	t := tracker.NewTracker(workDurationLimit)

	reader := bufio.NewReader(os.Stdin)

	t.Start()

	prompt.PrintCommands()

	for {
		command, _, err := reader.ReadRune()

		if err != nil {
			panic("Error reading command")
		}

		if command == '\n' {
			continue
		}

		switch command {
		case 's':
			t.Stop()
			prompt.PrintStatistics(t)
			return
		case 'p':
			if t.IsStopped {
				println("Tracker is stopped. You can't pause it")
				return
			}
			if t.IsPaused {
				println("Unpaused")
				t.Unpause()
			} else {
				println("Paused")
				t.Pause()
			}
			continue
		case 'q':
			if t.IsStopped {
				prompt.PrintStatistics(t)
			}
			return
		default:
			println("Invalid command. Available commands: ['s', 'p', 'q']")
		}
	}
}

func main() {
	var workDir string
	var workDurationLimit time.Duration

	flag.DurationVar(&workDurationLimit, "workDuration", time.Hour*8, "The max duration of work")
	flag.StringVar(&workDir, "reportDir", "./tracke_it_workdir", "Report directory")
	flag.Parse()

	runTracker(workDir, workDurationLimit)
}
