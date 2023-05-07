package FAQ

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ресурсы", "socialNetwork"),
		tgbotapi.NewInlineKeyboardButtonData("Контакты", "contacts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)

var SocialNetworkKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Сайт", "https://fishzlgroup.ru/"),
		tgbotapi.NewInlineKeyboardButtonURL("Vk", "https://vk.com/fishzlgroup"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)
