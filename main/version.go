package main

import (
	"fmt"

	"github.com/decardlabs/uuhappyserver/core"
	"github.com/decardlabs/uuhappyserver/main/commands/base"
)

var cmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "Show current version of uuhappyserver",
	Long: `Version prints the build information for uuhappyserver executables.
	`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
