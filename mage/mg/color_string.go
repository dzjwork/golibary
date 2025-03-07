package mg

import (
	"strconv"
)

func _() {
	var x [1]struct{}
	_ = x[Black-0]
	_ = x[Red-1]
	_ = x[Green-2]
	_ = x[Yellow-3]
	_ = x[Blue-4]
	_ = x[Magenta-5]
	_ = x[Cyan-6]
	_ = x[White-7]
	_ = x[BrightBlack-8]
	_ = x[BrightRed-9]
	_ = x[BrightGreen-10]
	_ = x[BrightYellow-11]
	_ = x[BrightBlue-12]
	_ = x[BrightMagenta-13]
	_ = x[BrightCyan-14]
	_ = x[BrightWhite-15]
}

// 所有的颜色名称
const _Color_name = "BlackRedGreenYellowBlueMagentaCyanWhiteBrightBlackBrightRedBrightGreenBrightYellowBrightBlueBrightMagentaBrightCyanBrightWhite"

// 每个颜色在颜色名称中的起始索引
var _Color_index = [...]uint8{0, 5, 8, 13, 19, 23, 30, 34, 39, 50, 59, 70, 82, 92, 105, 115, 126}

// 根据颜色索引获取到对应的名称
func (i Color) String() string {
	if i < 0 || i >= Color(len(_Color_index)-1) {
		return "Color(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}
