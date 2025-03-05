package target

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	errNewer = fmt.Errorf("newer item encountered")
)

// 查找是否有文件比目标时间新
func DirNewer(target time.Time, sources ...string) (bool, error) {
	// 该函数用于检查文件的最后修改时间是否在指定时间之后
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.ModTime().After(target) {
			return errNewer
		}
		return nil
	}

	for _, source := range sources {
		source = os.ExpandEnv(source)
		err := filepath.Walk(source, walkFn)

		if err == nil {
			continue
		}

		if err == errNewer {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// 检查源文件中是否有文件更新了
func PathNewer(target time.Time, sources ...string) (bool, error) {
	for _, source := range sources {
		// 1、 解析路径中的环境变量值
		source = os.ExpandEnv(source)
		// 2、 获取文件或目录的详细信息
		stat, err := os.Stat(source)
		// 3、 文件不存在退出
		if err != nil {
			return false, err
		}
		// 4、 查看文件是否在指定时间之后，如果是说明文件更新了
		if stat.ModTime().After(target) {
			return true, nil
		}
	}
	// 5、 遍历完所有都没退出说明所有文件都没更新
	return false, nil
}

// 检查匹配到的文件路径下的文件是否更新了
func GlobNewer(target time.Time, sources ...string) (bool, error) {
	for _, g := range sources {
		// 1、 获取到所有匹配的文件路径
		files, err := filepath.Glob(g)

		if err != nil {
			return false, err
		}

		if len(files) == 0 {
			return false, fmt.Errorf("glob didn't match any files: %s", g)
		}

		// 2、 检查路径下的文件是否更新了
		newer, err := PathNewer(target, files...)
		if err != nil {
			return false, err
		}
		if newer {
			return true, nil
		}
	}
	return false, nil
}

// 遍历指定目录中的所有文件，找到最旧的一个文件最后修改的时间
func OldesModTime(targets ...string) (time.Time, error) {
	t := time.Now().Add(time.Hour * 100000)
	for _, target := range targets {
		walkFn := func(_ string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			mtime := info.ModTime()
			if mtime.Before(t) {
				t = mtime
			}
			return nil
		}
		// 遍历根目录下的文件树，每个文件都执行walkFn函数
		if err := filepath.Walk(target, walkFn); err != nil {
			return t, err
		}
	}
	return t, nil
}

// 遍历指定目录中的所有文件，找到最旧的一个文件最后修改的时间
func NewestModTime(targets ...string) (time.Time, error) {
	t := time.Time{}

	for _, target := range targets {
		walkFn := func(_ string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			mTime := info.ModTime()
			if mTime.After(t) {
				t = mTime
			}
			return nil
		}

		if err := filepath.Walk(target, walkFn); err != nil {
			return t, err
		}
	}
	return t, nil
}
