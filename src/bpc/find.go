package bpc

import (
	"backup_period_checker/src/logging"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func (bpc Bpc) findLatestSnapshots() (latestSnapshots tFileInfos) {
	snapshotFolders := bpc.findFolders(
		bpc.Conf.ResticBackupFolder, "/snapshots$",
	)

	for _, val := range snapshotFolders {
		filesInFolder := bpc.findFiles(
			val.Path, ".*",
		)
		sort.Sort(filesInFolder)
		if len(filesInFolder) > 0 {
			latestSnapshots = append(latestSnapshots, filesInFolder[0])
		}
	}
	return
}

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
			maxDiff, matcher := bpc.getMaxDiffEntry(pth)
			newFileInfo = tFileInfo{
				Path:        pth,
				LastMod:     inf.ModTime(),
				LastModUnix: inf.ModTime().Unix(),
				Age:         bpc.Now.Sub(inf.ModTime()),
				MaxDiff:     maxDiff,
				Matcher:     matcher,
			}
		}
		if inf.IsDir() == isDir {
			fis = append(fis, newFileInfo)
		}
		bpc.Lg.IfErrError("root is not a folder", logging.F{"error": err})
	}
	return
}

func (bpc Bpc) getMaxDiffEntry(path string) (dur time.Duration, matcher string) {
	for _, el := range bpc.Conf.MaxDiffs {
		if rxMatch(el.Matcher, path) {
			return el.Dur, el.Matcher
		}
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
	root = bpc.resolvePath(root)
	var err error = filepath.Walk(root, bpc.visit(&files, rx))
	bpc.Lg.IfErrFatal(
		"unable to detect files",
		logging.F{"error": err},
	)
	return
}
