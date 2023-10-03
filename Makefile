start-tracker:
	go run ./cmd/tracker.go -workDuration=8h

report:
	go run ./cmd/tracker.go -generateReport