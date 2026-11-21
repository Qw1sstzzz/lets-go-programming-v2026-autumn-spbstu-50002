package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CommaFloat float64

func (c *CommaFloat) UnmarshalText(text []byte) error {
	s := strings.ReplaceAll(string(text), ",", ".")

	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return err
	}

	*c = CommaFloat(f)

	return nil
}

type Valute struct {
	NumCode  int        `xml:"NumCode"`
	CharCode string     `xml:"CharCode"`
	Value    CommaFloat `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is not set, use --config")
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Errorf("read config: %w", err))
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("parse config: %w", err))
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("config must contain input-file and output-file")
	}

	xmlFile, err := os.Open(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("open input file: %w", err))
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err = decoder.Decode(&curs); err != nil {
		panic(fmt.Errorf("decode xml: %w", err))
	}

	sort.SliceStable(curs.Valutes, func(i, j int) bool {
		return curs.Valutes[i].Value > curs.Valutes[j].Value
	})

	fmt.Println("valutes count:", len(curs.Valutes))

	for _, v := range curs.Valutes {
		fmt.Println(v.NumCode, v.CharCode, v.Value)
	}
}
