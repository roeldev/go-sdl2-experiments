// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentTag_Flags(t *testing.T) {
	tests := []struct {
		ct   ComponentTag
		want []ComponentTag
	}{
		{
			ct:   1 | 2,
			want: []ComponentTag{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.want, tc.ct.Flags())
		})
	}
}
