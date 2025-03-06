package target

import "os"

// 检查路径中的文件是否有更新
func Path(dst string, sources ...string) (bool, error) {
	// 1、 解析环境变量并尝试获取文件的信息
	stat, err := os.Stat(os.ExpandEnv(dst))

	// 2、 如果文件不存在
	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}
	return PathNewer(stat.ModTime(), sources...)
}

// 检查匹配到的文件路径下的文件是否更新了
func Glob(dst string, globs ...string) (bool, error) {
	stat, err := os.Stat(os.ExpandEnv(dst))

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}
	return GlobNewer(stat.ModTime(), globs...)
}

// 查找dst中是否有文件比目标时间新
func Dir(dst string, sources ...string) (bool, error) {
	dst = os.ExpandEnv(dst)
	stat, err := os.Stat(dst)

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	destTime := stat.ModTime()

	if stat.IsDir() {
		destTime, err = NewestModTime(dst)

		if err != nil {
			return false, err
		}
	}
	return DirNewer(destTime, sources...)
}
