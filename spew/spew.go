package spew

import (
	"fmt"
	"io"
)

// 格式化打印错误
func Errorf(format string, a ...interface{}) (err error) {
	return fmt.Errorf(format, convertArgs(a)...)
}

// 直接打印到指定位置
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, convertArgs(a)...)
}

// 格式化打印到指定位置
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, convertArgs(a)...)
}

// 换行打印到指定位置
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, convertArgs(a)...)
}

// 直接打印
func Print(a ...interface{}) (n int, err error) {
	return fmt.Print(convertArgs(a)...)
}

// 格式化打印
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, convertArgs(a)...)
}

// 换行打印
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(convertArgs(a)...)
}

// 直接打印
func Sprint(a ...interface{}) string {
	return fmt.Sprint(convertArgs(a)...)
}

// 格式化打印
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, convertArgs(a)...)
}

// 换行打印
func Sprintln(a ...interface{}) string {
	return fmt.Sprintln(convertArgs(a)...)
}

// 将所有参数转为对应的Formatter
func convertArgs(args []interface{}) (formatters []interface{}) {
	formatters = make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = NewFormatter(arg)
	}
	return formatters
}
