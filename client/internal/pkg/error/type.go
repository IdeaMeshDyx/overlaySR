package srError

import (
	"fmt"
	"time"
)

// myError type Basic Defines
// @Description:
type funcReturnError struct {
	text     string
	funcName string
}

func newFuncReturnError(text string, funcName string) *funcReturnError {
	return &funcReturnError{text: text, funcName: funcName}
}

func (error funcReturnError) Error() string {
	return fmt.Sprintf("Error: Time: %v,FuncName %v ,Error: %v ", time.Now().Format("2006-01-02 15:04:05"), error.funcName, error.text)
}
