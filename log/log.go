package log

import "fmt"

const (
	Debug = 0
	Info  = 1
	Error = 2
)

type Logger struct {
	Level uint8
}

func NewLogger(level uint8) *Logger {
	return &Logger{
		Level: level,
	}
}

func (l *Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l *Logger) Error(msg string) {
	fmt.Println(msg)
}
