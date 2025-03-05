package mg

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// 颜色对应的颜色码
var ansiColor = map[Color]string{
	Black:         "\u001b[30m",
	Red:           "\u001b[31m",
	Green:         "\u001b[32m",
	Yellow:        "\u001b[33m",
	Blue:          "\u001b[34m",
	Magenta:       "\u001b[35m",
	Cyan:          "\u001b[36m",
	White:         "\u001b[37m",
	BrightBlack:   "\u001b[30;1m",
	BrightRed:     "\u001b[31;1m",
	BrightGreen:   "\u001b[32;1m",
	BrightYellow:  "\u001b[33;1m",
	BrightBlue:    "\u001b[34;1m",
	BrightMagenta: "\u001b[35;1m",
	BrightCyan:    "\u001b[36;1m",
	BrightWhite:   "\u001b[37;1m",
}

// 用于重置端子颜色
const AnsiColorReset = "\033[0m"

// 默认的颜色码
var DefaultTargetAnsiColor = ansiColor[Cyan]

// 将字符串中所有字符转为小写
func toLowerCase(s string) string {
	buf := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		buf[i] = c
	}
	return string(buf)
}

// 根据名称获取到对应颜色的码值
func getAnsiColor(color string) (string, bool) {
	colorLower := toLowerCase(color)

	for k, v := range ansiColor {
		colorConstLower := toLowerCase(k.String())
		if colorLower == colorConstLower {
			return v, true
		}
	}
	return "", false
}
