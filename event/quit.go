// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

const Quit quit = "QUIT"

type quit string

func (q quit) ExitCode() int { return 0 }
func (q quit) Error() string { return string(q) }
