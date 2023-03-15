package bpc

import (
	"backup_period_checker/src/logging"
	"os"
)

func Init(configFile string, lg logging.Logging) (bpc Bpc) {
	bpc = Bpc{Lg: lg}
	conf := bpc.readTomlFile(configFile)

	bpc.Now = bpc.now()
	bpc.Conf.ResticBackupFolder = os.ExpandEnv(conf.ResticBackupFolder)

	return
}

func InitForTesting() (bpc Bpc) {
	lg := logging.Init("debug", "/dev/stdout", true, false)
	bpc.Lg = lg
	return
}
