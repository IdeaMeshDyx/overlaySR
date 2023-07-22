package srError

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

/*
  basicError  基础的错误类型，包含必要的信息
*/
// 核心 error 类型，抽像出不同类型的 error 行为
// 同时也是一开始第一个错误的全部类型
type basicError struct {
	// 导致错误的内部潜逃问题，处于问题构造链的深层次
	cause error
	// 所处文件以及文件执行相关代码的行数
	filepath string
	// 调用第三方包的错误
	msg string
	// 在msg 基础上添加的自己调用信息
	detail string
	// 错误信息栈
	stack errors.StackTrace
}

// 实现 Error 接口,针对于三类型的错误进行报错
// 每调用一次Error() 就会在错误堆栈中加入对应的错误信息，一共三种类型的错误：
// 1. 第三方调用正常，程序运行错误
// 2. 单纯程序运行错误
// 3. 第三方调用出现报错
func (be *basicError) Error() string {
	// 当前错误由于第三方调用包导致，是由于第三方问题导致了程序内的错误
	if be.detail != "" && be.cause != nil {
		return fmt.Sprintf("%v \n ===>  %v", be.detail, be.cause.Error())
	}
	// 当前错误由于程序内错误导致
	if be.detail != "" {
		return fmt.Sprintf("==%v \n ", be.detail)
	}
	// 第三方程序的错误
	// 这类情况是第三放调用没有成功运行，出现了问题；上述第一种是指第三方本身并没有出现问题，而是不正确的调用或者是参数错误导致他返回了不正确的结果
	if be.cause != nil {
		return fmt.Sprintf("**%v \n", be.cause.Error())
	}
	// 该阶段没有出现错误
	return ""
}

// 兼容其他的错误类型数据信息
func getOtherMsg(err error) string {
	type ErrorMsg interface {
		ErrorMsg() string
	}

	errorMsg, ok := err.(ErrorMsg)
	if ok {
		return errorMsg.ErrorMsg()
	} else {
		return err.Error()
	}
}

func (be *basicError) ErrorMsg() string {
	if be.msg != "" && be.cause != nil {
		return fmt.Sprintf("%v \ncause by ====> %v", be.msg, getOtherMsg(be.cause))
	}

	if be.msg != "" {
		return fmt.Sprintf("==%v", be.msg)
	}

	if be.cause != nil {
		return fmt.Sprintf("**%v", getOtherMsg(be.cause))
	}
	return ""
}

func (be *basicError) Cause() error {
	return be.cause
}

func (be *basicError) Unwrap() error {
	return be.cause
}

func (be *basicError) Format(s fmt.State, verb rune) {
	var stackTrace errors.StackTrace

	// 如果 error 堆栈为空，那么就往其中注入初始化的错误
	defer func() {
		if stackTrace != nil {
			stackTrace.Format(s, verb)
		}
	}()

	//依据不同的行为来执行格式化操作
	switch verb {
	case 'v':
		if s.Flag('+') {
			if be.detail != "" {
				_, _ = io.WriteString(s, be.detail)
			}
			if be.stack != nil {
				stackTrace = be.stack
			}
			if be.Cause() != nil {
				_, _ = fmt.Fprintf(s, "\n Cause by ===> %+v", be.Cause())
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, be.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", be.Error())
	default:
		_, _ = fmt.Fprintf(s, "unsupported format: %%!%c,use %%S:%s)", verb, be.Error())
	}
}

func (be *basicError) Msg() string {
	return be.msg
}

/*// myError type Basic Defines
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
}*/
