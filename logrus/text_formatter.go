package logrus

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
	"unicode/utf8"
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

	// 是否引用所有值
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

// 初始化TextFormatter会获取到日志级别的最大长度
func (f *TextFormatter) init(entry *Entry) {
	if entry.Logger != nil {
		f.isTerminal = checkIfTerminal(entry.Logger.Out)
	}

	for _, level := range AllLevels {
		levelTextLength := utf8.RuneCount([]byte(level.String()))
		if levelTextLength > f.levelTextMaxLength {
			f.levelTextMaxLength = levelTextLength
		}
	}
}

// 查看字符串是否需要引用
func (f *TextFormatter) needsQuoting(text string) bool {
	if f.ForceQuote {
		return true
	}

	if f.QuoteEmptyFields && len(text) == 0 {
		return true
	}

	if f.DisableQuote {
		return false
	}

	// 判断字符如果不是a-z、A-Z、0-9、-、.、_、/、@、^、+则返回true
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

// 将值加到缓冲区中
func (f *TextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !f.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}

}

// 将键加到缓冲区中
func (f *TextFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

// 打印日志是否使用颜色打印
func (f *TextFormatter) isColored() bool {
	isColored := f.ForceColors || (f.isTerminal && (runtime.GOOS != "windows"))

	if f.EnvironmentOverrideColors {
		switch force, ok := os.LookupEnv("CLICOLOR_FORCE"); {
		case ok && force != "0":
			isColored = true
		case ok && force == "0", os.Getenv("CLICOLOR") == "0":
			isColored = false
		}
	}
	return isColored && !f.DisableColors
}
