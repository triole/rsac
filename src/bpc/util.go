package bpc

import (
	"backup_period_checker/src/logging"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func (bpc Bpc) findFolder(root, matcher string) {

}

func (bpc Bpc) findFile(root, matcher string) {

}

func (bpc Bpc) find(root, matcher string, isDir bool) []string {
	inf, err := os.Stat(root)
	bpc.Lg.IfErrFatal("root is not a folder", logging.F{"error": err})

	if !inf.IsDir() {
		err = errors.New("not a folder")
		bpc.Lg.IfErrFatal("root is not a folder", logging.F{"error": err})
	}

	filelist := []string{}
	rxf, _ := regexp.Compile(matcher)

	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if rxf.MatchString(path) {
			inf, err := os.Stat(path)
			if err == nil {
				if inf.IsDir() == isDir {
					filelist = append(filelist, path)
				}
			} else {
				bpc.Lg.IfErrError("stat of file failed", logging.F{"path": path})
			}
		}
		return nil
	})
	bpc.Lg.IfErrFatal("file path walk failed", logging.F{
		"error": err,
		"path":  root,
	})
	return filelist
}

func (bpc Bpc) getLastModified(filename string) (r time.Time) {
	if filename != "" {
		stats, err := os.Stat(filename)
		if err != nil {
			fmt.Printf("Error os stat %q\n", filename)
		} else {
			r = stats.ModTime()
		}
	}
	return
}
