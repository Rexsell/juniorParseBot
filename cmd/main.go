package main

import (
	"juniorParseBot/internal/server"
	"juniorParseBot/internal/telegram"
	"juniorParseBot/model"
)
import log "github.com/sirupsen/logrus"

func main() {

	cfg, err := model.InitConfig("config.yaml")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	bot := telegram.New(cfg)

	serv := server.New(bot, 100)

	err = serv.Start(cfg)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
