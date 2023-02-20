package logger

import (
	"testing"

	"bitbucket.org/ayopop/ct-logger/constant"
)

func TestInit(t *testing.T) {
	type args struct {
		service  string
		env      string
		logLevel string
	}
	mockLogger := Init("test", "test", "error")
	tests := []struct {
		want *Logger
		args args
		name string
	}{
		{
			name: "OK",
			args: args{
				service:  "test",
				env:      "test",
				logLevel: "",
			},
			want: mockLogger,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := Init(tt.args.service, tt.args.env, tt.args.logLevel)
				if got.env != tt.args.env {
					t.Errorf("Init() = %v, want %v", got, tt.want)
				}
				if got.service != tt.args.service {
					t.Errorf("Init() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
func TestLogger_Debug_Dev(t *testing.T) {
	mockLogger := Init("test", constant.EnvDev, "error")
	mockLogger.Debug("test")
}

func TestLogger_Debug_Stage(t *testing.T) {
	mockLogger := Init("test", constant.EnvStage, "error")
	mockLogger.Debug("test")
}

func TestLogger_Debug_Prod(t *testing.T) {
	mockLogger := Init("test", "prod", "error")
	mockLogger.Debug("test")
}

func TestLogger_Debugf_Dev(t *testing.T) {
	mockLogger := Init("test", constant.EnvDev, "error")
	mockLogger.Debugf("%s", "test")
}

func TestLogger_Debugf_Stage(t *testing.T) {
	mockLogger := Init("test", constant.EnvStage, "error")
	mockLogger.Debugf("%s", "test")
}

func TestLogger_Debugf_Prod(t *testing.T) {
	mockLogger := Init("test", "prod", "debug")
	mockLogger.Debugf("%s", "test")
}

func TestLogger_Error(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Error("test")
}

func TestLogger_Errorf(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Errorf("%s", "test")
}

func TestLogger_Info(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Info("test")
}

func TestLogger_Infof(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Infof("%s", "test")
}

func TestLogger_Warn(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Warn("test")
}

func TestLogger_Warnf(t *testing.T) {
	mockLogger := Init("test", "test", "error")
	mockLogger.Warnf("%s", "test")
}
