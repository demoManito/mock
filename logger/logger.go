package logger

import (
	"log"
	"os"
)

type MockLog struct {
	*log.Logger
}

func NewMockLog() *MockLog {
	// Initialize an underlying logger to avoid nil pointer dereference when calling Print/Printf
	return &MockLog{Logger: log.New(os.Stdout, "", log.LstdFlags)}
}

func (m *MockLog) Log(args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Print(args...)
}

func (m *MockLog) Logf(format string, args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Printf(format, args...)
}

func (m *MockLog) Error(args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Print(args...)
}

func (m *MockLog) Errorf(format string, args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Printf(format, args...)
}

func (m *MockLog) Fatal(args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Logger.Fatal(args...)
}

func (m *MockLog) Fatalf(format string, args ...any) {
	if m == nil || m.Logger == nil {
		return
	}
	m.Logger.Fatalf(format, args...)
}
