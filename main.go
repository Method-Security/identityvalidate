package main

import (
	"flag"
	"os"

	"github.com/Method-Security/identityvalidation/cmd"
)

var Version = "none"

func main() {
	flag.Parse()

	identityvalidation := cmd.NewIdentityValidation(Version)

	if err := identityvalidation.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
