package bpc

import (
	"backup_period_checker/src/logging"
	"time"
)

type Bpc struct {
	Now  time.Time
	Conf tConf
	Lg   logging.Logging
}

type tFileInfo struct {
	Path        string
	LastMod     time.Time
	LastModUnix int64
	Age         time.Duration
	MaxDiff     time.Duration
}

type tFileInfos []tFileInfo

func (fi tFileInfos) Len() int {
	return len(fi)
}

func (fi tFileInfos) Less(i, j int) bool {
	return fi[i].LastModUnix > fi[j].LastModUnix
}

func (fi tFileInfos) Swap(i, j int) {
	fi[i], fi[j] = fi[j], fi[i]
}
