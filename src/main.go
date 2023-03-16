package main

import (
	"fut_clicker/logging"
	"rsac/src/rsac"

	"github.com/triole/logseal"
)

func main() {
	parseArgs()

	lg := logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Debug("init "+appName, logging.F{
		"config": CLI.Config,
	})

	rsac := rsac.Init(CLI.Config, lg)
	err := rsac.RunCheck()

	lg.IfErrError("check failed", logging.F{"error": err})
	if err == nil {
		lg.Debug("all snapshots are up to date", nil)
	}
}
