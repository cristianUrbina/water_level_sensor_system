package testutils

import (
	"fmt"

	"go.uber.org/mock/gomock"
)

type MyTestReporter struct {
	T gomock.TestReporter
}

func (m *MyTestReporter) Errorf(format string, args ...interface{}) {
	if m.T == nil {
		panic(fmt.Sprintf(format, args...))
	}
	m.T.Errorf(format, args...)
}

func (m *MyTestReporter) Fatalf(format string, args ...interface{}) {
	if m.T == nil {
		panic(fmt.Sprintf(format, args...))
	}
	m.T.Fatalf(format, args...)
}
