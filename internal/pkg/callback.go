package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"restaurantBot/internal/pkg/keyboards"
	"restaurantBot/internal/pkg/keyboards/FAQ"
	"restaurantBot/internal/pkg/keyboards/history"
)

func Callback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch update.CallbackQuery.Data {

	case "history":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип бронирования", history.MainKeyboard)
	case "standard":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История стандартных бронирований:", keyboards.MenuKeyboard)
	case "banquetHistory":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История банкетных бронирований:", keyboards.MenuKeyboard)

	case "FAQ":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Ресторан работает:\nПн-Вс 10:00-23:00\nАдрес пр. Динамо, 2А, Санкт-Петербург", FAQ.MainKeyboard)
	case "menu":
		sendMessageWithoutKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Для того, чтобы перейти в главное меню введите /menu")
	case "socialNetwork":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Наши социальные ресурсы", FAQ.SocialNetworkKeyboard)
	case "contacts":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Телефон: 8 (911) 752-55-55\nЧат ссобщества Vk https://vk.com/im?sel=-176492535", keyboards.MenuKeyboard)
	}

	//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	//if _, err := bot.Request(callback); err != nil {
	//	panic(err)
	//}
	//
	//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "cflooo")
	//if _, err := bot.Send(msg); err != nil {
	//	panic(err)
	//}
}
