package logrus

import (
	"io"
	"os"
	"sync"
)

type LogFunction func() []interface{}

type MutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

// 获取锁
func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

// 释放锁
func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

// 锁失效
func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

// 表示退出程序的函数
type exitFunc func(int)

type Logger struct {
	// 输出日志的流
	Out io.Writer
	// 日志输出时触发的狗子
	Hooks LevelHooks
	// 输出日志前的格式化程序
	Formatter Formatter

	// 是否记录呼叫着信息
	ReportCaller bool

	// 日志的级别
	Level Level
	// 为了保证同步写日志的锁
	mu MutexWrap
	// Reusable empty entry
	entryPool sync.Pool
	// 退出应用程序的函数,默认是os.exit()
	ExitFunc exitFunc
	// 日志缓冲池
	BufferPool BufferPool
}

// 创建一个新的Logger实例
func New() *Logger {
	return &Logger{
		Out:          os.Stderr,
		Formatter:    new(TextFormatter),
		Hooks:        make(LevelHooks),
		Level:        InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}
