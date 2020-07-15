package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Run(version string) {
	ver := flag.Bool("version", false, "version")

	flag.Parse()

	if *ver {
		fmt.Print(version)
		os.Exit(0)
	}
}
