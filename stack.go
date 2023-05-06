package weberrors

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
)

func NewStack() Stack {
	stack := make(Stack, 0)
	for i := 2; ; i++ {
		pc, fn, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, fn = filepath.Split(fn)

		functionName := ""
		if f := runtime.FuncForPC(pc); f != nil {
			functionName = f.Name()
		}

		stack = append(stack, StackItem{FileName: fn, LineNumber: line, FunctionName: functionName})
	}
	return stack
}

type Stack []StackItem

func (s Stack) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		buffer.WriteString(fmt.Sprintf("\n%s:%d %s", s[i].FileName, s[i].LineNumber, s[i].FunctionName))
	}
	return buffer.String()
}
