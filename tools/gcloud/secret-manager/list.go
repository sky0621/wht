package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type listCmd struct {
}

func newListCmd() *listCmd {
	return &listCmd{}
}

func (*listCmd) Name() string {
	return "listCmd"
}

func (*listCmd) Synopsis() string {
	return "list secrets"
}

func (*listCmd) Usage() string {
	return `usage: list secrets`
}

func (cmd *listCmd) SetFlags(f *flag.FlagSet) {}

func (cmd *listCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
