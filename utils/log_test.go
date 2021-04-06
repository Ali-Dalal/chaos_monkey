package utils

import "testing"

func TestLogError(t *testing.T) {
	InitLogrusLog()
	// test will fail and exit if anything goes wrong
	LogError("test error log")
}

func TestLogInfo(t *testing.T) {
	InitLogrusLog()
	LogInfo(" test info log")
}