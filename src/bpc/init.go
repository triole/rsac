package bpc

import (
	"backup_period_checker/src/logging"
	"os"
	"time"
)

type Bpc struct {
	Now  time.Time
	Conf tConf
	Lg   logging.Logging
}

func Init(configFile string, lg logging.Logging) (bpc Bpc) {
	bpc = Bpc{Lg: lg}
	conf := bpc.readTomlFile(configFile)

	bpc.Now = bpc.now()
	bpc.Conf.ResticBackupFolder = os.ExpandEnv(conf.ResticBackupFolder)

	return
}
