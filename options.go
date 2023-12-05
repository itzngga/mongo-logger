package logger

import (
	"io"
	"os"
)

// ref: https://github.com/gofiber/fiber/blob/v2.51.0/middleware/csrf/config.go
// special thanks!!!

type Level int

const (
	LevelInfo   = iota
	LevelError  = iota
	LevelSilent = iota
)

type Options struct {
	Writer io.Writer
	Level  Level
	Colors bool
}

var OptionsDefault = Options{
	Writer: os.Stdout,
	Level:  LevelInfo,
	Colors: true,
}

func optionsDefault(options ...Options) Options {
	if len(options) < 1 {
		return OptionsDefault
	}

	opt := options[0]
	if opt.Writer == nil {
		opt.Writer = os.Stdout
	}

	return opt
}
