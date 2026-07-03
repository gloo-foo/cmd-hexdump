package hexdump_test

import (
	"fmt"
	"os"

	"github.com/gloo-foo/testable"
	"github.com/gloo-foo/testable/run"

	command "github.com/gloo-foo/cmd-hexdump"
)

// This example demonstrates reading from a file instead of inline input.
func ExampleHexdump_fromFile_basic() {
	// hexdump testdata/binary.txt
	data, err := os.ReadFile("testdata/binary.txt")
	if err != nil {
		fmt.Printf("read testdata: %v\n", err)
		return
	}
	output, _ := testable.Test(command.Hexdump(), run.Input(data))
	fmt.Print(output)
	// Output:
	// 48 65 6c 6c 6f
}
