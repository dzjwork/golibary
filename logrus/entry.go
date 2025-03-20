package logrus

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
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

// ErrorKey 错误消息的键名
var ErrorKey = "error"

type Entry struct {
	// 指向log对象
	Logger *Logger

	// 包含用户设置的所有字段
	Data Fields

	// 该节点创建时间
	Time time.Time

	// 日志的级别
	Level Level

	// 带包名的调用时方法，每条日志中记录文件名、函数和行号
	Caller *runtime.Frame

	// 要记录的消息
	Message string

	// 格式化好的消息
	Buffer *bytes.Buffer

	// 用户设置的上下文，钩子函数使用
	Context context.Context

	// 记录字段格式化错误
	err string
}

// NewEntry 创建一个新的Entry实例
func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		// 默认是三个字段+一个可选字段
		Data: make(Fields, 6),
	}
}

// Dup 拷贝当前Entry得到一个新的Entry
func (entry *Entry) Dup() *Entry {
	// 拷贝段字段信息
	data := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		data[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: data, Time: entry.Time, Context: entry.Context, err: entry.err}
}

// 格式化Entry中的数据
func (entry *Entry) Bytes() ([]byte, error) {
	return entry.Logger.Formatter.Format(entry)
}

// 将Entry中数据格式化后返回字符串
func (entry *Entry) String() (string, error) {
	serialized, err := entry.Bytes()
	if err != nil {
		return "", err
	}
	str := string(serialized)
	return str, nil
}

// WithFields 根据当前Entry创建一个新的Entry，并将字段加到新的Entry中
func (entry *Entry) WithFields(fields Fields) *Entry {
	// 将原先entry中的段信息拷贝过来
	data := make(Fields, len(entry.Data)+len(fields))
	for k, v := range entry.Data {
		data[k] = v
	}

	// 记录无法成功添加的字段
	fieldErr := entry.err
	// 新加字段的值不能是一个函数，如果是一个函数将无法添加到字段集中并进行记录
	for k, v := range fields {
		isErrField := false
		if t := reflect.TypeOf(v); t != nil {
			switch {
			case t.Kind() == reflect.Func, t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Func:
				isErrField = true
			}
		}
		if isErrField {
			tmp := fmt.Sprintf("can not add field %q", k)
			if fieldErr != "" {
				fieldErr = entry.err + ", " + tmp
			} else {
				fieldErr = tmp
			}
		} else {
			data[k] = v
		}
	}
	return &Entry{Logger: entry.Logger, Data: data, Time: entry.Time, err: fieldErr, Context: entry.Context}
}

// WithField 添加单个字段到Entry中获取一个新的Entry实例
func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return entry.WithFields(Fields{key: value})
}

// 将错误作为单个字段添加到Entry中
func (entry *Entry) WithError(err error) *Entry {
	return entry.WithField(ErrorKey, err)
}

// WithContext 根据上下文创建一个新的Entry
func (entry *Entry) WithContext(ctx context.Context) *Entry {
	// 将原先entry中的段信息拷贝过来
	dataCopy := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		dataCopy[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: dataCopy, Time: entry.Time, err: entry.err, Context: ctx}
}

// WithTime 根据时间创建一个新的Entry
func (entry *Entry) WithTime(t time.Time) *Entry {
	// 将原先entry中的段信息拷贝过来
	dataCopy := make(Fields, len(entry.Data))
	for k, v := range entry.Data {
		dataCopy[k] = v
	}
	return &Entry{Logger: entry.Logger, Data: dataCopy, Time: t, err: entry.err, Context: entry.Context}
}

// 获取包名
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

// 是否有Caller函数
func (entry Entry) HasCaller() (has bool) {
	return entry.Logger != nil && entry.Logger.ReportCaller && entry.Caller != nil
}

func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// 获取日志缓冲池
func (entry *Entry) getBufferPool() (pool BufferPool) {
	if entry.Logger.BufferPool != nil {
		return entry.Logger.BufferPool
	}
	return bufferPool
}

// 触发该Entry对应级别及以下级别的钩子函数
func (entry *Entry) fireHooks() {
	var tmpHooks LevelHooks
	entry.Logger.mu.Lock()
	// 拷贝出来钩子函数
	tmpHooks = make(LevelHooks, len(entry.Logger.Hooks))
	for k, v := range entry.Logger.Hooks {
		tmpHooks[k] = v
	}
	entry.Logger.mu.Unlock()

	// 触发该日志级别的所有钩子函数
	err := tmpHooks.Fire(entry.Level, entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

// 写出日志，如果写失败会在标准错误输出中输出错误原因
func (entry *Entry) write() {
	entry.Logger.mu.Lock()
	defer entry.Logger.mu.Unlock()
	serialized, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		return
	}
	if _, err := entry.Logger.Out.Write(serialized); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}
}

