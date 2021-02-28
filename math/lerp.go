// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

func Lerp(cur, dest, t float64) float64 { return cur + (dest-cur)*t }

func LerpRound(cur, dest, t, e float64) float64 {
	cur += (dest - cur) * t
	if e > 0 {
		if dest > cur && dest-e < cur {
			cur = dest
		} else if dest < cur && dest+e > cur {
			cur = dest
		}
	}
	return cur
}
