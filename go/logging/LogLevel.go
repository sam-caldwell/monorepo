package logger

type LogLevel uint8

const (
	Fatal LogLevel = iota
	Error
	Warn
	Info
	Debug
	Trace
)
