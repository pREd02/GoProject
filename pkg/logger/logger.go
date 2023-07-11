package logger

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func GetLogger(level string, file io.Writer) Logger {
	var instance Logger
	var once sync.Once
	once.Do(func() {
		logrusLevel, err := logrus.ParseLevel(level)
		if err != nil {
			log.Fatalln(err)
		}

		l := logrus.New()
		l.Formatter = &logrus.JSONFormatter{}

		mw := io.MultiWriter(os.Stdout, file)
		l.SetOutput(mw)
		l.SetLevel(logrusLevel)

		instance = Logger{logrus.NewEntry(l)}
	})

	return instance
}
