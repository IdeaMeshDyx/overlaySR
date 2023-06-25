package srError

import (
	"github.com/pkg/errors"
	"runtime"
	"strconv"
	"strings"
)

// 错误堆栈的深度
const depth = 20

// 获取堆栈信息
func callers() errors.StackTrace {
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	stack := pcs[0:n]
	// 创建error的平台
	st := make([]errors.Frame, len(stack))

	for i := 0; i < len(st); i++ {
		st[i] = errors.Frame((stack)[i])
	}
	return st
}

func callerFuncInfo() string {
	pc, fileName, line, _ := runtime.Caller(2)

	f := runtime.FuncForPC(pc)

	callerFuncName := f.Name()
	callerFuncName = callerFuncName[strings.LastIndex(callerFuncName, ".")+1:]
	fileName = fileName[strings.LastIndex(fileName, "/")+1:]
	return callerFuncName + "(" + fileName + ":" + strconv.Itoa(line) + ")"
}
