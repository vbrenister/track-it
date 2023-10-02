package reporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/vbrenister/track-it/internal/fileutils"
	"github.com/vbrenister/track-it/internal/tracker"
)

type Reporter struct {
	workDir string
}

func NewReporter(workDir string) *Reporter {
	return &Reporter{workDir: workDir}
}

func (r *Reporter) MonthlyReport() {
	currentDate := time.Now()

	monthFolder := fmt.Sprintf("%s/%s/%s", r.workDir, fmt.Sprint(currentDate.Year()), currentDate.Month().String())
	if !fileutils.IsExist(monthFolder) {
		println("No data for this month")
		return
	}

	dailyFiles, err := os.ReadDir(monthFolder)
	if err != nil {
		panic(err)
	}

	var reportContent [][]string

	for _, dailyFile := range dailyFiles {
		f := fileutils.MustOpen(fmt.Sprintf("%s/%s", monthFolder, dailyFile.Name()))

		csvReader := csv.NewReader(f)

		c, err := csvReader.ReadAll()
		if err != nil {
			panic(err)
		}

		reportContent = append(reportContent, c...)
		f.Close()
	}

	reportFileName := fmt.Sprintf("%s/report.csv", monthFolder)
	reportFile, err := os.Create(reportFileName)
	if err != nil {
		panic(err)
	}
	defer reportFile.Close()

	csvWriter := csv.NewWriter(reportFile)
	if err := csvWriter.Write(
		[]string{"Date", "Start Time", "Stop Time", "Work time", "Break Time"}); err != nil {
		panic(err)
	}

	if err := csvWriter.WriteAll(reportContent); err != nil {
		panic(err)
	}
	csvWriter.Flush()
	if csvWriter.Error() != nil {
		panic(csvWriter.Error())
	}

	println("Report created")
}
func (r *Reporter) WriteStatistics(t *tracker.Tracker) {
	currentDate := time.Now()

	fileutils.CreateDirIfNotExists(r.workDir)

	monthFolder := fmt.Sprintf("%s/%s/%s", r.workDir, fmt.Sprint(currentDate.Year()), currentDate.Month().String())
	fileutils.CreateDirIfNotExists(monthFolder)

	dayFileName := fmt.Sprintf("%s/%s.csv", monthFolder, fmt.Sprint(currentDate.Day()))
	fileutils.CreateFileIfNotExists(dayFileName)

	dayFile, err := os.OpenFile(dayFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer dayFile.Close()

	csvWriter := csv.NewWriter(dayFile)

	if err := csvWriter.Write(
		[]string{
			t.StartTime.Format("2006-01-02"),
			t.StartTime.Format("15:04:05"),
			t.StopTime.Format("15:04:05"),
			t.WorkedDuration().Round(time.Second).String(),
			t.BreaksDuration().Round(time.Second).String(),
		}); err != nil {
		panic(err)
	}

	csvWriter.Flush()
	if csvWriter.Error() != nil {
		panic(csvWriter.Error())
	}
}
