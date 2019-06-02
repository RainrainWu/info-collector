package tgbot

import (
	"log"

	"info-collector/pkg/crawler"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Tgbot struct {

	Bot		*tgbotapi.BotAPI
	TunaBatch	int
	BnextBatch	int
}

func (tgbot *Tgbot) AuthBot() {

	bot, err := tgbotapi.NewBotAPI("770132626:AAGckrbOi_95zZ8jOo9YS28wXLJy0MK_i_0")
	if err != nil {

		log.Panic(err)
	}

	tgbot.Bot = bot
	tgbot.Bot.Debug = true
	log.Printf("Authorized on account %s", tgbot.Bot.Self.UserName)
}


func (tgbot *Tgbot) Reply() {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120
	updates, err := tgbot.Bot.GetUpdatesChan(u)
	if err != nil {

		log.Panic(err)
	}

	for update := range updates {

		if update.Message == nil {

			continue
		} else if update.Message.Text == "/tuna_update" {

			tgbot.BypassReply(crawler.Tuna_press(), tgbot.TunaBatch, update)
		} else if update.Message.Text == "/bnext_update" {

			tgbot.BypassReply(crawler.Business_next(), tgbot.BnextBatch, update)
		}
	}
}

func (tgbot *Tgbot) BypassReply(ch chan string, batch int, update tgbotapi.Update) {

	for i := 0; i < batch; i++ {

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, <-ch)
		msg.ReplyToMessageID = update.Message.MessageID
		tgbot.Bot.Send(msg)
	}
}
