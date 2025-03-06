package logrus

type Hook interface {
	// 获取该hook支持的所有日志级别
	Levels() []Level
	Fire(*Entry) error
}

// 存储每种级别的日志对应的hook处理器
type LevelHooks map[Level][]Hook

// 在所有日志级别中注册钩子函数
func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

// 触发对应日志级别下的所有钩子函数
func (hooks LevelHooks) Fire(level Level, entry *Entry) error {
	for _, hook := range hooks[level] {
		// 调用对应的触发函数
		if err := hook.Fire(entry); err != nil {
			return err
		}
	}
	return nil
}
