package hexdump_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-hexdump"
	"github.com/gloo-foo/testable"
)

func ExampleHexdump_basic() {
	// echo "Hello" | hexdump
	output, _ := testable.Test(command.Hexdump(), "Hello\n")
	fmt.Print(output)
	// Output:
	// 48 65 6c 6c 6f
}

func ExampleHexdump_canonical() {
	// echo "Hello" | hexdump -C
	output, _ := testable.Test(command.Hexdump(command.HexdumpCanonical), "Hello\n")
	fmt.Print(output)
	// Output:
	// 00000000  48 65 6c 6c 6f                                    |Hello|
}
