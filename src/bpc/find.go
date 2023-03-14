package bpc

import (
	"backup_period_checker/src/logging"
	"log"
	"os"
	"path/filepath"
)

func (bpc Bpc) findFiles(root, rx string) (fis tFileInfos) {
	fileList := bpc.find(root, rx)
	fis = bpc.filterFileList(fileList, false)
	return
}

func (bpc Bpc) findFolders(root, rx string) (fis tFileInfos) {
	fileList := bpc.find(root, rx)
	fis = bpc.filterFileList(fileList, true)
	return
}

func (bpc Bpc) filterFileList(fileList []string, isDir bool) (fis tFileInfos) {
	for _, pth := range fileList {
		inf, err := os.Stat(pth)
		bpc.Lg.IfErrError("failed to stat file", logging.F{
			"error": err,
			"path":  pth,
		})
		newFileInfo := tFileInfo{Path: pth}
		if !inf.IsDir() && !isDir {
			newFileInfo = tFileInfo{
				Path:        pth,
				LastMod:     inf.ModTime(),
				LastModUnix: inf.ModTime().Unix(),
				Age:         bpc.Now.Sub(inf.ModTime()),
			}
		}
		if inf.IsDir() == isDir {
			fis = append(fis, newFileInfo)
		}
		bpc.Lg.IfErrError("root is not a folder", logging.F{"error": err})
	}
	return
}

func (bpc Bpc) visit(files *[]string, rx string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if rxMatch(rx, path) {
			*files = append(*files, path)
		}
		return nil
	}
}

func (bpc Bpc) find(root string, rx string) (files []string) {
	lnk, _ := os.Readlink(root)
	if lnk != "" {
		root = lnk
	}
	err := filepath.Walk(root, bpc.visit(&files, rx))
	if err != nil {
		panic(err)
	}
	return
}
