package bpc

import (
	"backup_period_checker/src/logging"
	"os"
	"path/filepath"
)

func (bpc Bpc) resolvePath(path string) (r string) {
	r = bpc.abs(path)
	if bpc.isSymLink(r) {
		lnk, err := os.Readlink(r)
		bpc.Lg.IfErrError("failed to resolve symlink", logging.F{"error": err})
		if err == nil {
			bpc.Lg.Debug(
				"resolve symlink",
				logging.F{
					"source": r,
					"target": lnk,
				},
			)
			r = lnk
		}
	}
	return
}

func (bpc Bpc) abs(path string) (r string) {
	var err error
	r, err = filepath.Abs(path)
	bpc.Lg.IfErrError("unable to construct absolute path", logging.F{"error": err})
	return r
}

func (bpc Bpc) isSymLink(path string) bool {
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}
