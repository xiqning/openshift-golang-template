package main

import (
	"os"

	"github.com/amsokol/openshift-golang-template/example/pkg/cmd/cli"
	"github.com/amsokol/openshift-golang-template/example/pkg/cmd/server"
)

func main() {
	if err := server.Start(cli.GetPort()); err != nil {
		os.Exit(1)
	}
}
