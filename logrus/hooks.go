package logrus

type Hook interface {
	// 获取该hook支持的所有日志级别
	Levels() []Level
	Fire(*Entry) error
}

// 存储每种级别的日志对应的hook处理器
type LevelHooks map[Level][]Hook

func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

func (hooks LevelHooks) Fire(level Level, entry *Entry) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(entry); err != nil {
			return err
		}
	}
	return nil
}
