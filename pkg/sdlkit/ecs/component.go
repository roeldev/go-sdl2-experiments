// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecs

type ComponentTag uint64

func (ct ComponentTag) Flags() []ComponentTag {
	res := make([]ComponentTag, 0, 64)

	var flag ComponentTag
	for i := 0; i < 64; i++ {
		flag = 1 << i
		if ct&flag != 0 {
			res = append(res, flag)
		}
	}
	return res
}
