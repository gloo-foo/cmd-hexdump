package command

// hexdumpCanonicalFlag enables the canonical hex+ASCII display (-C).
type hexdumpCanonicalFlag bool

const (
	HexdumpCanonical   hexdumpCanonicalFlag = true
	HexdumpNoCanonical hexdumpCanonicalFlag = false
)

// flags holds the parsed hexdump flag set.
type flags struct {
	canonicalEnabled hexdumpCanonicalFlag
}

// fold partitions opts: hexdump's own option values are folded into the flag
// set, and every other argument is passed through unchanged for the framework's
// positional classifier.
func fold(opts []any) (flags, []any) {
	var f flags
	rest := make([]any, 0, len(opts))
	for _, o := range opts {
		switch v := o.(type) {
		case hexdumpCanonicalFlag:
			f.canonicalEnabled = v
		default:
			rest = append(rest, o)
		}
	}
	return f, rest
}
