package rsac

import "testing"

var (
	rsac = InitForTesting()
)

func TestResolveSymlink(t *testing.T) {
	assertResolvePath(
		rsac.resolvePath("/doesnotexist"), rsac.abs("/doesnotexist"), t,
	)
	assertResolvePath(
		rsac.resolvePath("../../tmp/testdata"), rsac.abs("../../tmp/testdata"), t,
	)
	assertResolvePath(
		rsac.resolvePath("../../tmp/testdata_sl"), rsac.abs("../../tmp/testdata"), t,
	)
	assertResolvePath(
		rsac.resolvePath("../../tmp/testdata_sl_broken"), rsac.abs("../../tmp/broken_sl_target"), t,
	)
}

func assertResolvePath(res, exp string, t *testing.T) {
	if res != exp {
		t.Errorf("failed resolve symlink %q != %q", res, exp)
	}
}
