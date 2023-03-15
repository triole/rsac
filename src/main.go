package main

import (
	"backup_period_checker/src/bpc"
	"backup_period_checker/src/logging"
)

func main() {
	parseArgs()

	lg := logging.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Debug("init "+appName, logging.F{
		"config": CLI.Config,
	})

	bpc := bpc.Init(CLI.Config, lg)
	err := bpc.RunCheck()

	lg.IfErrError("check failed", logging.F{"error": err})
	if err == nil {
		lg.Debug("all snapshots up to date", nil)
	}
}
