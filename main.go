package main

import (
	"github.com/promptc/pot/oper/help"
	"github.com/promptc/pot/oper/pack"
	"github.com/promptc/pot/oper/remote"
	"github.com/promptc/pot/oper/shared"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help.Welcome()
		return
	}
	shared.InitPath()
	verb := args[0]
	args = args[1:]
	handler := verbParser(verb)
	if handler == nil {
		help.Welcome()
		return
	}
	handler(args)
}

func verbParser(verb string) func(args []string) {
	switch verb {
	case "help":
		return help.HelpHandler
	case "update":
		return remote.UpdateHandler
	case "upgrade":
		return remote.UpgradeHandler
	case "remove":
		return pack.RemoveHandler
	case "install":
		return pack.InstallHandler
	case "list":
		return pack.ListHandler
	case "search":
		return pack.SearchHandler
	}
	return nil
}
