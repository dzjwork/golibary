package logrus

import (
	"runtime"
	"sync"
	"time"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

var baseTimestamp time.Time

// 文本格式器
type TextFormatter struct {
	// 是否在输出颜色之前绕过tty检查
	ForceColors bool

	// 是否禁用颜色
	DisableColors bool

	// 是否引用所有制
	ForceQuote bool

	// 是否禁用所有值的引用
	DisableQuote bool

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool

	// 是否禁用时间戳日志记录
	DisableTimestamp bool

	// 是输出在tty时输出完整的时间戳
	FullTimestamp bool

	// 时间戳的格式
	TimestampFormat string

	// 是否对输出字段进行排序
	DisableSorting bool

	// 对键进行排序的函数
	SortingFunc func([]string)

	// Disables the truncation of the level text to 4 characters.
	DisableLevelTruncation bool

	// PadLevelText Adds padding the level text so that all the levels output at the same length
	// PadLevelText is a superset of the DisableLevelTruncation option
	PadLevelText bool

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	// 是否输出到终端
	isTerminal bool

	// 为所有支持的字段定义的新名称
	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	terminalInitOnce sync.Once

	// 日志级别名称的最大长度
	levelTextMaxLength int
}
