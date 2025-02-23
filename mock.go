package mock

import (
	"github.com/demoManito/mock/close"
	"github.com/demoManito/mock/inject"
)

// MFunc is a function that takes a mock as an argument.
type MFunc func(*Mock)

// M is an interface that defines the methods of a mock.
type M interface {
	Log(...any)
	Logf(string, ...any)
	Error(...any)
	Errorf(string, ...any)
	Fatal(...any)
	Fatalf(string, ...any)
}

// Mock is a mock implementation of io.Closer.
type Mock struct {
	M
	injector inject.Injector
	closers  []close.Closer
}

// Option is a function that takes an injector and an IMockEnv as arguments.
type Option struct {
	Injector inject.Injector
}

// NewMock returns a new mock
func NewMock(m M, opts ...*Option) *Mock {
	mock := &Mock{
		M:       m,
		closers: make([]close.Closer, 0),
	}
	if len(opts) != 0 {
		mock.injector = opts[0].Injector
	}
	return mock
}

// Run runs the given functions.
func (m *Mock) Run(funcs ...MFunc) *Mock {
	for _, f := range funcs {
		f(m)
	}
	return m
}

// Injector returns the injector.
func (m *Mock) Injector() inject.Injector {
	return m.injector
}

// AppendCloser appends the given closers to the list of closers.
func (m *Mock) AppendCloser(closers ...close.Closer) {
	if len(closers) == 0 {
		return
	}
	m.closers = append(m.closers, closers...)
}

// Close closes all closers in the list.
func (m *Mock) Close() {
	for _, closer := range m.closers {
		if err := closer.Close(); err != nil {
			m.Errorf("close error:", err)
		}
	}
}
