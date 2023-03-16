package rsac

import (
	"regexp"
	"time"

	"github.com/triole/logseal"
	"github.com/xhit/go-str2duration/v2"
)

func (rsac Rsac) now() time.Time {
	return time.Now()
}

func rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

func (rsac Rsac) str2dur(s string) (dur time.Duration, err error) {
	dur, err = str2duration.ParseDuration(s)
	rsac.Lg.IfErrError("can not parse string to duration",
		logseal.F{"error": err},
	)
	return
}
