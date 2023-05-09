package pkg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"restaurantBot/internal/pkg/keyboards"
	"restaurantBot/internal/pkg/keyboards/FAQ"
	"restaurantBot/internal/pkg/keyboards/banquet"
	"restaurantBot/internal/pkg/keyboards/booking"
	"restaurantBot/internal/pkg/keyboards/history"
	"strings"
)

func addButton(text, callbackData string) (keyboard tgbotapi.InlineKeyboardMarkup) {
	keyboard = tgbotapi.NewInlineKeyboardMarkup()
	row := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(text, callbackData),
	)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	return
}

func Callback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch update.CallbackQuery.Data {
	case "menu":
		sendMessageWithoutKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Для того, чтобы перейти в главное меню введите /menu")

	case "booking":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип бронирования", booking.MainKeyboard)
	case "firstFloor":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип бронирования", booking.FirstFloorKeyboard)
	case "secondFloor":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип бронирования", booking.SecondFloorKeyboard)
	case "bar":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00BarBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "table":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00TableBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "fireplace":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00FireplaceBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "scene":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00SceneBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "skipFirst":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00SkipFirst")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "skipSecond":
		//clickhouse
		keyboard := addButton("1.05 10:00", "1.05_10:00SkipSecond")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)

	case "banquet":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип пакета", banquet.MainKeyboard)
	case "personalize":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Спасибо за Вашу заявку.\nВ скором времени,наш менеджер свяжется с Вами для подробного обсуждения всех моментов!",
			banquet.PersonalizeKeyboard)
	case "ready":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"В нашем ресторане у нас есть два готовых пакета для Ваших мероприятий: два блюда + алкоголь на выбор "+
				"или два блюда без алкоголя ресторана, но с пробочным сбором (40% от стоимости всего алкоголя).\n Какой "+
				"подойдет Вам больше?", banquet.ReadyKeyboard)

	case "withFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете первое блюдо",
			banquet.FeeKeyboard)
	case "VegetableFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете второе блюдо",
			banquet.FeeKeyboardTwoVegetable)
	case "MeatFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете второе блюдо",
			banquet.FeeKeyboardTwoMeat)
	case "sorbetFeeVegetable":
		//clickhouse
		keyboard := addButton("7.05 18:00", "7.05_18:00sorbetFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuFeeVegetable":
		//clickhouse
		keyboard := addButton("6.05 18:00", "6.05_18:00tiramisuFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "sorbetFeeMeat":
		//clickhouse
		keyboard := addButton("5.05 10:00", "5.05_10:00sorbetFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuFeeMeat":
		//clickhouse
		keyboard := addButton("3.05 14:00", "3.05_14:00tiramisuFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)

	case "withoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете первое блюдо",
			banquet.WithoutFeeKeyboard)
	case "VegetableWithoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете второе блюдо",
			banquet.WithoutFeeKeyboardTwoVegetable)
	case "MeatWithoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете второе блюдо",
			banquet.WithoutFeeKeyboardTwoMeat)
	case "sorbetWithoutFeeMeat":
		//clickhouse
		keyboard := addButton("3.05 14:00", "3.05_14:00sorbetWithoutFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuWithoutFeeMeat":
		//clickhouse
		keyboard := addButton("3.05 14:00", "3.05_14:00tiramisuWithoutFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "sorbetWithoutFeeVegetable":
		//clickhouse
		keyboard := addButton("3.05 14:00", "3.05_14:00sorbetWithoutFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuWithoutFeeVegetable":
		//clickhouse
		keyboard := addButton("3.05 14:00", "3.05_14:00tiramisuWithoutFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)

	case "history":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выбери тип бронирования", history.MainKeyboard)
	case "standard":
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История стандартных бронирований:", keyboards.MenuKeyboard)
	case "banquetHistory":
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История банкетных бронирований:", keyboards.MenuKeyboard)

	case "FAQ":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Ресторан работает:\nПн-Вс 10:00-23:00\nАдрес пр. Динамо, 2А, Санкт-Петербург", FAQ.MainKeyboard)
	case "socialNetwork":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Наши социальные ресурсы", FAQ.SocialNetworkKeyboard)
	case "contacts":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Телефон: 8 (911) 752-55-55\nЧат ссобщества Vk https://vk.com/im?sel=-176492535", keyboards.MenuKeyboard)

	default:
		getCallback(update, bot)
	}
}

func getCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if strings.Contains(update.CallbackData(), "BarBooking") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "BarBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "TableBooking") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "TableBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "FireplaceBooking") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "FireplaceBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "SceneBooking") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "SceneBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "skipFirst") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "skipFirst")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "skipSecond") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "skipSecond")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetFeeVegetable") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuFeeVegetable") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, овощным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetFeeMeat") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, мясным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuFeeMeat") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, мясным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetWithoutFeeMeat") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, мясным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetWithoutFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuWithoutFeeMeat") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, мясным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuWithoutFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetWithoutFeeVegetable") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetWithoutFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuWithoutFeeVegetable") {
		//clickhouse
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuWithoutFeeVegetable")[0]), keyboards.MenuKeyboard)
	}
}
