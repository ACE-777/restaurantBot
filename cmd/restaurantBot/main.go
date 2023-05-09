package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"restaurantBot/internal/pkg"
)

func main() {
	//bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	bot, err := tgbotapi.NewBotAPI("6017463153:AAEhp1iy3e4MuyxpyTyfoyuN64W3ZzQ3F-M")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Runing bot: %s ...", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			pkg.Commands(update, bot, update.Message)
		} else if update.CallbackQuery != nil {
			pkg.Callback(update, bot)
		}
	}
}
