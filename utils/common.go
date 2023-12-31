package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"time"
)

func GetRootDir() string {
	CacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	AppDir := path.Join(CacheDir, "shak-daemon")
	_, err = os.Stat(AppDir)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(AppDir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	return AppDir
}

func GetFolderDumpDir(BundleName string) string {
	dir := path.Join(RootDir, BundleName, "folders")
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	return dir
}

func GetFileDumpDir(BundleName string) string {
	dir := path.Join(RootDir, BundleName, "files")
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	return dir
}

func GetCommandDumpDir(BundleName string, CommandName string) string {
	dir := path.Join(RootDir, BundleName, "commands")
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	return path.Join(dir, CommandName+".output")
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func GetBundleName() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Bundle-%s-%d", hostname, time.Now().Unix())
}

func GetBundleDir(BundleName string) string {
	return path.Join(RootDir, BundleName)
}

func GetBundleArchivePath(BundleName string) string {
	return path.Join(RootDir, BundleName+".tar.gzip")
}
