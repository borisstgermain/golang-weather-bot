package main

import (
	"fmt"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	BOT_TOKEN = "238663056:AAE3aizY-QZ-UTQcR9GmdXKDewOJ1-Ersh0"
	API_KEY   = "2446ff5aa914f740113b505a5fddca4b"
)

func main() {
	reqWeather := ReqWeather{}
	reqWeather.Init(API_KEY)
	bot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	fmt.Printf("%s is ready for action!!11", bot.Self.UserName)

	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// TODO проверка сообщения
		args := strings.Split(update.Message.Text, " ")
		city := args[0]

		text := city + ":\n" + reqWeather.GetWeather(args)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	}
}
