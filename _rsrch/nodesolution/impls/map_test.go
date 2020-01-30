package impls

import (
	"strconv"
	"testing"

	wish "github.com/warpfork/go-wish"

	ipld "github.com/ipld/go-ipld-prime/_rsrch/nodesolution"
	"github.com/ipld/go-ipld-prime/must"
)

var tableStrInt = [25]struct {
	s string
	i int
}{}

func init() {
	for i := 1; i <= 25; i++ {
		tableStrInt[i-1] = struct {
			s string
			i int
		}{"k" + strconv.Itoa(i), i}
	}
}

func TestGennedMapIntValues(t *testing.T) {
	CheckMaps(t, Type__Map_K_T{})
}
func TestGenericMapIntValues(t *testing.T) {
	CheckMaps(t, Style__Map{})
}

func buildMapStrIntN3(ns ipld.NodeStyle) ipld.Node {
	nb := ns.NewBuilder()
	ma, err := nb.BeginMap(3)
	must.NotError(err)
	must.NotError(ma.AssembleKey().AssignString("whee"))
	must.NotError(ma.AssembleValue().AssignInt(1))
	must.NotError(ma.AssembleKey().AssignString("woot"))
	must.NotError(ma.AssembleValue().AssignInt(2))
	must.NotError(ma.AssembleKey().AssignString("waga"))
	must.NotError(ma.AssembleValue().AssignInt(3))
	must.NotError(ma.Done())
	n, err := nb.Build()
	must.NotError(err)
	return n
}

func buildMapStrIntN25(ns ipld.NodeStyle) ipld.Node {
	nb := ns.NewBuilder()
	ma, err := nb.BeginMap(25)
	must.NotError(err)
	for i := 1; i <= 25; i++ {
		must.NotError(ma.AssembleKey().AssignString(tableStrInt[i-1].s))
		must.NotError(ma.AssembleValue().AssignInt(tableStrInt[i-1].i))
	}
	must.NotError(ma.Done())
	n, err := nb.Build()
	must.NotError(err)
	return n
}

func CheckMaps(t *testing.T, ns ipld.NodeStyle) {
	t.Run("map node, str:int, 3 entries", func(t *testing.T) {
		n := buildMapStrIntN3(ns)
		t.Run("reads back out", func(t *testing.T) {
			wish.Wish(t, n.Length(), wish.ShouldEqual, 3)

			v, err := n.LookupString("whee")
			wish.Wish(t, err, wish.ShouldEqual, nil)
			v2, err := v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 1)

			v, err = n.LookupString("waga")
			wish.Wish(t, err, wish.ShouldEqual, nil)
			v2, err = v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 3)

			v, err = n.LookupString("woot")
			wish.Wish(t, err, wish.ShouldEqual, nil)
			v2, err = v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 2)
		})
		t.Run("reads via iteration", func(t *testing.T) {
			itr := n.MapIterator()

			wish.Wish(t, itr.Done(), wish.ShouldEqual, false)
			k, v, err := itr.Next()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			k2, err := k.AsString()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, k2, wish.ShouldEqual, "whee")
			v2, err := v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 1)

			wish.Wish(t, itr.Done(), wish.ShouldEqual, false)
			k, v, err = itr.Next()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			k2, err = k.AsString()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, k2, wish.ShouldEqual, "woot")
			v2, err = v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 2)

			wish.Wish(t, itr.Done(), wish.ShouldEqual, false)
			k, v, err = itr.Next()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			k2, err = k.AsString()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, k2, wish.ShouldEqual, "waga")
			v2, err = v.AsInt()
			wish.Wish(t, err, wish.ShouldEqual, nil)
			wish.Wish(t, v2, wish.ShouldEqual, 3)

			wish.Wish(t, itr.Done(), wish.ShouldEqual, true)
			k, v, err = itr.Next()
			wish.Wish(t, err, wish.ShouldEqual, ipld.ErrIteratorOverread{})
			wish.Wish(t, k, wish.ShouldEqual, nil)
			wish.Wish(t, v, wish.ShouldEqual, nil)
		})
		t.Run("reads for absent keys error sensibly", func(t *testing.T) {
			v, err := n.LookupString("nope")
			wish.Wish(t, err, wish.ShouldBeSameTypeAs, ipld.ErrNotExists{})
			wish.Wish(t, err.Error(), wish.ShouldEqual, `key not found: "nope"`)
			wish.Wish(t, v, wish.ShouldEqual, nil)
		})
	})
	t.Run("repeated key should error", func(t *testing.T) {
		nb := ns.NewBuilder()
		ma, err := nb.BeginMap(3)
		if err != nil {
			panic(err)
		}
		if err := ma.AssembleKey().AssignString("whee"); err != nil {
			panic(err)
		}
		if err := ma.AssembleValue().AssignInt(1); err != nil {
			panic(err)
		}
		if err := ma.AssembleKey().AssignString("whee"); err != nil {
			wish.Wish(t, err, wish.ShouldBeSameTypeAs, ipld.ErrRepeatedMapKey{})
			// No string assertion at present -- how that should be presented for typed stuff is unsettled
			//  (and if it's clever, it'll differ from untyped, which will mean no assertion possible!).
		}
	})
	t.Run("builder reset works", func(t *testing.T) {
		// TODO
	})
}