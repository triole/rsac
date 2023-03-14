package bpc

import (
	"backup_period_checker/src/logging"
	"os"
	"time"

	"github.com/pelletier/go-toml"
)

type tConf struct {
	ResticBackupFolder string `toml:"restic_backup_folder"`
	DefaultMaxDiff     string `toml:"default_max_diff"`
	MaxDiffs           []tSmd `toml:"smds"`
}

type tSmd struct {
	Matcher    string `toml:"matcher"`
	MaxDiffStr string `toml:"max_diff"`
	Duration   time.Duration
}

func (bpc Bpc) readTomlFile(filename string) (conf tConf) {
	content, err := os.ReadFile(filename)
	bpc.Lg.IfErrFatal("Can not read file", logging.F{
		"error": err,
		"file":  filename,
	})
	err = toml.Unmarshal(content, &conf)
	bpc.Lg.IfErrFatal("Unable to decode toml", logging.F{
		"error": err,
	})
	return
}
