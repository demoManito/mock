package close

// Closer .
type Closer interface {
	Close() error
}

// NewMockCloser returns a Closer that wraps the given function.
func NewMockCloser(f func() error) *MockCloser {
	return &MockCloser{
		close: f,
	}
}

// MockCloser is an interface that wraps the Close method.
type MockCloser struct {
	close func() error
}

// Close calls the wrapped function.
func (c *MockCloser) Close() error {
	return c.close()
}
