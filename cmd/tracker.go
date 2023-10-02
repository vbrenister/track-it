package main

import (
	"bufio"
	"flag"
	"os"
	"time"

	"github.com/vbrenister/track-it/internal/prompt"
	"github.com/vbrenister/track-it/internal/reporter"
	"github.com/vbrenister/track-it/internal/tracker"
)

func runTracker(r *reporter.Reporter, workDurationLimit time.Duration) {
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
			r.WriteStatistics(t)
			return
		case 'p':
			if t.IsStopped {
				println("Tracker is stopped. You can't pause it")
				prompt.PrintStatistics(t)
				r.WriteStatistics(t)
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
				r.WriteStatistics(t)
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
	var isMonthlyReport bool

	flag.DurationVar(&workDurationLimit, "workDuration", time.Hour*8, "The max duration of work")
	flag.StringVar(&workDir, "reportDir", "./reports", "Report directory")
	flag.BoolVar(&isMonthlyReport, "generateReport", false, "Monthly report")
	flag.Parse()

	r := reporter.NewReporter(workDir)

	if isMonthlyReport {
		r.MonthlyReport()
		return
	}

	runTracker(r, workDurationLimit)
}
