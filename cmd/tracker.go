package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

type Break struct {
	startTime time.Time
	endTime   time.Time
}

type Tracker struct {
	stopChan    chan interface{}
	stopResult  chan time.Duration
	pauseChan   chan bool
	pauseResult chan []Break
}

func (t *Tracker) Start() {
	startTime := time.Now()

	var breaks []Break
	var currentBreakIndex int

	for {
		select {
		case <-t.stopChan:
			t.pauseResult <- breaks
			t.stopResult <- time.Since(startTime)
			return
		case isPause := <-t.pauseChan:
			if isPause {
				breaks = append(breaks, Break{startTime: time.Now()})
				currentBreakIndex = len(breaks) - 1
			} else {
				breaks[currentBreakIndex].endTime = time.Now()
			}
		}
	}
}

func formatDuration(duration time.Duration) string {
	return fmt.Sprintf("%.0f:%.0f:%.0f",
		math.Floor(duration.Hours()),
		math.Floor(duration.Minutes()),
		math.Floor(duration.Seconds()))
}

func RunCmd(workDir string, workHours int) {
	tracker := Tracker{
		stopChan:    make(chan interface{}),
		stopResult:  make(chan time.Duration),
		pauseChan:   make(chan bool),
		pauseResult: make(chan []Break),
	}

	reader := bufio.NewReader(os.Stdin)

	var isPause bool

	go tracker.Start()

	println("Press 's' to stop tracking")
	println("Press 'p' to pause/unpause tracking")

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
			tracker.stopChan <- struct{}{}

			breaks := <-tracker.pauseResult
			var totalBreaks time.Duration

			for _, b := range breaks {
				totalBreaks += b.endTime.Sub(b.startTime)
			}
			fmt.Printf("Your total break time: %s\n", formatDuration(totalBreaks))

			workedHours := <-tracker.stopResult
			workedHours -= totalBreaks

			fmt.Printf("Your total worked hours today: %s\n", formatDuration(workedHours))
			return
		case 'p':
			if isPause {
				println("Unpaused")
				isPause = false
			} else {
				println("Paused")
				isPause = true
			}
			tracker.pauseChan <- isPause
			continue
		default:
			println("Invalid command. Available commands: 's', 'p'")
		}
	}
}

func main() {
	var workDir string
	var workHours int

	flag.IntVar(&workHours, "hours", 8, "work hours")
	flag.StringVar(&workDir, "workDir", "./tracke_it_workdir", "work directory")
	flag.Parse()

	RunCmd(workDir, workHours)
}
