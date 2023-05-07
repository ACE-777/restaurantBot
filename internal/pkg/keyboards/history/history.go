package history

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Стандартное", "standard"),
		tgbotapi.NewInlineKeyboardButtonData("Банкет", "banquetHistory"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)
