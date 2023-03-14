package bpc

import (
	"fmt"
	"sort"
)

func (bpc Bpc) Run() {
	latestSnapshots := bpc.findLatestSnapshots()

	for _, el := range latestSnapshots {
		fmt.Printf("%+v\n", el)
	}
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
