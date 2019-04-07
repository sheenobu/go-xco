package xco

import "testing"

type manyerrors interface {
	Errors() []error
}

func TestErrorCast(t *testing.T) {
	var err error = &multiError{}
	e, ok := err.(manyerrors)
	if !ok {
		t.Errorf("can't cast multiError to Errors interface")
	}
	e.Errors()
}
