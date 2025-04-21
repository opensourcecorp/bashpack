package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	if err := newRootCmd().Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
