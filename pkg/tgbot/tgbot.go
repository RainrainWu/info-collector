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
	MeetBatch	int
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
		} else if update.Message.Text == "/help" {

			content := `Hi I'm info-collector, a service to collect tech news and stories, you can use the following commands to chat with me.
			
			tuna_update	吐納商業評論
			bnext_update	數位時代
			meet_update	創業小聚
			
			Author: Rain
			Repo:	https://github.com/RainrainWu/info-collector`
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, content)
			msg.ReplyToMessageID = update.Message.MessageID
			tgbot.Bot.Send(msg)
		} else if update.Message.Text == "/tuna_update" {

			tgbot.BypassReply(crawler.Tuna_press(), tgbot.TunaBatch, update)
		} else if update.Message.Text == "/bnext_update" {

			tgbot.BypassReply(crawler.Business_next(), tgbot.BnextBatch, update)
		} else if update.Message.Text == "/meet_update" {

			tgbot.BypassReply(crawler.Meet_bnext(), tgbot.MeetBatch, update)
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
