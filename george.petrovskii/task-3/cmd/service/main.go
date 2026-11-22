package main

import (
	"errors"
	"flag"

	"george.petrovskii/task-3/internal/config"
	"george.petrovskii/task-3/internal/currency"
	"george.petrovskii/task-3/internal/writer"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		return errors.New("config path is not set, use --config")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return err
	}

	currencies, err := currency.Load(cfg.InputFile)
	if err != nil {
		return err
	}

	currency.SortByValueDesc(currencies)

	return writer.WriteJSON(cfg.OutputFile, currencies)
}
