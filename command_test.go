package command_test

import (
	"fmt"
	"strings"
	"testing"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-hexdump"
)

// dump runs the command over input and returns the rendered lines, failing the
// test on any execution error.
func dump(t *testing.T, cmd gloo.Command[[]byte, []byte], input string) []string {
	t.Helper()
	lines, err := testable.TestLines(cmd, input)
	if err != nil {
		t.Fatal(err)
	}
	return lines
}

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %d lines %q, want %d lines %q", len(got), got, len(want), want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("line %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestHexdump_DefaultRendersSpaceSeparatedHex(t *testing.T) {
	got := dump(t, command.Hexdump(), "abc\n")
	assertLines(t, got, []string{"61 62 63"})
}

func TestHexdump_DefaultPreservesMultipleLines(t *testing.T) {
	got := dump(t, command.Hexdump(), "Hi\nok\n")
	assertLines(t, got, []string{"48 69", "6f 6b"})
}

func TestHexdump_DefaultEmptyInputProducesNoLines(t *testing.T) {
	got := dump(t, command.Hexdump(), "")
	assertLines(t, got, nil)
}

func TestHexdump_DefaultBlankLineRendersEmpty(t *testing.T) {
	got := dump(t, command.Hexdump(), "\n")
	assertLines(t, got, []string{""})
}

func TestHexdump_CanonicalRendersOffsetHexAndAscii(t *testing.T) {
	got := dump(t, command.Hexdump(command.HexdumpCanonical), "Hi\n")
	assertLines(t, got, []string{
		"00000000  48 69                                             |Hi|",
	})
}

func TestHexdump_CanonicalNonPrintableBytesShowAsDots(t *testing.T) {
	// NUL (0x00) and DEL (0x7f) are non-printable and collapse to '.' in the
	// sidebar, while their hex columns still render the true byte values.
	got := dump(t, command.Hexdump(command.HexdumpCanonical), "A\x00B\x7f\n")
	assertLines(t, got, []string{
		"00000000  41 00 42 7f                                       |A.B.|",
	})
}

func TestHexdump_CanonicalFullRowInsertsMidpointGap(t *testing.T) {
	// A full 16-byte row exercises the gap inserted after the eighth byte.
	got := dump(t, command.Hexdump(command.HexdumpCanonical), "0123456789abcdef\n")
	assertLines(t, got, []string{
		"00000000  30 31 32 33 34 35 36 37  38 39 61 62 63 64 65 66  |0123456789abcdef|",
	})
}

func TestHexdump_CanonicalRendersEachLineIndependently(t *testing.T) {
	got := dump(t, command.Hexdump(command.HexdumpCanonical), "abc\nxyz\n")
	assertLines(t, got, []string{
		"00000000  61 62 63                                          |abc|",
		"00000000  78 79 7a                                          |xyz|",
	})
}

func TestHexdump_NoCanonicalMatchesDefault(t *testing.T) {
	// The disabled flag form must behave exactly like passing no flag at all.
	got := dump(t, command.Hexdump(command.HexdumpNoCanonical), "abc\n")
	assertLines(t, got, []string{"61 62 63"})
}

func ExampleHexdump() {
	lines, _ := testable.TestLines(command.Hexdump(), "Hello\n")
	fmt.Println(strings.Join(lines, "\n"))
	// Output:
	// 48 65 6c 6c 6f
}

func ExampleHexdump_canonical() {
	lines, _ := testable.TestLines(command.Hexdump(command.HexdumpCanonical), "Hi\n")
	fmt.Println(strings.Join(lines, "\n"))
	// Output:
	// 00000000  48 69                                             |Hi|
}
