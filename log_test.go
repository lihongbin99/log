package log

import "testing"

func TestMyUtils(t *testing.T) {
	Error("error")
	Warn("warn")
	Info("info")
	Debug("debug")
	Trace("trace")

	ChangeLevel("Trace")

	Error("error")
	Warn("warn")
	Info("info")
	Debug("debug")
	Trace("trace")
}
