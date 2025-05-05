package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	log := fmt.Sprintf("[%s] - [%s] %s\n",
		entry.Time.Format("2006-01-02 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	return []byte(log), nil
}

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	logFilePath := filepath.Join(logDir, "application.log")
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	Log.SetOutput(io.MultiWriter(os.Stdout, file))
	Log.SetFormatter(&CustomFormatter{})
	Log.SetLevel(logrus.DebugLevel)
	Log.SetReportCaller(false)
}