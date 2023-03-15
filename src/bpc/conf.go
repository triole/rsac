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

type tDiffs struct {
	Default  tDiff `toml:"default_duration"`
	Specific []tDiff
}

type tDiff struct {
	Matcher string `toml:"matcher"`
	Str     string `toml:"duration"`
	Dur     time.Duration
}

func (bpc *Bpc) readTomlFile(filename string) (conf tConf) {
	content, err := os.ReadFile(filename)
	bpc.Lg.IfErrFatal("can not read file", logging.F{
		"error": err,
		"file":  filename,
	})
	err = toml.Unmarshal(content, &conf)
	bpc.Lg.IfErrFatal("unable to decode toml", logging.F{
		"error": err,
	})

	bpc.Conf.ResticBackupFolder = conf.ResticBackupFolder

	// parse default duration string
	defaultDur, err := bpc.str2dur(conf.MaxDiffs.Default.Str)
	bpc.Lg.IfErrFatal(
		"no default duration specified",
		logging.F{"config": filename, "error": err},
	)
	if err == nil {
		bpc.Conf.MaxDiffs.Default.Dur = bpc.addDurationTolerance(defaultDur)
	}

	// assemble specific duration list, add 30m tolerance to durations
	for _, el := range conf.MaxDiffs.Specific {
		if el.Str != "" {
			dur, err := bpc.str2dur(el.Str)
			if err == nil {
				newEl := el
				newEl.Dur = bpc.addDurationTolerance(dur)
				bpc.Conf.MaxDiffs.Specific = append(
					bpc.Conf.MaxDiffs.Specific, newEl,
				)
			}
		}
	}
	bpc.Lg.Info("apply configuration", logging.F{"config": fmt.Sprintf("%+v", bpc.Conf)})
	return
}
