package alias_test

import (
	"slices"
	"testing"

	"github.com/gloo-foo/testable"

	hexdump "github.com/gloo-foo/cmd-hexdump/alias"
)

// The alias package re-exports the constructor and flag constants under
// unprefixed names. A mis-wired re-export (say, Canonical bound to the disabled
// constant, or Hexdump bound to the wrong function) compiles cleanly, so only
// behavior can prove the wiring. Each test exercises one re-export and asserts
// the exact hexdump output it must produce.

const input = "Hi\n"

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_DefaultRendersSpaceSeparatedHex(t *testing.T) {
	lines, err := testable.TestLines(hexdump.Hexdump(), input)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"48 69"})
}

func TestAlias_CanonicalRendersOffsetHexAndAscii(t *testing.T) {
	// -C: offset, padded hex field, and ASCII sidebar.
	lines, err := testable.TestLines(hexdump.Hexdump(hexdump.Canonical), input)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{
		"00000000  48 69                                             |Hi|",
	})
}

func TestAlias_NoCanonicalMatchesDefault(t *testing.T) {
	// The NoCanonical constant is the disabled form: it must behave exactly
	// like passing no flag at all.
	lines, err := testable.TestLines(hexdump.Hexdump(hexdump.NoCanonical), input)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"48 69"})
}
