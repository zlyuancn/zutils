package zutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var TemplateDir = templateDirCli{}

type templateDirCli struct{}

type TemplateDirReleaseDirArgs struct {
	TemplateDirName  string                                  // 模板文件夹路径
	DestDirPath      string                                  // 释放到磁盘的路径
	CleanDestDirPath bool                                    // 清理目标文件夹, 如果设为true, 会将目标文件夹清空
	MkdirFn          func(absPath string) error              // 创建文件夹函数, 可以为空
	WriteFileFn      func(absPath string, data []byte) error // 创建文件函数, 可以为空
}

// 释放模板文件夹到指定路径
func (f templateDirCli) ReleaseDir(args TemplateDirReleaseDirArgs) error {
	// 如果未提供创建文件夹函数，则使用默认的mkdir函数
	if args.MkdirFn == nil {
		args.MkdirFn = f.mkdir
	}
	// 如果未提供创建文件函数，则使用默认的writeFileFn函数
	if args.WriteFileFn == nil {
		args.WriteFileFn = f.writeFileFn
	}

	// 将目标目录路径转化为绝对路径
	destDirPath := filepath.Clean(args.DestDirPath)
	if !filepath.IsAbs(destDirPath) {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("获取当前工作目录失败: %v", err)
		}
		destDirPath = filepath.Join(cwd, destDirPath)
	}
	args.DestDirPath = destDirPath

	// 如果需要清空目标目录，则先清空
	if args.CleanDestDirPath {
		err := os.RemoveAll(args.DestDirPath)
		if err != nil {
			return fmt.Errorf("清理目标目录失败: %v", err)
		}
	}

	// 创建目标目录
	if err := args.MkdirFn(args.DestDirPath); err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	return f.dispatchDirs(args, args.TemplateDirName)
}

func (f templateDirCli) dispatchDirs(args TemplateDirReleaseDirArgs, templateDirPath string) error {
	dirs, err := os.ReadDir(templateDirPath)
	if err != nil {
		return fmt.Errorf("读取目录资源失败: %v", err)
	}

	for _, dir := range dirs {
		path := templateDirPath + "/" + dir.Name()
		if dir.IsDir() {
			err = f.releaseDir(args, path)
		} else {
			err = f.releaseFile(args, path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (f templateDirCli) releaseDir(args TemplateDirReleaseDirArgs, templateDirPath string) error {
	path := strings.TrimPrefix(templateDirPath, args.TemplateDirName) // 相对路径, 要去掉模板文件夹基础路径
	path = strings.TrimPrefix(path, "/")                              // 要去掉前导斜杠
	destPath, err := f.dirJoin(args.DestDirPath, path)
	if err != nil {
		return err
	}

	err = args.MkdirFn(destPath)
	if err != nil {
		return err
	}

	return f.dispatchDirs(args, templateDirPath)
}

func (f templateDirCli) releaseFile(args TemplateDirReleaseDirArgs, templateFilePath string) error {
	data, err := os.ReadFile(templateFilePath)
	if err != nil {
		return fmt.Errorf("读取文件资源失败: %v", err)
	}

	path := strings.TrimPrefix(templateFilePath, args.TemplateDirName) // 相对路径, 要去掉模板文件夹基础路径
	path = strings.TrimPrefix(path, "/")                               // 要去掉前导斜杠
	destPath, err := f.dirJoin(args.DestDirPath, path)
	if err != nil {
		return err
	}

	return args.WriteFileFn(destPath, data)
}

func (templateDirCli) mkdir(path string) error {
	return os.MkdirAll(path, 0755)
}

func (templateDirCli) writeFileFn(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func (templateDirCli) dirJoin(path1, path2 string) (string, error) {
	// 处理路径分隔符
	path1 = strings.ReplaceAll(path1, "\\", "/")
	path2 = strings.ReplaceAll(path2, "\\", "/")
	if filepath.IsAbs(path2) {
		// 如果是绝对路径，则使用path2作为路径
		return filepath.Clean(path1 + "/" + path2), nil
	}

	// 否则将path2添加到path1上，并使用filepath.Join连接路径
	path := filepath.Join(path1, path2)
	if !strings.HasPrefix(strings.ReplaceAll(path, "\\", "/"), path1) {
		return "", fmt.Errorf("路径非法: %s", path)
	}
	return filepath.Clean(path), nil
}
