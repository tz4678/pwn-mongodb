package main

import (
	"os"
	"pwn-mongodb/internal"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cfg := &internal.Config{}

	kingpin.Flag("timeout", "Connect timeout").
		Short('t').
		Default("5s").
		DurationVar(&cfg.Timeout)

	kingpin.Flag("concurrency", "Number of concurrent workers").
		Short('c').
		Default("10").
		IntVar(&cfg.Concurrency)

	kingpin.Flag("rate-limit", "Requests per second").
		Short('l').
		Default("100").
		IntVar(&cfg.RateLimit)

	kingpin.Flag("input", "Input filename").
		Short('i').
		Default(os.Stdin.Name()).
		FileVar(&cfg.Input)

	kingpin.Flag("output", "Output filename").
		Short('o').
		Default(os.Stdout.Name()).
		OpenFileVar(&cfg.Output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	err := internal.Run(cfg)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
