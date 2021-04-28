package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Log returns a logrus.Logger, initialising it if necessary
func Log() *logrus.Logger {
	if log != nil {
		return log
	}
	lvlStr, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvlStr = "INFO"
	}
	lvl, err := logrus.ParseLevel(lvlStr)
	if err != nil {
		panic(err)
	}
	log = logrus.New()
	log.SetLevel(lvl)
	return log
}
