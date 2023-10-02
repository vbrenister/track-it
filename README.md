# track-it
Simple command line application which tracks your work time with capabilities to pause and unpause work and generate weekly and monthly CSV reports 


## How to use

### Run locally
Make sure you have go installed. Follow the instruction from [here](https://go.dev/doc/install)

```
go run ./cmd/tracker.go -help
```

Supported flags:
- help - prints help
- workDuration - work duration limit (which can be parsed by `time.ParseDuration`. Ex: 1h30m15s, 25m, 10s). Default: 8h
- reportDir - directory where reports will be stored. Default: `./reports`
- generateReport - flag which indicates if report should be generated. Default: false, if is true, then tracker won't start and will generate report from last month

Example:
```
go run ./cmd/tracker.go -workDuration 8h -reportDir ./reports -generateReport
```

### Using binary
You can download binary from [releases](www.google.com) based on your OS and architecture.

For Linux and Mac users make sure you granted execution permission to binary
```
chmod +x tracker
```

Run binary with flags
```
./tracker -workDuration 8h -reportDir ./reports -generateReport
```