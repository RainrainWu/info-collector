package main

import (
	"net/http"
	"os"
	"fmt"

	"info-collector/pkg/tgbot"
)

func main() {

	bot := tgbot.Tgbot{TunaBatch:2, BnextBatch:4, MeetBatch:3}
        bot.AuthBot()
        bot.Reply()

	updates := bot.Bot.ListenForWebhook("/" + bot.Bot.Token)
	port := os.Getenv("PORT")
	go http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", nil)

	for update := range updates {

		fmt.Println("%+v\n", update)
	}
}
