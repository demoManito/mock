package mock_test

import (
	"testing"

	"github.com/demoManito/mock"
	"github.com/demoManito/mock/close"
)

var (
	testVar = "test_mock"
)

type testInjector struct {
	kv map[string]any
}

func (i *testInjector) Register(key string, value any) {
	i.kv[key] = value
}

func (i *testInjector) Load(key string) (value any, ok bool) {
	v, ok := i.kv[key]
	return v, ok
}

func (*testInjector) Delete(string) {
	panic("implement me")
}

func testMockInject() mock.MFunc {
	return func(m *mock.Mock) {
		testVar = "test_mock_mock"
		m.Injector().Register("result", "success")
		m.AppendCloser(close.NewMockCloser(func() error {
			m.Log("test_mock_1 closing")
			return nil
		}))
	}
}

func testMockGlobal() mock.MFunc {
	return func(m *mock.Mock) {
		testVar = "test_mock_mock"
		m.AppendCloser(close.NewMockCloser(func() error {
			m.Log("test_mock_2 closing")
			return nil
		}))
	}
}

func TestMock(t *testing.T) {
	injector := &testInjector{kv: map[string]any{"result": "failed"}}
	defer mock.New(t, &mock.Option{Injector: injector}).Run(testMockInject(), testMockGlobal()).Close()
	defer mock.New(t).Run(testMockGlobal()).Close()

	if injector.kv["result"] != "success" {
		t.Errorf("expect success, but got %s", injector.kv["result"])
	}
	if testVar != "test_mock_mock" {
		t.Errorf("expect test_mock_mock, but got %s", testVar)
	}
}
