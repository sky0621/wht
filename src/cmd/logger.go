package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func globalLogSetting() {
	// FIXME: for GCP logging
}

func globalLocalLogSetting() {
	// Pretty logging
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}
