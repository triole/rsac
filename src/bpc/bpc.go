package bpc

import (
	"backup_period_checker/src/logging"
	"errors"
	"sort"
	"strconv"
)

func (bpc Bpc) RunCheck() (err error) {
	latestSnapshots := bpc.findLatestSnapshots()
	errCounter := 0
	for _, el := range latestSnapshots {
		if el.Age <= el.MaxDiff {
			bpc.Lg.Info(
				"snapshot up to date",
				logging.F{
					"age": el.Age, "max_diff": el.MaxDiff, "path": el.Path,
				},
			)
		} else {
			bpc.Lg.Warn(
				"latest snapshoot outdated",
				logging.F{
					"age": el.Age, "max_diff": el.MaxDiff, "path": el.Path,
				},
			)
			errCounter += 1
			err = errors.New(strconv.Itoa(errCounter) + " snapshots exceed their expected maximum age")
		}
	}
	return
}

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
