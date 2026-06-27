package command

import (
	"fmt"
	"strings"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// bytesPerLine is the number of input bytes rendered per canonical row, matching
// the 16-column layout of hexdump -C / xxd.
const bytesPerLine = 16

// midpoint is the column after which the canonical layout inserts an extra gap,
// splitting the hex field into two octets of eight bytes.
const midpoint = 8

// Hexdump returns a command that produces a hex dump of each input line.
//
// Default: space-separated hex bytes for each line.
// With HexdumpCanonical (-C): offset, hex bytes with midpoint gap, and ASCII sidebar.
func Hexdump(opts ...any) gloo.Command[[]byte, []byte] {
	params := gloo.NewParameters[gloo.File, flags](opts...)
	render := renderer(bool(params.Flags.canonical))
	return patterns.Map(func(line []byte) ([]byte, error) {
		return []byte(render(line)), nil
	})
}

// renderer selects the line formatter for the active mode.
func renderer(canonical bool) func([]byte) string {
	if canonical {
		return formatCanonical
	}
	return formatHex
}

// formatHex returns space-separated lowercase hex bytes.
func formatHex(data []byte) string {
	parts := make([]string, len(data))
	for i, b := range data {
		parts[i] = fmt.Sprintf("%02x", b)
	}
	return strings.Join(parts, " ")
}

// formatCanonical returns offset + hex field + ASCII sidebar for a single line.
func formatCanonical(data []byte) string {
	return fmt.Sprintf("%08x  %s |%s|", 0, hexField(data), asciiField(data))
}

// hexField renders the fixed-width 16-column hex field with a midpoint gap,
// padding absent bytes with spaces so the ASCII sidebar always aligns.
func hexField(data []byte) string {
	var b strings.Builder
	for i := range bytesPerLine {
		if i == midpoint {
			b.WriteByte(' ')
		}
		b.WriteString(hexCell(data, i))
	}
	return b.String()
}

// hexCell renders one column: the byte as "%02x " or three spaces when absent.
func hexCell(data []byte, i int) string {
	if i < len(data) {
		return fmt.Sprintf("%02x ", data[i])
	}
	return "   "
}

// asciiField renders the ASCII sidebar: printable bytes verbatim, others as '.'.
func asciiField(data []byte) string {
	var b strings.Builder
	for _, c := range data {
		b.WriteByte(printable(c))
	}
	return b.String()
}

// printable maps a byte to its sidebar glyph: itself if printable, else '.'.
func printable(c byte) byte {
	if c >= 0x20 && c <= 0x7e {
		return c
	}
	return '.'
}
