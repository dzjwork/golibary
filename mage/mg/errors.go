package mg

import (
	"errors"
	"fmt"
)

// 表示致命的错误
type fatalErr struct {
	code int
	error
}

// 查看退出的状态
func (f fatalErr) ExitStatus() int {
	return f.code
}

type exitStatus interface {
	ExitStatus() int
}

// 根据错误信息获取一个致命错误实例
func Fatal(code int, args ...interface{}) error {
	return fatalErr{
		code:  code,
		error: errors.New(fmt.Sprint(args...)),
	}
}

// 根据错信息格式化获取一个致命错误实例
func Fatalf(code int, format string, args ...interface{}) error {
	return fatalErr{
		code:  code,
		error: fmt.Errorf(format, args...),
	}
}

// 尝试获取错误的状态码（尝试将错误转为致命错误类型）
func ExitStatus(err error) int {
	if err == nil {
		return 0
	}
	exit, ok := err.(exitStatus)
	if !ok {
		return 1
	}
	return exit.ExitStatus()
}
