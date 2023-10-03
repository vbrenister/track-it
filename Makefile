start-standard-tracker:
	go run ./cmd/tracker.go

report:
	go run ./cmd/tracker.go -generateReport

.PHONY: start-standard-tracker report