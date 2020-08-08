package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func globalLogSetting() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
}

func globalLocalLogSetting() {
	// Pretty logging
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}
	log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
}
