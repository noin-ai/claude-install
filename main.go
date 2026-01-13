package main

import (
	"fmt"
	"os"

	"github.com/noin-ai/claude-install/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
