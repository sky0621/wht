package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	os.Exit(int(execMain()))
}

func execMain() subcommands.ExitStatus {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(newCreateCmd(), "create")
	subcommands.Register(newListCmd(), "list")
	flag.Parse()
	return subcommands.Execute(context.Background())
}
