package command

type hexdumpCanonicalFlag bool

const (
	HexdumpCanonical   hexdumpCanonicalFlag = true
	HexdumpNoCanonical hexdumpCanonicalFlag = false
)

type flags struct {
	canonical hexdumpCanonicalFlag
}

func (c hexdumpCanonicalFlag) Configure(flags *flags) { flags.canonical = c }
