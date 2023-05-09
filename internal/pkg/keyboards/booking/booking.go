package booking

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Первый этаж", "firstFloor"),
		tgbotapi.NewInlineKeyboardButtonData("Второй этаж", "secondFloor"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)

var FirstFloorKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Бар", "bar"),
		tgbotapi.NewInlineKeyboardButtonData("Стол", "table"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
		tgbotapi.NewInlineKeyboardButtonData("Пропустить", "skipFirst"),
	),
)

var SecondFloorKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Камин", "fireplace"),
		tgbotapi.NewInlineKeyboardButtonData("Сцена", "scene"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
		tgbotapi.NewInlineKeyboardButtonData("Пропустить", "skipSecond"),
	),
)
