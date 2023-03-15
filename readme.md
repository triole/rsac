# Restic Snapshot Age Checker ![build](https://github.com/triole/rsac/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

The `Restic Snapshot Age Checker` checks a [restic](https://github.com/restic/restic) backup folder for the age of the latest created snapshots. All checks are run on file level which does not require any access data for certain repositories. It scans a folder for `snapshots` sub folders and looks up the latest entries inside these folders. Afterwards `rsac` checks the age of the latest snapshots. Rules can be configured in a toml file which looks the one below.

Max diff rules are applied in order. The first that fits is used. Make sure your most common rule is at the end of the list.

If any snapshot is outdated the program exits with status code 1. For example a use case would be to periodically run [goss](https://github.com/goss-org/goss) tests to detect repositories with outdated latest snapshots. By doing this one can find out which backups might have gone wrong.

```go mdox-exec="cat examples/conf.toml"
restic_backup_folder = "${HOME}/rolling/golang/projects/rsac/tmp/testdata"

[[max_diffs]]
matcher = ".*/user01/.*"
duration = "6h"

[[max_diffs]]
matcher = ".*/user02/repo02.*"
duration = "18h30m"

[[max_diffs]]
matcher = ".*/user03/repo3.*"
duration = "1w"

[[max_diffs]]
matcher = ".*"
duration = "1d"
```

## Help

```go mdox-exec="r -h"

checks if latest restic snapshots are up to date

Arguments:
  [<config>]    config file path

Flags:
  -h, --help                      Show context-sensitive help.
  -l, --log-file="/dev/stdout"    log file
  -e, --log-level="info"          log level
  -n, --log-no-colors             disable output colours, print plain text
  -j, --log-json                  enable json log, instead of text one
  -V, --version-flag              display version
```
