package bpc

import (
	"regexp"
	"time"
)

func (bpc Bpc) now() time.Time {
	return time.Now()
}

func rxMatch(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}
