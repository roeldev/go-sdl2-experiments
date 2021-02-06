// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"fmt"
	"os"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

var FatalErrorTitle = "fatal error"

const QUIT quit = "QUIT"

type quit string

func (q quit) ExitCode() int { return 0 }
func (q quit) Error() string { return string(q) }

// FailOnErr shows a simple message box and exits the program when it receives
// a non-nil error.
func FailOnErr(possibleErr error) {
	if possibleErr == nil || errors.Is(possibleErr, QUIT) {
		return
	}

	_ = sdl.ShowSimpleMessageBox(
		sdl.MESSAGEBOX_ERROR,
		FatalErrorTitle,
		possibleErr.Error(),
		nil,
	)

	exitCode := 1
	if exitErr, ok := possibleErr.(errors.ExitCoder); ok {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n%+v", FatalErrorTitle, possibleErr)
		exitCode = exitErr.ExitCode()
	}

	os.Exit(exitCode)
}
