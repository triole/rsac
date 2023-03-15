package rsac

import (
	"errors"
	"rsac/src/logging"
	"strconv"
	"time"
)

func (rsac Rsac) RunCheck() (err error) {
	latestSnapshots := rsac.findLatestSnapshots()
	errCounter := 0
	for _, el := range latestSnapshots {
		if el.Age <= el.MaxDiff {
			rsac.Lg.Info(rsac.makeSnapInfo("snapshot up to date", el))
		} else {
			rsac.Lg.Warn(rsac.makeSnapInfo("snapshot outdated", el))
			errCounter += 1
			err = errors.New(strconv.Itoa(errCounter) + " snapshots exceed their expected maximum age")
		}
	}
	return
}

func (rsac Rsac) makeSnapInfo(msg string, fi tFileInfo) (string, logging.F) {
	return msg, logging.F{
		"age":      rsac.roundDuration(fi.Age),
		"max_diff": fi.MaxDiff,
		"matcher":  fi.Matcher,
		"path":     fi.Path,
	}
}

func (rsac Rsac) roundDuration(dur time.Duration) time.Duration {
	return dur.Round(time.Second)
}
