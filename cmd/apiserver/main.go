package main

import (
	"github.com/BurntSushi/toml"
	"flag"
	//"fmt"
	"github.com/kirusha123/http-rest-api/iternal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init(){
	flag.StringVar(&configPath,"config-path","configs/apiserver.toml","path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath,config)

	if err != nil {
		log.Fatal(err)
	}

	APIServer := apiserver.New(config)

	if err := APIServer.Start(); err != nil{
		log.Fatal(err)
	}

}
