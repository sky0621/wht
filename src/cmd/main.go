package main

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/sky0621/wht/cmd/setup"
)

func main() {
	os.Exit(int(setup.ExecMain()))
}
