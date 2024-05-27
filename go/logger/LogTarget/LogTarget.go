package LogTarget

// LogTarget - The target to which logs will be written
type LogTarget interface {
	Configure(cfg ConfigureTarget)
}
