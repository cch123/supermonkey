package supermonkey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:noinline
func hey() string {
	return "hello"
}

func TestPatchByFullSymbolName(t *testing.T) {
	Convey("[TestPatchByFullSymbolName]", t, func() {
		patchGuard := PatchByFullSymbolName("github.com/cch123/supermonkey.hey", func() string {
			return "ok"
		})
		So(hey(), ShouldEqual, "ok")
		patchGuard.Unpatch()
	})
}

func TestPatch(t *testing.T) {
	Convey("[TestPatchByFullSymbolName]", t, func() {
		patchGuard := Patch((*person).do, func(_ *person, s string) string {
			return "Linda"
		})
		p := &person{}
		So(p.do("Lance"), ShouldEqual, "Linda")
		patchGuard.Unpatch()
	})
}

type person struct{ name string }

//go:noinline
func (p *person) do(s string) string {
	return s
}
