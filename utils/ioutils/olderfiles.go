package ioutils

import (
	"os"
	"time"
)

func isOlderThan(t time.Time, duration time.Duration) bool {
	return time.Since(t) > duration
}

func FindFilesOlderThan(dir string, duration time.Duration) (files []os.DirEntry, err error) {
	tmpfiles, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {

		fileinfo, err := file.Info()
		if err == nil {

			if fileinfo.Mode().IsRegular() && !fileinfo.IsDir() {
				if isOlderThan(fileinfo.ModTime(), duration) {
					files = append(files, file)
				}
			}
		}
	}
	return
}
