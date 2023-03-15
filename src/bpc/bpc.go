package bpc

import (
	"backup_period_checker/src/logging"
	"errors"
	"strconv"
	"time"
)

func (bpc Bpc) RunCheck() (err error) {
	latestSnapshots := bpc.findLatestSnapshots()
	errCounter := 0
	for _, el := range latestSnapshots {
		if el.Age <= el.MaxDiff {
			bpc.Lg.Info(bpc.makeSnapInfo("snapshot up to date", el))
		} else {
			bpc.Lg.Warn(bpc.makeSnapInfo("snapshot outdated", el))
			errCounter += 1
			err = errors.New(strconv.Itoa(errCounter) + " snapshots exceed their expected maximum age")
		}
	}
	return
}

func (bpc Bpc) makeSnapInfo(msg string, fi tFileInfo) (string, logging.F) {
	return msg, logging.F{
		"age": bpc.roundDuration(fi.Age), "max_diff": fi.MaxDiff, "path": fi.Path,
	}
}

func (bpc Bpc) roundDuration(dur time.Duration) time.Duration {
	return dur.Round(time.Second)
}
