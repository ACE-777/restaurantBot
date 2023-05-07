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

var EmptyKeyboard = tgbotapi.NewInlineKeyboardMarkup()

//var buttons = []tgbotapi.KeyboardButton{
//	tgbotapi.NewKeyboardButton("Button 1"),
//	tgbotapi.NewKeyboardButton("Button 2"),
//}
//
//var rows = [][]tgbotapi.KeyboardButton{
//	buttons,
//}
//
//var Keyboard = tgbotapi.ReplyKeyboardMarkup{
//	Keyboard:       rows,
//	ResizeKeyboard: true,
//}
