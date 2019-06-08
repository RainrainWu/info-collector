package main

import (
	"net/http"
	"os"

	"info-collector/pkg/tgbot"
)

func main() {

	bot := tgbot.Tgbot{TunaBatch:2, BnextBatch:4, MeetBatch:3}
	bot.AuthBot()
	bot.Reply()

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

