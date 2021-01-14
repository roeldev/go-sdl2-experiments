// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dump

import (
	"fmt"
	"strings"
	"sync"
)

var dumpPool = sync.Pool{New: func() interface{} {
	return &dumper{}
}}

func newDumper(target interface{}) *dumper {
	d := dumpPool.Get().(*dumper)
	d.target = target
	d.buf.Reset()
	return d
}

func dump(d *dumper) string {
	res := d.Sprint()
	d.target = nil
	dumpPool.Put(d)
	return res
}

type dumper struct {
	buf    strings.Builder
	target interface{}
}

func (d *dumper) Const(n string, val interface{}, str, unknown string) {
	if str == "" {
		if unknown == "" {
			d.Number(n, val)
			return
		}
		str = unknown
	}
	_, _ = fmt.Fprintf(&d.buf, "\t%s: %s (%d)\n", n, str, val)
}

func (d *dumper) Number(n string, val interface{}) {
	_, f32 := val.(float32)
	_, f64 := val.(float64)

	format := "\t%s: %T(%d)\n"
	if f32 || f64 {
		format = "\t%s: %T(%f)\n"
	}

	_, _ = fmt.Fprintf(&d.buf, format, n, val, val)
}

func (d *dumper) Struct(n string, val *dumper) {
	_, _ = fmt.Fprintf(&d.buf, "\t%s: %s\n",
		n, strings.ReplaceAll(dump(val), "\n", "\n\t"),
	)
}

func (d *dumper) String(n string, val string) {
	_, _ = fmt.Fprintf(&d.buf, "\t%s: %s\n", n, val)
}

func (d *dumper) Sprint() string {
	if d.buf.Len() > 0 {
		return fmt.Sprintf("%T{\n%s}", d.target, d.buf.String())
	}

	return fmt.Sprintf("%T{}", d.target)
}
