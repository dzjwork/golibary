package yaml

const (
	// 输入流的原始缓冲区大小
	input_raw_buffer_size = 512

	// 输入缓冲区大小
	input_buffer_size = input_raw_buffer_size * 3

	// 输出缓冲区的大小
	output_buffer_size = 128

	// 输出原始缓冲区的大小
	output_raw_buffer_size = (output_buffer_size*2 + 2)

	// 初始栈大小
	initial_stack_size = 16
	// 初始队列大小
	initial_queue_size = 16
	// 初始字符长度
	initial_string_size = 16
)

// 检查给定位置的字符是否是一个字母、字符、数字、“_”或“-”
func is_alpha(b []byte, i int) bool {
	return b[i] >= '0' && b[i] <= '9' || b[i] >= 'A' && b[i] <= 'Z' || b[i] >= 'a' && b[i] <= 'z' || b[i] == '_' || b[i] == '-'
}

// 检查给定位置的字符是否是一个数字
func is_digit(b []byte, i int) bool {
	return b[i] >= '0' && b[i] <= '9'
}

// 获取给定位置的字符，转为数字类型
func as_digit(b []byte, i int) int {
	return int(b[i]) - '0'
}

// 检查给定位置的字符是否是一个十六进制字符
func is_hex(b []byte, i int) bool {
	return b[i] >= '0' && b[i] <= '9' || b[i] >= 'A' && b[i] <= 'F' || b[i] >= 'a' && b[i] <= 'f'
}

// 获取指定位置的十六进制字符，将字符转为整数返回
func as_hex(b []byte, i int) int {
	bi := b[i]
	if bi >= 'A' && bi <= 'F' {
		return int(bi) - 'A' + 10
	}
	if bi >= 'a' && bi <= 'f' {
		return int(bi) - 'a' + 10
	}
	return int(bi) - '0'
}

// 检查指定位置字符是否是一个ASCII码
func is_ascii(b []byte, i int) bool {
	return b[i] <= 0x7F
}

// 检查开头的字符是否可以不带转义的输出
func is_printable(b []byte, i int) bool {
	return ((b[i] == 0x0A) || // . == #x0A
		(b[i] >= 0x20 && b[i] <= 0x7E) || // #x20 <= . <= #x7E
		(b[i] == 0xC2 && b[i+1] >= 0xA0) || // #0xA0 <= . <= #xD7FF
		(b[i] > 0xC2 && b[i] < 0xED) ||
		(b[i] == 0xED && b[i+1] < 0xA0) ||
		(b[i] == 0xEE) ||
		(b[i] == 0xEF && // #xE000 <= . <= #xFFFD
			!(b[i+1] == 0xBB && b[i+2] == 0xBF) && // && . != #xFEFF
			!(b[i+1] == 0xBF && (b[i+2] == 0xBE || b[i+2] == 0xBF))))
}

// 检查指定位置的字符是否是NULL
func is_z(b []byte, i int) bool {
	return b[i] == 0x00
}

// 检查换缓冲的开头是否是BOM
func is_bom(b []byte, i int) bool {
	return b[0] == 0xEF && b[1] == 0xBB && b[2] == 0xBF
}

// 检查指定位置字符是否是空格
func is_space(b []byte, i int) bool {
	return b[i] == ' '
}

// 检查指定位置的字符是否是制表符
func is_tab(b []byte, i int) bool {
	return b[i] == '\t'
}

// 检查指定位置的字符是否是空格或制表符
func is_blank(b []byte, i int) bool {
	return is_space(b, i) || is_tab(b, i)
}

// 检查指定位置的字符是否为换行符
func is_break(b []byte, i int) bool {
	return (b[i] == '\r' || // CR (#xD)
		b[i] == '\n' || // LF (#xA)
		b[i] == 0xC2 && b[i+1] == 0x85 || // NEL (#x85)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA8 || // LS (#x2028)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA9) // PS (#x2029)
}

// 检查指定位置的字符是否为换行符
func is_crlf(b []byte, i int) bool {
	return b[i] == '\r' && b[i+1] == '\n'
}

// Check if the character is a line break or NUL.
func is_breakz(b []byte, i int) bool {
	//return is_break(b, i) || is_z(b, i)
	return (
	// is_break:
	b[i] == '\r' || // CR (#xD)
		b[i] == '\n' || // LF (#xA)
		b[i] == 0xC2 && b[i+1] == 0x85 || // NEL (#x85)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA8 || // LS (#x2028)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA9 || // PS (#x2029)
		// is_z:
		b[i] == 0)
}

// Check if the character is a line break, space, or NUL.
func is_spacez(b []byte, i int) bool {
	//return is_space(b, i) || is_breakz(b, i)
	return (
	// is_space:
	b[i] == ' ' ||
		// is_breakz:
		b[i] == '\r' || // CR (#xD)
		b[i] == '\n' || // LF (#xA)
		b[i] == 0xC2 && b[i+1] == 0x85 || // NEL (#x85)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA8 || // LS (#x2028)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA9 || // PS (#x2029)
		b[i] == 0)
}

// Check if the character is a line break, space, tab, or NUL.
func is_blankz(b []byte, i int) bool {
	//return is_blank(b, i) || is_breakz(b, i)
	return (
	// is_blank:
	b[i] == ' ' || b[i] == '\t' ||
		// is_breakz:
		b[i] == '\r' || // CR (#xD)
		b[i] == '\n' || // LF (#xA)
		b[i] == 0xC2 && b[i+1] == 0x85 || // NEL (#x85)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA8 || // LS (#x2028)
		b[i] == 0xE2 && b[i+1] == 0x80 && b[i+2] == 0xA9 || // PS (#x2029)
		b[i] == 0)
}

// 查看字符的宽度
func width(b byte) int {
	if b&0x80 == 0x00 {
		return 1
	}
	if b&0xE0 == 0xC0 {
		return 2
	}
	if b&0xF0 == 0xE0 {
		return 3
	}
	if b&0xF8 == 0xF0 {
		return 4
	}
	return 0

}
