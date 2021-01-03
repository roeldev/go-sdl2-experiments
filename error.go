// Copyright (c) 2020, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdlkit

import (
	"fmt"
	"os"

	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/go-pogo/sdlkit/event"
)

var FatalErrorTitle = "fatal error"

func FailOnErr(possibleErr error) {
	if possibleErr == nil || errors.Is(possibleErr, event.Quit) {
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
