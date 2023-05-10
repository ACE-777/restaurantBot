package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var StartKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Забронировать", "booking"),
		tgbotapi.NewInlineKeyboardButtonData("Банкет", "banquet"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("История", "history"),
		tgbotapi.NewInlineKeyboardButtonData("FAQ", "FAQ"),
	),
)

var MainMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Забронировать", "booking"),
		tgbotapi.NewInlineKeyboardButtonData("Банкет", "banquet"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("История", "history"),
		tgbotapi.NewInlineKeyboardButtonData("FAQ", "FAQ"),
	),
)

var MenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)
