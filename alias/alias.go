package alias

import command "github.com/gloo-foo/cmd-hexdump"

// Hexdump is the command constructor.
var Hexdump = command.Hexdump

// -C flag: canonical hex+ASCII display
const Canonical = command.HexdumpCanonical

// default: not canonical
const NoCanonical = command.HexdumpNoCanonical
