package fileutil

import (
	"fmt"
	"os"
)

// CheckDir 检查路径是否存在，如果不存在则创建
func CheckDir(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 路径不存在，创建目录
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("创建目录失败: %w", err)
		}
	} else if err != nil {
		// 其他错误（比如权限问题）
		return fmt.Errorf("检查路径出错: %w", err)
	} else if !info.IsDir() {
		// 路径存在但不是目录
		return fmt.Errorf("路径 %s 已存在但不是目录", path)
	}
	return nil
}
