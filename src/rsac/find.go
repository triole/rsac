package rsac

import (
	"log"
	"os"
	"path/filepath"
	"rsac/src/logging"
	"sort"
	"time"
)

func (rsac Rsac) findLatestSnapshots() (latestSnapshots tFileInfos) {
	snapshotFolders := rsac.findFolders(
		rsac.Conf.ResticBackupFolder, "/snapshots$",
	)

	for _, val := range snapshotFolders {
		filesInFolder := rsac.findFiles(
			val.Path, ".*",
		)
		sort.Sort(filesInFolder)
		if len(filesInFolder) > 0 {
			latestSnapshots = append(latestSnapshots, filesInFolder[0])
		}
	}
	return
}

func (rsac Rsac) findFiles(root, rx string) (fis tFileInfos) {
	fileList := rsac.find(root, rx)
	fis = rsac.filterFileList(fileList, false)
	return
}

func (rsac Rsac) findFolders(root, rx string) (fis tFileInfos) {
	fileList := rsac.find(root, rx)
	fis = rsac.filterFileList(fileList, true)
	return
}

func (rsac Rsac) filterFileList(fileList []string, isDir bool) (fis tFileInfos) {
	for _, pth := range fileList {
		inf, err := os.Stat(pth)
		rsac.Lg.IfErrError("failed to stat file", logging.F{
			"error": err,
			"path":  pth,
		})
		newFileInfo := tFileInfo{Path: pth}
		if !inf.IsDir() && !isDir {
			maxDiff, matcher := rsac.getMaxDiffEntry(pth)
			newFileInfo = tFileInfo{
				Path:        pth,
				LastMod:     inf.ModTime(),
				LastModUnix: inf.ModTime().Unix(),
				Age:         rsac.Now.Sub(inf.ModTime()),
				MaxDiff:     maxDiff,
				Matcher:     matcher,
			}
		}
		if inf.IsDir() == isDir {
			fis = append(fis, newFileInfo)
		}
		rsac.Lg.IfErrError("root is not a folder", logging.F{"error": err})
	}
	return
}

func (rsac Rsac) getMaxDiffEntry(path string) (dur time.Duration, matcher string) {
	for _, el := range rsac.Conf.MaxDiffs {
		if rxMatch(el.Matcher, path) {
			return el.Dur, el.Matcher
		}
	}
	return
}

func (rsac Rsac) visit(files *[]string, rx string) filepath.WalkFunc {
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

func (rsac Rsac) find(root string, rx string) (files []string) {
	root = rsac.resolvePath(root)
	var err error = filepath.Walk(root, rsac.visit(&files, rx))
	rsac.Lg.IfErrFatal(
		"unable to detect files",
		logging.F{"error": err},
	)
	return
}
