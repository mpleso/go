// Copyright © 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package builtin provides goes that run in the same context as the
// goes cli or script instead of as a sub-process.
package builtin

import (
	"github.com/platinasystems/go/goes/builtin/apropos"
	"github.com/platinasystems/go/goes/builtin/bang"
	"github.com/platinasystems/go/goes/builtin/cd"
	"github.com/platinasystems/go/goes/builtin/cli"
	"github.com/platinasystems/go/goes/builtin/complete"
	"github.com/platinasystems/go/goes/builtin/env"
	"github.com/platinasystems/go/goes/builtin/exit"
	"github.com/platinasystems/go/goes/builtin/export"
	"github.com/platinasystems/go/goes/builtin/help"
	"github.com/platinasystems/go/goes/builtin/license"
	"github.com/platinasystems/go/goes/builtin/man"
	"github.com/platinasystems/go/goes/builtin/patents"
	"github.com/platinasystems/go/goes/builtin/resize"
	"github.com/platinasystems/go/goes/builtin/show_commands"
	"github.com/platinasystems/go/goes/builtin/source"
	"github.com/platinasystems/go/goes/builtin/usage"
	"github.com/platinasystems/go/goes/builtin/version"
)

func New() []interface{} {
	return []interface{}{
		apropos.New(),
		bang.New(),
		cli.New(),
		cd.New(),
		complete.New(),
		env.New(),
		exit.New(),
		export.New(),
		help.New(),
		license.New(),
		man.New(),
		patents.New(),
		resize.New(),
		show_commands.New(),
		source.New(),
		usage.New(),
		version.New(),
	}
}