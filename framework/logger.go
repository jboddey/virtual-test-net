package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(
		&nested.Formatter{})
}

func getLogger() *logrus.Logger {
	return log
}
