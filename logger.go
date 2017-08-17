package vidar

import (
	"fmt"
	"log"
	"os"
	"sync"
)

const (
	DefaultLevel  = LevelInfo
	DefaultOutput = os.Stdout
	DefaultFlag   = log.Ldate | log.Ltime | log.Lshortfile
)

const (
	LevelError = itoa
	LevelWarning
	LevelInfo
	LevelDebug
)

var levelName = map[string]int{
	"Error":   LevelError,
	"Warning": LevelWarning,
	"Info":    LevelInfo,
	"Debug":   LevelDebug,
}

type Logger struct {
	level int
	err   *log.Logger
	war   *log.Logger
	inf   *log.Logger
	deb   *log.Logger
}

func (l *Logger) Error(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}

	l.err.Output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > log.level {
		return
	}

	l.war.Output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(format string, v ...interface{}) {
	if LevelInfo > log.level {
		return
	}

	l.logger.SetPrefix("[INFO]: ")
	l.inf.Output(3, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > log.level {
		return
	}

	l.logger.SetPrefix("[DEBUG]: ")
	l.deb.Output(3, fmt.Sprintf(format, v...))
}

var mutex = new(sync.RWMutex)

func (l *Logger) SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()
	l.level = level
}

func (l *Logger) SetLevelByName(level string) {
	mutex.Lock()
	defer mutex.Unlock()
	l.level = levelName[level]
}

func (l *Logger) Level() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return l.level
}

func NewLogger() *Logger {
	l := &Logger{
		level: DefaultLevel,
	}

	l.err = log.New(DefaultOutput, "[ERROR]: ", DefaultFlag)
	l.war = log.New(DefaultOutput, "[WARNING]: ", DefaultFlag)
	l.inf = log.New(DefaultOutput, "[INFO]: ", DefaultFlag)
	l.deb = log.New(DefaultOutput, "[DEBUG]: ", DefaultFlag)

	return l
}
