package utils

import "github.com/fatih/color"

type logFunc func(format string, a ...interface{})

type log struct {
	Info    logFunc
	Success logFunc
	Warn    logFunc
	Error   logFunc
	Debug   logFunc
}

var Log = log{
	Info:    color.White,
	Success: color.Green,
	Warn:    color.Yellow,
	Error:   color.Red,
	Debug:   color.Cyan,
}
