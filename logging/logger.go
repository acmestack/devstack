package logging

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 's Sugar A Sugar wraps the base Logger functionality in a slower, but less
// verbose, API. Any Logger can be converted to a SugaredLogger with its Sugar
// method.
//
// Unlike the Logger, the SugaredLogger doesn't insist on structured logging.
// For each log level, it exposes four methods:
//
//   - methods named after the log level for log.Print-style logging
//   - methods ending in "w" for loosely-typed structured logging
//   - methods ending in "f" for log.Printf-style logging
//   - methods ending in "ln" for log.Println-style logging
//
// For example, the methods for InfoLevel are:
//
//	Info(...any)           Print-style logging
//	Infow(...any)          Structured logging (read as "info with")
//	Infof(string, ...any)  Printf-style logging
//	Infoln(...any)         Println-style logging
type Logger struct {
	logr.Logger
	level LogLevel
	Sugar *zap.SugaredLogger
}

func InitLogger(level LogLevel) Logger {
	logger := initZapLogger(level)
	return Logger{
		Logger: zapr.NewLogger(logger),
		Sugar:  logger.Sugar(),
	}
}

// WithName returns a new Logger instance with the specified name element added
// to the Logger's name.  Successive calls with WithName append additional
// suffixes to the Logger's name.  It's strongly recommended that name segments
// contain only letters, digits, and hyphens (see the package documentation for
// more information).
func (l Logger) WithName(name string) Logger {
	logger := initZapLogger(l.level)

	return Logger{
		Logger: zapr.NewLogger(logger).WithName(name),
		Sugar:  l.Sugar,
	}
}

// WithValues returns a new Logger instance with additional key/value pairs.
// See Info for documentation on how key/value pairs work.
func (l Logger) WithValues(keysAndValues ...interface{}) Logger {
	l.Logger = l.Logger.WithValues(keysAndValues...)
	return l
}

func initZapLogger(level LogLevel) *zap.Logger {
	parseLevel, _ := zapcore.ParseLevel(string(level))
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(parseLevel))

	return zap.New(core, zap.AddCaller())
}
