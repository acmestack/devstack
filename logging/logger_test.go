package logging

import (
	"errors"
	"testing"
)

func TestInitLogger(t *testing.T) {
	type args struct {
		level LogLevel
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "init log",
			args: args{level: "error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := InitLogger(tt.args.level)
			logger.Error(errors.New("error"), "error occur")

			logger.WithName("sys").WithValues("key", "value").Info("msg", "infok", "infovalue")
		})
	}
}
