package rsac

import (
	"errors"
	"rsac/src/logging"
	"strconv"
	"strings"
	"time"
)

func (rsac Rsac) RunCheck() (err error) {
	latestSnapshots := rsac.findLatestSnapshots()
	errCounter := 0
	for _, el := range latestSnapshots {
		if el.MaxDiff > 0 {
			if el.Age <= el.MaxDiff {
				rsac.Lg.Info(rsac.makeSnapInfo("up to date", el))
			} else {
				rsac.Lg.Warn(rsac.makeSnapInfo("outdated", el))
				errCounter += 1
				err = errors.New(strconv.Itoa(errCounter) + " snapshots exceed their expected maximum age")
			}
		} else {
			rsac.Lg.Debug(rsac.makeSnapInfo("skip, no matcher did apply", el))
		}
	}
	return
}

func (rsac Rsac) makeSnapInfo(msg string, fi tFileInfo) (string, logging.F) {
	fields := logging.F{
		"age":  rsac.roundDuration(fi.Age),
		"path": fi.Path,
	}
	if !strings.HasPrefix(msg, "skip") {
		fields["max_diff"] = fi.MaxDiff
		fields["matcher"] = fi.Matcher
	}
	return msg, fields
}

func (rsac Rsac) roundDuration(dur time.Duration) time.Duration {
	return dur.Round(time.Second)
}