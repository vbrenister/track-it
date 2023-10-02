# track-it
Simple command line application which tracks your work time with capabilities to pause and unpause work and generate weekly and monthly CSV reports 


## How to use

### Run locally
```
go run ./cmd/tracker.go -help
```

Supported flags:
- workDuration - work duration limit (which can be parsed by `time.ParseDuration`. Ex: 1h30m15s, 25m, 10s). Default: 8h
- reportDir - directory where reports will be stored. Default: `./reports`
- generateReport - flag which indicates if report should be generated. Default: false, if is true, then tracker won't start and will generate report from last month

Example:
```
go run ./cmd/tracker.go -workDuration 8h -reportDir ./reports -generateReport
```

You can download binary from [releases](www.google.com) based on your OS and architecture.