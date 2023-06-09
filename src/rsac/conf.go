package rsac

import (
	"fmt"
	"os"
	"time"

	"github.com/triole/logseal"

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

func (rsac *Rsac) readTomlFile(filename string) (conf tConf) {
	content, err := os.ReadFile(rsac.resolvePath(filename))
	rsac.Lg.IfErrFatal("can not read file", logseal.F{
		"error": err,
		"file":  filename,
	})
	err = toml.Unmarshal(content, &conf)
	rsac.Lg.IfErrFatal("unable to decode toml", logseal.F{
		"error": err,
	})

	rsac.Conf.ResticBackupFolder = conf.ResticBackupFolder

	// assemble max diff list, add tolerance to durations
	for _, el := range conf.MaxDiffs {
		if el.Str != "" {
			dur, err := rsac.str2dur(el.Str)
			if err == nil {
				newEl := el
				newEl.Dur = dur
				rsac.Conf.MaxDiffs = append(
					rsac.Conf.MaxDiffs, newEl,
				)
			}
		}
	}
	rsac.Lg.Debug("applied configuration", logseal.F{"config": fmt.Sprintf("%+v", rsac.Conf)})
	return
}
