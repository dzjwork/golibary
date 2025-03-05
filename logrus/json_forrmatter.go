package logrus

import "runtime"

type fieldKey string

// 存储支持的默认字段名对应的新名称（用户自定义的名称）
type FieldMap map[fieldKey]string

// 获取到指定key对应的值，如果key不存在则返回key
func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}
	return string(key)
}

// 用于将日志解析为json
type JSONFormatter struct {
	// 设置时间戳的格式
	TimestampFormat string

	// 是否禁用时间戳打印
	DisableTimestamp bool

	// 是否进制HTML的转义
	DisableHTMLEscape bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey string

	// 存储支持的默认字段名对应的新名称
	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// 是否缩进json日志
	PrettyPrint bool
}

func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(Fields, len(entry.Data)+4)

	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
}
