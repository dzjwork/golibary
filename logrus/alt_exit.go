package logrus

import (
	"fmt"
	"os"
)

var handlers = []func(){}

// 运行单个处理程序，保证处理完后不触发系统级错误
func runHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "Error: Logrus exit handler error:", err)
		}
	}()

	handler()
}

// 运行所有的处理程序
func runHandlers() {
	for _, handler := range handlers {
		runHandler(handler)
	}
}

// Exit 执行所有logrus退出的处理程序，然后终止程序
func Exit(code int) {
	runHandlers()
	os.Exit(code)
}

// RegisterExitHandler 添加一个新的退出处理程序
func RegisterExitHandler(handler func()) {
	handlers = append(handlers, handler)
}

// DeferExitHandler 在退出处理程序头部注册一个退出处理程序
func DeferExitHandler(handler func()) {
	handlers = append([]func(){handler}, handlers...)
}
