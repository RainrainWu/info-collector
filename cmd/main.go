package main

import (
	"net/http"
	"os"

	"info-collector/pkg/tgbot"
)

func main() {

	http.HandleFunc("/", bot)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func bot (w http.ResponseWriter, r *http.Request) {

	bot := tgbot.Tgbot{TunaBatch:2, BnextBatch:4, MeetBatch:3}
        bot.AuthBot()
        bot.Reply()
}
