package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger Primary logging wrapper
type Logger struct {
	logger zerolog.Logger
}

// Default instance of a Logger
var Default Logger = New("fatal", "default")

// New create a new logger
func New(level string, command string) Logger {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Str("command", command).Timestamp().Logger()

	return Logger{logger: logger.Level(lvl)}
}

// Print Wrapper method for generic debug
func (log Logger) Print(format string, v ...interface{}) {
	log.logger.Debug().Msgf(format, v...)
}

// Trace outputs a log line at trace the level
func (log Logger) Trace() *zerolog.Event {
	return log.logger.Trace()
}

func (log Logger) Info() *zerolog.Event {
	return log.logger.Info()
}

func (log Logger) Debug() *zerolog.Event {
	return log.logger.Debug()
}

func (log Logger) With() zerolog.Context {
	return log.logger.With()
}

func (log Logger) Child(command string) Logger {
	sub := log.logger.With().Str("command", command).Logger()
	return Logger{logger: sub}
}
