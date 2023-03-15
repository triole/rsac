package bpc

import (
	"backup_period_checker/src/logging"
	"regexp"
	"time"

	"github.com/xhit/go-str2duration/v2"
)

func (bpc Bpc) now() time.Time {
	return time.Now()
}

func rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func (bpc Bpc) str2dur(s string) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	bpc.Lg.IfErrError("can not parse string to duration",
		logging.F{"error": err},
	)
	return
}

func (bpc Bpc) addDurationTolerance(dur time.Duration) time.Duration {
	return dur + time.Duration(time.Minute*30)
}
