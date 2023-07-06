package logging

// LogLevel defines a log level for system logs.
type LogLevel string

const (
	// LogLevelDebug defines the "debug" logging level.
	LogLevelDebug LogLevel = "debug"

	// LogLevelInfo defines the "Info" logging level.
	LogLevelInfo LogLevel = "info"

	// LogLevelWarn defines the "Warn" logging level.
	LogLevelWarn LogLevel = "warn"

	// LogLevelError defines the "Error" logging level.
	LogLevelError LogLevel = "error"
)
