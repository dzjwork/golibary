package spew

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ConfigState struct {
	// 缩进符号
	Indent string

	// 解析的最大栈深度，0表示没限制
	MaxDepth int

	// DisableMethods specifies whether or not error and Stringer interfaces are
	// invoked for types that implement them.
	DisableMethods bool

	// DisablePointerMethods specifies whether or not to check for and invoke
	// error and Stringer interfaces on types which only accept a pointer
	// receiver when the current type is not a pointer.
	//
	// NOTE: This might be an unsafe action since calling one of these methods
	// with a pointer receiver could technically mutate the value, however,
	// in practice, types which choose to satisify an error or Stringer
	// interface with a pointer receiver should not be mutating their state
	// inside these interface methods.  As a result, this option relies on
	// access to the unsafe package, so it will not have any effect when
	// running in environments without access to the unsafe package such as
	// Google App Engine or with the "safe" build tag specified.
	DisablePointerMethods bool

	// 是否禁止打印指针地址
	DisablePointerAddresses bool

	// 是否禁止打印容量值
	DisableCapacities bool

	// ContinueOnMethod specifies whether or not recursion should continue once
	// a custom error or Stringer interface is invoked.  The default, false,
	// means it will print the results of invoking the custom error or Stringer
	// interface and return immediately instead of continuing to recurse into
	// the internals of the data type.
	//
	// NOTE: This flag does not have any effect if method invocation is disabled
	// via the DisableMethods or DisablePointerMethods options.
	ContinueOnMethod bool

	// SortKeys specifies map keys should be sorted before being printed. Use
	// this to have a more deterministic, diffable output.  Note that only
	// native types (bool, int, uint, floats, uintptr and string) and types
	// that support the error or Stringer interfaces (if methods are
	// enabled) are supported, with other types sorted according to the
	// reflect.Value.String() output which guarantees display stability.
	SortKeys bool

	// SpewKeys specifies that, as a last resort attempt, map keys should
	// be spewed to strings and sorted by those strings.  This is only
	// considered if SortKeys is true.
	SpewKeys bool
}

// 顶级函数的活动配置
var Config = ConfigState{Indent: " "}

// 格式化输出错误
func (c *ConfigState) Errorf(format string, a ...interface{}) (err error) {
	return fmt.Errorf(format, c.convertArgs(a)...)
}

// 直接输出到指定位置
func (c *ConfigState) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, c.convertArgs(a)...)
}

// 格式化输出到指定为止
func (c *ConfigState) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, c.convertArgs(a)...)
}

// 换行输出到指定位置
func (c *ConfigState) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, c.convertArgs(a)...)
}

// 直接打印
func (c *ConfigState) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(c.convertArgs(a)...)
}

// 格式化打印
func (c *ConfigState) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, c.convertArgs(a)...)
}

// 换行打印
func (c *ConfigState) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(c.convertArgs(a)...)
}

// 直接打印
func (c *ConfigState) Sprint(a ...interface{}) string {
	return fmt.Sprint(c.convertArgs(a)...)
}

// 格式化打印
func (c *ConfigState) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, c.convertArgs(a)...)
}

// 换行打印
func (c *ConfigState) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(c.convertArgs(a)...)
}

// 新建一个格式化器
func (c *ConfigState) NewFormatter(v interface{}) fmt.Formatter {
	return newFormatter(c, v)
}

// 将数据格式化后输出到指定的位置
func (c *ConfigState) Fdump(w io.Writer, a ...interface{}) {
	fdump(c, w, a...)
}

// 将数据格式化后输出到控制台中
func (c *ConfigState) Dump(a ...interface{}) {
	fdump(c, os.Stdout, a...)
}

// 将格式化后的数据作为字符串返回
func (c *ConfigState) Sdump(a ...interface{}) string {
	var buf bytes.Buffer
	fdump(c, &buf, a...)
	return buf.String()
}

// 将所有参数转为对应的Formatter
func (c *ConfigState) convertArgs(args []interface{}) (formatters []interface{}) {
	formatters = make([]interface{}, len(args))
	for index, arg := range args {
		formatters[index] = newFormatter(c, arg)
	}
	return formatters
}

// 获取到一个具有默认配置的ConfigState实力
func NewDefaultConfig() *ConfigState {
	return &ConfigState{Indent: " "}
}
