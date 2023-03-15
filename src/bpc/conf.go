package bpc

import (
	"backup_period_checker/src/logging"
	"fmt"
	"os"
	"time"

	"github.com/pelletier/go-toml"
)

type tConf struct {
	ResticBackupFolder string `toml:"restic_backup_folder"`
	MaxDiffs           tDiffs `toml:"max_diffs"`
}

type tDiffs []tDiff

type tDiff struct {
	Matcher string `toml:"matcher"`
	Str     string `toml:"duration"`
	Dur     time.Duration
}

func (bpc *Bpc) readTomlFile(filename string) (conf tConf) {
	content, err := os.ReadFile(bpc.resolvePath(filename))
	bpc.Lg.IfErrFatal("can not read file", logging.F{
		"error": err,
		"file":  filename,
	})
	err = toml.Unmarshal(content, &conf)
	bpc.Lg.IfErrFatal("unable to decode toml", logging.F{
		"error": err,
	})

	bpc.Conf.ResticBackupFolder = conf.ResticBackupFolder

	// assemble max diff list, add tolerance to durations
	for _, el := range conf.MaxDiffs {
		if el.Str != "" {
			dur, err := bpc.str2dur(el.Str)
			if err == nil {
				newEl := el
				newEl.Dur = bpc.addDurationTolerance(dur)
				bpc.Conf.MaxDiffs = append(
					bpc.Conf.MaxDiffs, newEl,
				)
			}
		}
	}
	bpc.Lg.Debug("applied configuration", logging.F{"config": fmt.Sprintf("%+v", bpc.Conf)})
	return
}
