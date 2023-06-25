package srError

import (
	"errors"
	"fmt"
)

func New(msg string) error {
	return &basicError{
		msg:    msg,
		detail: fmt.Sprintf("%v, %v", callerFuncInfo(), msg),
		stack:  callers(),
	}
}

func Newf(format string, args ...interface{}) error {
	return &basicError{
		msg:    fmt.Sprintf(format, args...),
		detail: fmt.Sprintf("%v， %v", callerFuncInfo(), fmt.Sprintf(format, args...)),
		stack:  callers(),
	}
}

// Wrap 使用传入的信息包装错误, 携带堆栈信息
// 如果传入的 err 已经有堆栈, 不再设置堆栈
// 如果传入的 err 为 nil, Wrap 将返回 nil
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	wrapErr := &basicError{
		cause:  err,
		msg:    msg,
		detail: fmt.Sprintf("%v, %v", callerFuncInfo(), msg),
	}
	var fd *basicError
	if errors.As(err, &fd) {
	} else {
		// 链路上没有同类型错误的时候，证明是首次包装, 添加上堆栈信息
		wrapErr.stack = callers()
	}
	return wrapErr
}

// Wrapf 使用 format 格式的信息包装错误, 携带堆栈信息
// 如果传入的 err 已经有堆栈, 不再设置堆栈
// 如果传入的 err 为 nil, Wrapf 将返回 nil
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	wrapErr := &basicError{
		cause:  err,
		msg:    fmt.Sprintf(format, args...),
		detail: fmt.Sprintf("%v, %v", callerFuncInfo(), fmt.Sprintf(format, args...)),
	}
	var fd *basicError
	if !errors.As(err, &fd) {
		// 链路上没有同类型错误的时候，证明是首次包装, 添加上堆栈信息
		wrapErr.stack = callers()
	}
	return wrapErr
}

// WrapWithCode 使用传入的信息包装错误, 携带堆栈信息
// 如果传入的 err 已经有堆栈, 不再设置堆栈
// 如果传入的 err 为 nil, WrapWithCode 将返回 nil
func WrapWithCode(err error, code int, msg string) error {
	if err == nil {
		return nil
	}
	wrapErr := &basicError{
		cause:  err,
		msg:    msg,
		detail: fmt.Sprintf("%v, %v", callerFuncInfo(), msg),
	}
	var fd *basicError
	if !errors.As(err, &fd) {
		// 链路上没有同类型错误的时候，证明是首次包装, 添加上堆栈信息
		wrapErr.stack = callers()
	}
	return wrapErr
}

// WrapWithCodef 使用 format 格式的信息包装错误, 携带堆栈信息
// 如果传入的 err 已经有堆栈, 不再设置堆栈
// 如果传入的 err 为 nil, WrapWithCodef 将返回 nil
func WrapWithCodef(err error, code int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	wrapErr := &basicError{
		cause:  err,
		msg:    fmt.Sprintf(format, args...),
		detail: fmt.Sprintf("%v, %v", callerFuncInfo(), fmt.Sprintf(format, args...)),
	}
	var fd *basicError
	if !errors.As(err, &fd) {
		// 链路上没有同类型错误的时候，证明是首次包装, 添加上堆栈信息
		wrapErr.stack = callers()
	}
	return wrapErr
}

// Msg 获取 msg
func Msg(e error) string {
	if e == nil {
		return ""
	}
	const unknownMsg = "unknown error"
	err, ok := e.(*basicError)
	if !ok {
		return unknownMsg
	}
	return err.Msg()
}

// Cause 如果可能，Cause 返回最内层的错误。
// 如果错误实现了以下内容，那么它就有一个原因
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// 如果错误没有实现原因，则原始错误将被返回。
// 如果错误为 nil，则返回 nil 而不进行进一步处理
func Cause(err error) error {
	type causer interface {
		Cause() error
	}
	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		if cause.Cause() == nil {
			break
		}
		err = cause.Cause()
	}
	return err
}

// Is 方便使用者, 无需为了使用 Is 引入 errors 标准库, 定义别名
func Is(err, target error) bool { return errors.Is(err, target) }

// As 方便使用者, 无需为了使用 As 引入 errors 标准库, 定义别名
func As(err error, target interface{}) bool { return errors.As(err, target) }

// Unwrap 方便使用者, 无需为了使用 Unwrap 引入 errors 标准库, 定义别名
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

/*// handle it after main thread by using defer
func funcErrHandler(originErr error, funcname string) error {
	defer func() {
		// sleep 600 seconds
		for i := 0; i < 600; i++ {
			time.Sleep(time.Second)
		}
	}()

	// combine the current error and funcname
	err := newFuncReturnError(originErr.Error(), funcname)
	fmt.Printf("\033[1;31m %s \033[0m \033[32m Program will exit in 10 mins \033[0m \n", err.Error())
	return err
}*/
