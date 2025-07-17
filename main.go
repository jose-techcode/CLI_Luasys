package main

import (
	"github.com/jose-techcode/CLI_Luasys/cmd"
	"github.com/jose-techcode/CLI_Luasys/api"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "api" {
		api.Server()
	} else {
		cmd.Execute()
	}
}