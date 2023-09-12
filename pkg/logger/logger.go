package logger

import (
	"github.com/dmitriyGarden/test-user-service/model"
	log "github.com/sirupsen/logrus"
)

func New() model.ILogger {
	l := log.New()
	l.SetLevel(log.DebugLevel)
	return l
}
