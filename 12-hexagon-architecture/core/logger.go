package core

import (
	"fmt"
	"log"
)

type LogLevel int64

const (
	ERROR LogLevel = 0
	WARN           = 1
	INFO           = 2
	DEBUG          = 3
)

type Logger struct {
	level LogLevel
}

func (s Logger) Info(text string) {
	if s.level >= INFO {
		log.Println(fmt.Sprintf("[INFO] %v", text))
	}
}

func (s Logger) Debug(text string) {
	if s.level >= DEBUG {
		log.Println(fmt.Sprintf("[DEBUG] %v", text))
	}
}

func (s Logger) Warn(text string) {
	if s.level >= WARN {
		log.Println(fmt.Sprintf("[WARN] %v", text))
	}
}

func (s Logger) Error(text string) {
	if s.level >= ERROR {
		log.Println(fmt.Sprintf("[ERROR] %v", text))
	}
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
	}
}
