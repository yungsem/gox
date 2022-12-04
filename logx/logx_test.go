package logx

import (
	"errors"
	"testing"
)

func TestNewLog(t *testing.T) {
	log := NewStdoutLog("WARN")

	log.Warn("test warn string log")
	log.WarnF("test %s log", "warn")
	log.Debug("test debug string log")
	log.DebugF("test %s log", "debug")
	log.Info("test info string log")
	log.InfoF("test %s log", "info")
	log.Error("test error string log")
	log.ErrorF("test %s log", "error")

	err := errors.New("test error")
	log.ErrorE(err)
}
