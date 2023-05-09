package pkg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"restaurantBot/internal/pkg/keyboards"
)

func Commands(update tgbotapi.Update, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch update.Message.Text {
	case "/start":
		sendMessageWithKeyboard(bot, update.Message.Chat.ID, fmt.Sprintf("Привет %s !\nЯ помогу тебе сделать бронирование!", update.Message.From),
			keyboards.StartKeyboard)
	case "/help":
		sendMessageWithoutKeyboard(bot, update.Message.Chat.ID,
			"Доступные комманды:\n/menu - Главное меню бота\n/start - Старт бота\n/help - Показать список команд")
	case "/menu":
		sendMessageWithKeyboard(bot, update.Message.Chat.ID, "Главное меню:", keyboards.MainMenuKeyboard)
	default:
		sendMessageWithoutKeyboard(bot, update.Message.Chat.ID,
			"Неизвестная команда. Введи /help для получения списка команд")
	}
}
