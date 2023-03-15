package rsac

import (
	"os"
	"path/filepath"
	"rsac/src/logging"
)

func (rsac Rsac) resolvePath(path string) (r string) {
	r = rsac.abs(path)
	if rsac.isSymLink(r) {
		lnk, err := os.Readlink(r)
		rsac.Lg.IfErrError("failed to resolve symlink", logging.F{"error": err})
		if err == nil {
			rsac.Lg.Debug(
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

func (rsac Rsac) abs(path string) (r string) {
	var err error
	r, err = filepath.Abs(path)
	rsac.Lg.IfErrError("unable to construct absolute path", logging.F{"error": err})
	return r
}

func (rsac Rsac) isSymLink(path string) bool {
	if info, err := os.Lstat(path); err == nil && info.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
}
