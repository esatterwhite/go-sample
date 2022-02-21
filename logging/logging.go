package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func New(level string, command string) Logger {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Str("command", command).Timestamp().Logger()
	logger.Level(lvl)

	return Logger{logger: logger}
}

func (self Logger) Trace() *zerolog.Event {
	return self.logger.Trace()
}

func (self Logger) Info() *zerolog.Event {
	return self.logger.Info()
}

func (self Logger) Debug() *zerolog.Event {
	return self.logger.Debug()
}

func (self Logger) With() zerolog.Context {
	return self.logger.With()
}

func (self Logger) Child(command string) Logger {
	sub := self.logger.With().Str("command", "fake").Logger()
	return Logger{logger: sub}
}
