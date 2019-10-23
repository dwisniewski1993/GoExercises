package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

// Poziomy logowania
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

const defaultLogLevel = LevelInfo

// Logger reprezentuje obiekt rejestrujący, który obsługuje wpisy dziennika na podstawie bieżącego poziomu dziennika
type Logger struct {
	mu     sync.Mutex // for serialization
	prefix string     // prefix to write at beginning of each log entry
	Level  Level
	w      io.Writer    // writer for output
	buf    bytes.Buffer // internal buffer
}

// Tworzenie nowego loggera
func New(w io.Writer, prefix string) *Logger {
	return &Logger{w: w, prefix: prefix, Level: defaultLogLevel}
}

// Konsola tworzy nowy loger który zwraca wyjście do Stderr
var Console = New(os.Stderr, "")

func (l *Logger) Debug(v ...interface{}) {
	if LevelDebug < l.Level {
		return
	}
	l.WriteEntry(LevelDebug, fmt.Sprintln(v...))
}

func (l *Logger) Info(v ...interface{}) {
	if LevelInfo < l.Level {
		return
	}
	l.WriteEntry(LevelInfo, fmt.Sprintln(v...))
}

func (l *Logger) Warning(v ...interface{}) {
	if LevelWarning < l.Level {
		return
	}
	l.WriteEntry(LevelWarning, fmt.Sprintln(v...))
}

func (l *Logger) Error(v ...interface{}) {
	if LevelError < l.Level {
		return
	}
	l.WriteEntry(LevelError, fmt.Sprintln(v...))
}

func (l *Logger) Critical(v ...interface{}) {
	if LevelCritical < l.Level {
		return
	}
	l.WriteEntry(LevelCritical, fmt.Sprintln(v...))
}

// Ustawianie poziomu logowania
func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = lvl
}

// Pobieranie obecnego poziomu logowania
func (l *Logger) GetLevel() Level {
	return l.Level
}

// Zapisywanie komunikatu określonego poziomu do domyślnego writera
func (l *Logger) WriteEntry(lvl Level, msg string) error {
	l.w.Write([]byte(msg))
	return nil
}
