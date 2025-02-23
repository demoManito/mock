package inject

// Injector is the interface that New must implement.
type Injector interface {
	Register(key string, value any)
	Load(key string) (value any, ok bool)
	Delete(key string)
}