// 按照日志级别打印日志
func (entry *Entry) log(level Level, msg string) {
	var buffer *bytes.Buffer

	newEntry := entry.Dup()

	// 设置新Entry的创建时间
	if newEntry.Time.IsZero() {
		newEntry.Time = time.Now()
	}

	// 设置日志消息和日志级别
	newEntry.Level = level
	newEntry.Message = msg

	newEntry.Logger.mu.Lock()
	reportCaller := newEntry.Logger.ReportCaller
	bufPool := newEntry.getBufferPool()
	newEntry.Logger.mu.Unlock()

	if reportCaller {
		newEntry.Caller = getCaller()
	}

	// 触发对应的钩子函数
	newEntry.fireHooks()
	buffer = bufPool.Get()
	// 释放资源并将缓冲池的内存返还
	defer func() {
		newEntry.Buffer = nil
		buffer.Reset()
		bufPool.Put(buffer)
	}()
	buffer.Reset()
	newEntry.Buffer = buffer

	// 写出去日志
	newEntry.write()

	newEntry.Buffer = nil

	if level <= PanicLevel {
		panic(newEntry)
	}
}

// Log 根据日志Entry日志级别打印日志，Entry的日志级别必须大于指定的级别才会打印日志，否则不打印
func (entry *Entry) Log(level Level, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.log(level, fmt.Sprint(args...))
	}
}

// Trace 打印跟踪日志
func (entry *Entry) Trace(args ...interface{}) {
	entry.Log(TraceLevel, args...)
}

// Debug 打印debug日志
func (entry *Entry) Debug(args ...interface{}) {
	entry.Log(DebugLevel, args...)
}

// Print 打印普通日志
func (entry *Entry) Print(args ...interface{}) {
	entry.Info(args...)
}

// Info 打印普通日志
func (entry *Entry) Info(args ...interface{}) {
	entry.Log(InfoLevel, args...)
}

// Warn 打印警告日志
func (entry *Entry) Warn(args ...interface{}) {
	entry.Log(WarnLevel, args...)
}

// Warning 警告日志
func (entry *Entry) Warning(args ...interface{}) {
	entry.Warn(args...)
}

// 警告日志
func (entry *Entry) Error(args ...interface{}) {
	entry.Log(ErrorLevel, args...)
}

// Fatal 致命错误日志
func (entry *Entry) Fatal(args ...interface{}) {
	entry.Log(FatalLevel, args...)
	entry.Logger.Exit(1)
}

// Panic 系统错误日志
func (entry *Entry) Panic(args ...interface{}) {
	entry.Log(PanicLevel, args...)
}

// Logf 按照指定的格式打印日志，Entry的日志级别必须大于指定的级别才会打印日志，否则不打印
func (entry *Entry) Logf(level Level, format string, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.Log(level, fmt.Sprintf(format, args...))
	}
}

// Tracef 格式化打印跟踪日志
func (entry *Entry) Tracef(format string, args ...interface{}) {
	entry.Logf(TraceLevel, format, args...)
}

// Debugf 格式化打印debug日志
func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.Logf(DebugLevel, format, args...)
}

// Infof 格式化打印普通日志
func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.Logf(InfoLevel, format, args...)
}

// Printf 格式化打印普通日志
func (entry *Entry) Printf(format string, args ...interface{}) {
	entry.Infof(format, args...)
}

// Warnf 格式化打印警告日志
func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.Logf(WarnLevel, format, args...)
}

// Warningf 格式化打印警告日志
func (entry *Entry) Warningf(format string, args ...interface{}) {
	entry.Warnf(format, args...)
}

// Errorf 格式化打印错误日志
func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.Logf(ErrorLevel, format, args...)
}

// 格式化打印
func (entry *Entry) Fatalf(format string, args ...interface{}) {
	entry.Logf(FatalLevel, format, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicf(format string, args ...interface{}) {
	entry.Logf(PanicLevel, format, args...)
}

// 获取到有换行符的字符
func (entry *Entry) sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}

// 打印可以换行的日志
func (entry *Entry) Logln(level Level, args ...interface{}) {
	if entry.Logger.IsLevelEnabled(level) {
		entry.Log(level, entry.sprintlnn(args...))
	}
}

func (entry *Entry) Traceln(args ...interface{}) {
	entry.Logln(TraceLevel, args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.Logln(DebugLevel, args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.Logln(InfoLevel, args...)
}

func (entry *Entry) Println(args ...interface{}) {
	entry.Infoln(args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.Logln(WarnLevel, args...)
}

func (entry *Entry) Warningln(args ...interface{}) {
	entry.Warnln(args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.Logln(ErrorLevel, args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.Logln(FatalLevel, args...)
	entry.Logger.Exit(1)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.Logln(PanicLevel, args...)
}
