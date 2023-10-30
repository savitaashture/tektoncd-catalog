// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package sha1

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/internal/pkg"
)

func init() {
	pkg.Register("crypto/sha1", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name:  "Size",
		Const: "20",
	}, {
		Name:  "BlockSize",
		Const: "64",
	}, {
		Name: "Sum",
		Params: []pkg.Param{
			{Kind: adt.BytesKind | adt.StringKind},
		},
		Result: adt.BytesKind | adt.StringKind,
		Func: func(c *pkg.CallCtxt) {
			data := c.Bytes(0)
			if c.Do() {
				c.Ret = Sum(data)
			}
		},
	}},
}