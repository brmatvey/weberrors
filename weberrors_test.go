package weberrors_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/brmatvey/weberrors"
)

func TestNew(t *testing.T) {
	err := (error)(nil)
	err = weberrors.New(http.StatusNotFound)
	castedErr, ok := err.(*weberrors.Error)
	if !ok {
		t.Fatal("must be casted")
	}
	if len(castedErr.Stack) == 0 {
		t.Fatal("must be not 0")
	}
	if castedErr.Stack[0].FileName != "weberrors_test.go" {
		t.Fatal("must be weberrors_test.go")
	}
	if castedErr.Stack[0].LineNumber != 14 {
		t.Fatal("must be 14")
	}
	if !strings.Contains(castedErr.Stack[0].FunctionName, "TestNew") {
		t.Fatal("must contain")
	}
	if castedErr.Code != http.StatusNotFound {
		t.Fatal("must be not found")
	}
}

func TestNewWithArgs(t *testing.T) {
	expectedMessage := fmt.Sprintf("must %d %d %d", 1, 2, 3)
	err := (error)(nil)
	err = weberrors.NewWithArgs(http.StatusNotFound, "must %d %d %d", 1, 2, 3)
	castedErr, ok := err.(*weberrors.Error)
	if !ok {
		t.Fatal("must be casted")
	}
	if len(castedErr.Stack) == 0 {
		t.Fatal("must be not 0")
	}
	if castedErr.Stack[0].FileName != "weberrors_test.go" {
		t.Fatal("must be weberrors_test.go")
	}
	if castedErr.Stack[0].LineNumber != 39 {
		t.Fatal("must be 39")
	}
	if !strings.Contains(castedErr.Stack[0].FunctionName, "TestNewWithArgs") {
		t.Fatal("must contain")
	}
	if castedErr.Code != http.StatusNotFound {
		t.Fatal("must be not found")
	}
	if castedErr.Message != expectedMessage {
		t.Fatal("must be eq" + expectedMessage)
	}
}
