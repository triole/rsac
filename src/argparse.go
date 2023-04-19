package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "restic snapshot age checker"
	appDescription = "checks if latest restic snapshots are up to date"
	appMainversion = "0.1"
)

var CLI struct {
	Config      string `help:"config file path" arg:"" optional:""`
	LogFile     string `help:"log file" short:"l" default:"/dev/stdout"`
	LogLevel    string `help:"log level" short:"e" default:"info" enum:"debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text" short:"n"`
	LogJSON     bool   `help:"enable json log, instead of text one" short:"j"`
	VersionFlag bool   `help:"display version" short:"V"`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
		kong.Vars{},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	if CLI.Config == "" {
		fmt.Printf("\n%s\n", "[error] positional arg required, please pass config file path")
		os.Exit(1)
	}
	// ctx.FatalIfErrorf(err)
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}

// func getBindir() (s string) {
// 	ex, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	s = filepath.Dir(ex)
// 	return
// }
