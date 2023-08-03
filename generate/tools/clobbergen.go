//go:build ignore

package main

import (
	"os"

	"github.com/progrium/macdriver/generate/oldgen"
)

// go run ./generate/tools/clobbergen.go [dir, ex: ./macos/appkit]
func main() {
	oldgen.RemoveGeneratedCode(os.Args[1])
}
