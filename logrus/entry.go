package logrus

import (
	"sync"
)

var (
	// 首次使用时缓存的限定包名
	logrusPackage string

	// 跟踪以报告调用方法时调用堆栈的位置
	minimumCallerDepth int

	// 用于呼叫着信息初始化
	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

func init() {
	// 包名缓存前从堆栈底部开始
	minimumCallerDepth = 1
}

var ErrorKey = "error"

type Entry struct {
	Logger *Logger

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context

	// err may contain a field formatting error
	err string
}