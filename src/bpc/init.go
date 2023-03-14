package bpc

import (
	"backup_period_checker/src/logging"

	str2duration "github.com/xhit/go-str2duration/v2"
)

type Bpc struct {
	Conf tConf
	Lg   logging.Logging
}

func Init(configFile string, lg logging.Logging) (bpc Bpc) {
	bpc = Bpc{Lg: lg}
	conf := bpc.readTomlFile(configFile)

	bpc.Conf.ResticBackupFolder = conf.ResticBackupFolder
	bpc.Conf.DefaultMaxDiff = conf.DefaultMaxDiff

	for _, el := range conf.MaxDiffs {
		dur, err := str2duration.ParseDuration(el.MaxDiffStr)
		bpc.Lg.IfErrError("can not parse duration config entry",
			logging.F{"error": err},
		)

		if err == nil {
			entry := tSmd{
				Matcher:    el.Matcher,
				MaxDiffStr: el.MaxDiffStr,
				Duration:   dur,
			}
			bpc.Conf.MaxDiffs = append(bpc.Conf.MaxDiffs, entry)
		}
	}
	return
}
