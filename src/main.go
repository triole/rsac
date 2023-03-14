package main

import (
	"backup_period_checker/src/bpc"
	"backup_period_checker/src/logging"
)

func main() {
	parseArgs()

	lg := logging.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Info("Init backup period checker", logging.F{
		"config": CLI.Config,
	})

	bpc := bpc.Init(CLI.Config, lg)
	bpc.Run()

}
