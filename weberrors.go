package weberrors

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

type Error struct {
	Code    int
	Message string
	Stack   Stack
}

func (e *Error) ErrorCode() int { return e.Code }
func (e *Error) Error() string  { return e.Message }

type StackItem struct {
	FileName     string
	LineNumber   int
	FunctionName string
}

func New(code int) *Error {

	stack := NewStack()
	logger.Printf("Error has occurred, call stack:\n%s", stack.String())
	return &Error{
		Code:  code,
		Stack: stack,
	}
}

func NewWithArgs(code int, message string, args ...interface{}) *Error {

	decoratedMessage := ""
	if len(args) == 0 {
		decoratedMessage = message
	} else {
		decoratedMessage = fmt.Sprintf(message, args...)
	}

	stack := NewStack()
	logger.Printf("Error [%s] has occurred, call stack:\n%s", decoratedMessage, stack.String())
	return &Error{
		Code:    code,
		Message: decoratedMessage,
		Stack:   stack,
	}
}
