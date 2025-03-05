package mg

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

// 编译时二进制文件存储的位置
const CacheEnv = "MAGEFILE_CACHE"

// 编译时输出详细信息
const VerboseEnv = "MAGEFILE_VERBOSE"

// debug模式进行编译
const DebugEnv = "MAGEFILE_DEBUG"

// 指定go的二进制文件，希望用于magefile编译
const GoCmdEnv = "MAGEFILE_GOCMD"

// 忽略magefile中指定的默认目标
const IgnoreDefaultEnv = "MAGEFILE_IGNOREDEFAULT"

// 快速散列磁盘中的文件，用于确定磁盘中二进制文件是否要重建
const HashFastEnv = "MAGEFILE_HASHFAST"

// 是否支持彩色输出
const EnableColorEnv = "MAGEFILE_ENABLE_COLOR"

// 着色使用的颜色
const TargetColorEnv = "MAGEFILE_TARGET_COLOR"

// 查看编译时是否要输出详细信息
func Verbose() bool {
	b, _ := strconv.ParseBool(os.Getenv(VerboseEnv))
	return b
}

// 查看是否以Debug模式进行编译
func Debug() bool {
	b, _ := strconv.ParseBool(os.Getenv(DebugEnv))
	return b
}

// 获取到go的可执行文件
func GoCmd() string {
	if cmd := os.Getenv(GoCmdEnv); cmd != "" {
		return cmd
	}
	return "go"
}

// 查看是否快速散列磁盘中的文件，用于确定磁盘中二进制文件是否要重建
func HashFast() bool {
	b, _ := strconv.ParseBool(os.Getenv(HashFastEnv))
	return b
}

// 查看是否忽略magefile中指定的默认目标
func IgnoreDefault() bool {
	b, _ := strconv.ParseBool(os.Getenv(IgnoreDefaultEnv))
	return b
}

// 查看是否支持彩色输出
func EnableColor() bool {
	b, _ := strconv.ParseBool(os.Getenv(EnableColorEnv))
	return b
}

// 查看着色使用的颜色
func TargetColor() string {
	s, exists := os.LookupEnv(TargetColorEnv)
	if exists {
		if c, ok := getAnsiColor(s); ok {
			return c
		}
	}
	return DefaultTargetAnsiColor
}

// 查看编译好的二进制文件存储位置
func CacheDir() string {
	d := os.Getenv(CacheEnv)
	if d != "" {
		return d
	}
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"), "magefile")
	default:
		return filepath.Join(os.Getenv("HOME"), ".magefile")
	}
}

// 允许对类似的命令进行分组
type Namespace struct{}
