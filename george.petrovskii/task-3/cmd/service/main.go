package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
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
	Valutes []Valute `xml:"Valute"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	xmlFile, err := os.Open(cfg.InputFile)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	err = decoder.Decode(&curs)
	if err != nil {
		panic(err)
	}

	fmt.Println("valutes count:", len(curs.Valutes))

	for _, v := range curs.Valutes {
		fmt.Println(v.NumCode, v.CharCode, v.Value)
	}
}
