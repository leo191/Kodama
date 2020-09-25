package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger(path string) *StandardLogger {
	var baseLogger = logrus.New()
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		baseLogger.Fatal(err)
	}

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}
	standardLogger.SetOutput(file)
	return standardLogger
}

// Log method stream out as per proper level
func (l *StandardLogger) Log(level int8, msg string, err error) {
	switch {
	case err == nil:
		l.Infoln(msg)
	case level == 1:
		l.Warningf("Issue: %s", err)
	case level == 2:
		l.Errorf("Issue: %s", err)
	case level == 3:
		l.Fatalf("Issue: %s", err)
	}
}
