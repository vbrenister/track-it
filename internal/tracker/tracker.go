package tracker

import "time"

type Break struct {
	startTime time.Time
	endTime   time.Time
}

type Tracker struct {
	workDurationLimit time.Duration
	startTime         time.Time
	stopTime          time.Time
	IsStopped         bool
	IsPaused          bool
	breaks            []Break

	stopSignal  chan interface{}
	pauseSignal chan interface{}
}

func NewTracker(workDurationLimit time.Duration) *Tracker {
	return &Tracker{
		workDurationLimit: workDurationLimit,
		stopSignal:        make(chan interface{}),
		pauseSignal:       make(chan interface{}),
	}
}

func (t *Tracker) WorkedDuration() time.Duration {
	return t.TotalDuration() - t.BreaksDuration()
}

func (t *Tracker) TotalDuration() time.Duration {
	if t.IsStopped {
		return t.stopTime.Sub(t.startTime)
	}
	return time.Since(t.startTime)
}

func (t *Tracker) Start() {
	t.startTime = time.Now()

	go func() {
		for {
			select {
			case <-t.stopSignal:
				t.IsStopped = true
				t.stopTime = time.Now()

				return
			case <-t.pauseSignal:
				if t.IsPaused {
					t.breaks = append(t.breaks, Break{startTime: time.Now()})
				} else {
					t.breaks[len(t.breaks)-1].endTime = time.Now()
				}
			default:
				if t.WorkedDuration() >= t.workDurationLimit {
					t.IsStopped = true
					t.stopTime = time.Now()
					println("Work time exceeded. Stopping time tracking")
					return
				}
			}
		}
	}()
}

func (t *Tracker) Stop() {
	if t.IsStopped {
		return
	}
	t.stopSignal <- struct{}{}
}

func (t *Tracker) Pause() {
	if t.IsPaused {
		return
	}

	t.IsPaused = true
	t.pauseSignal <- struct{}{}
}

func (t *Tracker) Unpause() {
	if !t.IsPaused {
		return
	}
	t.IsPaused = false
	t.pauseSignal <- struct{}{}
}

func (t *Tracker) BreaksDuration() time.Duration {
	var totalBreaks time.Duration

	for _, b := range t.breaks {
		totalBreaks += b.endTime.Sub(b.startTime)
	}

	return totalBreaks
}
