package main

import (
	"os"
	"rsac/src/rsac"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()

	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Debug("init "+appName, logseal.F{
		"config": CLI.Config,
	})

	rsac := rsac.Init(CLI.Config, lg)
	err, errCounter := rsac.RunCheck()

	if errCounter == 0 && err == nil {
		lg.Debug("all snapshots are up to date", nil)
	} else {
		lg.IfErrError("check failed, snapshots outdated",
			logseal.F{"error": err},
		)
		os.Exit(1)
	}
}
