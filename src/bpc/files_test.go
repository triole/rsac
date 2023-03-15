package bpc

import "testing"

var (
	bpc = InitForTesting()
)

func TestResolveSymlink(t *testing.T) {
	assertResolvePath(
		bpc.resolvePath("/doesnotexist"), bpc.abs("/doesnotexist"), t,
	)
	assertResolvePath(
		bpc.resolvePath("../../tmp/testdata"), bpc.abs("../../tmp/testdata"), t,
	)
	assertResolvePath(
		bpc.resolvePath("../../tmp/testdata_sl"), bpc.abs("../../tmp/testdata"), t,
	)
	assertResolvePath(
		bpc.resolvePath("../../tmp/testdata_sl_broken"), bpc.abs("../../tmp/broken_sl_target"), t,
	)
}

func assertResolvePath(res, exp string, t *testing.T) {
	if res != exp {
		t.Errorf("failed resolve symlink %q != %q", res, exp)
	}
}
