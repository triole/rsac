package rsac

import (
	"os"

	"github.com/triole/logseal"
)

func Init(configFile string, lg logseal.Logseal) (rsac Rsac) {
	rsac = Rsac{Lg: lg}
	conf := rsac.readTomlFile(configFile)

	rsac.Now = rsac.now()
	rsac.Conf.ResticBackupFolder = os.ExpandEnv(conf.ResticBackupFolder)

	return
}

func InitForTesting() (rsac Rsac) {
	lg := logseal.Init("debug", nil, true, false)
	rsac.Lg = lg
	return
}
