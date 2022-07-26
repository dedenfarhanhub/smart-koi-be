package logging

import (
	"flag"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger() Logger {
	logger := log.New()
	f, err := os.OpenFile("api.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("cannot open 'testlogfile', (%s)", err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(f)
	return Logger{
		Log: logger,
	}
}

func (l Logger) Logging(c echo.Context) *log.Entry {
	logs := l.Log

	if c == nil {
		return logs.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logs.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func (l Logger) LogEntry(request, response interface{}) *log.Entry {
	logs := l.Log

	return logs.WithFields(log.Fields{
		"at":       time.Now().Format("2006-01-02 15:04:05"),
		"request":  request,
		"response": response,
	})
}
