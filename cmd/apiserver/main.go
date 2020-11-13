package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kirusha123/rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	cfg := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, cfg)

	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
