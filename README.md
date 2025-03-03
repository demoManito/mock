## mock

> Provides a unified interface for automatic call testing, which facilitates unified management of mock data, generation of mock data, storage of mock data, and clearing of mock data


### Install

```bash
go get -u github.com/demoManito/mock
```

### Usage

```go
package xxx

import (
    "fmt"
	"testing"

    "github.com/demoManito/mock"
)

func TestXxx(t *testing.T) {
    // Create a new mock object
    m := mock.New(t).Run(MockDB(), MockRedis())
	defer m.Close()
	
	// Use the mock object to test the function
    // ...
}
```
