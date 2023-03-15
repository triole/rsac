package main

import (
	"rsac/src/logging"
	"rsac/src/rsac"
)

func main() {
	parseArgs()

	lg := logging.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Debug("init "+appName, logging.F{
		"config": CLI.Config,
	})

	rsac := rsac.Init(CLI.Config, lg)
	err := rsac.RunCheck()

	lg.IfErrError("check failed", logging.F{"error": err})
	if err == nil {
		lg.Debug("all snapshots up to date", nil)
	}
}
