package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-hexdump"
)

// Hexdump re-exports the constructor.
func Hexdump(opts ...any) gloo.Command[[]byte, []byte] { return command.Hexdump(opts...) }

// -C flag: canonical hex+ASCII display
const Canonical = command.HexdumpCanonical

// default: not canonical
const NoCanonical = command.HexdumpNoCanonical
