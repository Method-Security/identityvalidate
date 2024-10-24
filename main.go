package main

import (
	"flag"
	"os"

	"github.com/Method-Security/identityvalidate/cmd"
)

var Version = "none"

func main() {
	flag.Parse()

	identityvalidate := cmd.Newidentityvalidate(Version)
	identityvalidate.InitRootCommand()
	identityvalidate.InitPortalCommand()

	if err := identityvalidate.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
