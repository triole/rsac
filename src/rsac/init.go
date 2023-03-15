package rsac

import (
	"os"
	"rsac/src/logging"
)

func Init(configFile string, lg logging.Logging) (rsac Rsac) {
	rsac = Rsac{Lg: lg}
	conf := rsac.readTomlFile(configFile)

	rsac.Now = rsac.now()
	rsac.Conf.ResticBackupFolder = os.ExpandEnv(conf.ResticBackupFolder)

	return
}

func InitForTesting() (rsac Rsac) {
	lg := logging.Init("debug", "/dev/stdout", true, false)
	rsac.Lg = lg
	return
}
