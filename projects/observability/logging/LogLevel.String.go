package logger

func (level *LogLevel) String() string {
	switch *level {
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	case Trace:
		return "Trace"
	default:
		panic("unknown log level")
	}
}
