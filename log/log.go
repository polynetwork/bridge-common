/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

const (
	FATAL = log.LvlCrit
	ERROR = log.LvlError
	WARN  = log.LvlWarn
	INFO  = log.LvlInfo
	DEBUG = log.LvlDebug
	TRACE = log.LvlTrace
)

var (
	JSON      = false
	VERBOSITY = INFO
	VMODULE   = ""

	glogger *log.GlogHandler

	New    = log.New
	Root   = log.Root
	Output = log.Output
	Trace  = log.Trace
	Debug  = log.Debug
	Info   = log.Info
	Warn   = log.Warn
	Error  = log.Error
	Crit   = log.Crit
	Fatal  = log.Crit
)

func init() {
	glogger = log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	verbosity, _ := strconv.Atoi(os.Getenv("LOG"))
	if verbosity > 0 {
		VERBOSITY = log.Lvl(verbosity)
	}
	glogger.Verbosity(VERBOSITY)
	log.Root().SetHandler(glogger)
}

func Init() {
	var ostream log.Handler
	output := io.Writer(os.Stderr)
	if JSON {
		ostream = log.StreamHandler(output, log.JSONFormat())
	} else {
		usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
		if usecolor {
			output = colorable.NewColorableStderr()
		}
		ostream = log.StreamHandler(output, log.TerminalFormat(usecolor))
	}
	glogger.SetHandler(ostream)

	// logging
	if VERBOSITY > 0 {
		glogger.Verbosity(log.Lvl(VERBOSITY))
	}
	if VMODULE != "" {
		glogger.Vmodule(VMODULE)
	}
	/*
		debug := ctx.GlobalBool(debugFlag.Name)
		if ctx.GlobalIsSet(debugFlag.Name) {
			debug = ctx.GlobalBool(debugFlag.Name)
		}

		log.PrintOrigins(debug)

		backtrace := ctx.GlobalString(backtraceAtFlag.Name)
		glogger.BacktraceAt(backtrace)
	*/

	log.Root().SetHandler(glogger)
}

func Json(lvl log.Lvl, body interface{}) {
	v, _ := json.Marshal(body)
	if lvl <= VERBOSITY {
		fmt.Println(string(v))
	}
}
