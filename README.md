# track-it
Simple command line application which tracks your work time with capabilities to pause and unpause work and generate weekly and monthly CSV reports 


## How to use

### Run locally
Make sure you have go installed. Follow the instruction from [here](https://go.dev/doc/install)

```
go run ./cmd/tracker.go -help
```

Supported flags:
- `help` - prints help
- `workDuration` - work duration limit which can be parsed by [time.ParseDuration](https://pkg.go.dev/time#ParseDuration). Default: `8h`
- `reportDir` - directory where reports will be stored. Default: `./reports`
- `generateReport` - flag which indicates if report should be generated. If it's passed, then tracker won't start and will generate report from last month

Example:
```
go run ./cmd/tracker.go -workDuration 8h -reportDir ./reports -generateReport
```

#### Using Makefile

Besides runnging plain `go run` command you can use Makefile to run tracker.

Start tracker with default flags
```
make start-tracker
```

Generate monthly report
```
make report
```

### Using binary
You can download binary from [releases](www.google.com) based on your OS and architecture.

> Note: For Linux and Mac users make sure you granted execution permission to binary
```
chmod +x tracker
```

Run binary with flags
```
./tracker -workDuration 8h -reportDir ./reports -generateReport
```

## Reporting

When tracker stops tracking, it will create a day entry in `{reportDir}/{CURRENT_YEAR}/{CURRENT_MONTH}/{CURRENT_DAY}.csv` file.

When you run tracker with `generateReport` flag, it will collect all daily entries from current month and generate a report in `{reportDir}/{CURRENT_YEAR}/{CURRENT_MONTH}/report.csv` file.
