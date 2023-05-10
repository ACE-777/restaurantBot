package pkg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"restaurantBot/internal/pkg/adapter"
	"restaurantBot/internal/pkg/clickHouse"
	"restaurantBot/internal/pkg/keyboards"
	"restaurantBot/internal/pkg/keyboards/FAQ"
	"restaurantBot/internal/pkg/keyboards/banquet"
	"restaurantBot/internal/pkg/keyboards/booking"
	"restaurantBot/internal/pkg/keyboards/history"
	"strings"
)

func Callback(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *clickHouse.DBConnect) {
	userDatabaseAdapter := adapter.CreateAdapter(db)
	switch update.CallbackQuery.Data {
	case "menu":
		sendMessageWithoutKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Для того, чтобы перейти в главное меню введите /menu")

	case "booking":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выберите тип бронирования", booking.MainKeyboard)
	case "firstFloor":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выберите тип бронирования", booking.FirstFloorKeyboard)
	case "secondFloor":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выберите тип бронирования", booking.SecondFloorKeyboard)
	case "bar":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "table":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "fireplace":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "scene":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "skipFirst":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "skipSecond":
		keyboard := makeKeyboardBooking(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)

	case "banquet":
		checkAvailabilityBanquet(update, bot, db)

	case "personalize":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Спасибо за Вашу заявку.\nВ скором времени, наш менеджер свяжется с Вами для подробного обсуждения всех моментов!",
			banquet.PersonalizeKeyboard)
	case "ready":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"В нашем ресторане у нас есть два готовых пакета для Ваших мероприятий: два блюда + алкоголь на выбор "+
				"или два блюда без алкоголя ресторана, но с пробочным сбором (40% от стоимости всего алкоголя).\n Какой "+
				"подойдет Вам больше?", banquet.ReadyKeyboard)

	case "withFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете первое блюдо:",
			banquet.FeeKeyboard)
	case "VegetableFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете десерт:",
			banquet.FeeKeyboardTwoVegetable)
	case "MeatFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете десерт:",
			banquet.FeeKeyboardTwoMeat)
	case "sorbetFeeVegetable":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuFeeVegetable":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "sorbetFeeMeat":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)
	case "tiramisuFeeMeat":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю", keyboard)

	case "withoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете первое блюдо:",
			banquet.WithoutFeeKeyboard)
	case "VegetableWithoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете десерт:",
			banquet.WithoutFeeKeyboardTwoVegetable)
	case "MeatWithoutFee":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, выберете десерт:",
			banquet.WithoutFeeKeyboardTwoMeat)
	case "sorbetWithoutFeeMeat":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю:", keyboard)
	case "tiramisuWithoutFeeMeat":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю:", keyboard)
	case "sorbetWithoutFeeVegetable":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю:", keyboard)
	case "tiramisuWithoutFeeVegetable":
		keyboard := makeKeyboardBanquet(db, update)
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Пожалуйста, укажите предпочтительную дату и время из доступных на ближайшую неделю:", keyboard)

	case "history":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выберите тип бронирования", history.MainKeyboard)
	case "standard":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История стандартных бронирований:\n"+GetUsersHistory(userDatabaseAdapter, update, "standard")+"\n",
			keyboards.MenuKeyboard)
	case "banquetHistory":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"История банкетных бронирований:\n"+GetUsersHistory(userDatabaseAdapter, update, "banquetHistory"),
			keyboards.MenuKeyboard)

	case "FAQ":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Ресторан работает:\nПн-Вс 10:00-23:00\nАдрес пр. Динамо, 2А, Санкт-Петербург", FAQ.MainKeyboard)
	case "socialNetwork":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Наши социальные ресурсы:", FAQ.SocialNetworkKeyboard)
	case "contacts":
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Телефон: 8 (911) 752-55-55\nЧат ссобщества Vk https://vk.com/im?sel=-176492535", keyboards.MenuKeyboard)

	default:
		getCallback(update, bot, db)
	}
}

func getCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *clickHouse.DBConnect) {
	userDatabaseAdapter := adapter.CreateAdapter(db)
	if strings.Contains(update.CallbackData(), "barBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронировано место на первом этаже на барной стойке на ", "barBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "barBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tableBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронирован стол на первом этаже на ", "tableBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "tableBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "fireplaceBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронирован стол на втором этаже у камина на ", "fireplaceBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "fireplaceBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sceneBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронирован стол на втором этаже у сцены на ", "sceneBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "sceneBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "skipFirstBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронирован стол на первом этаже на ", "skipFirstBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "skipFirstBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "skipSecondBooking") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "standard",
			"Забронирован стол на втором этаже на ", "skipFirstBooking")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы забронировали на %s", strings.Split(update.CallbackQuery.Data, "skipSecondBooking")[0]),
			keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetFeeVegetable") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет с пробковым сбором.\nПервое блюдо: овощной салат.\nВторое блюдо: сорбет.\nБанкет забронирован на ", "sorbetFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuFeeVegetable") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет с пробковым сбором.\nПервое блюдо: овощной салат.\nВторое блюдо: тирамису.\nБанкет забронирован на ", "tiramisuFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, овощным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetFeeMeat") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет с пробковым сбором.\nПервое блюдо: мясной салат.\nВторое блюдо: сорбет.\nБанкет забронирован на ", "sorbetFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, мясным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuFeeMeat") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет с пробковым сбором.\nПервое блюдо: мясной салат.\nВторое блюдо: тирамису.\nБанкет забронирован на ", "tiramisuFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет с пробковым сбором, мясным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetWithoutFeeMeat") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет без пробкового сбора.\nПервое блюдо: мясной салат.\nВторое блюдо: сорбет.\nБанкет забронирован на ", "sorbetWithoutFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, мясным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetWithoutFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuWithoutFeeMeat") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет без пробкового сбора.\nПервое блюдо: мясной салат.\nВторое блюдо: тирамису.\nБанкет забронирован на ", "tiramisuWithoutFeeMeat")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, мясным салатом и тирамису на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuWithoutFeeMeat")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "sorbetWithoutFeeVegetable") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет без пробкового сбора.\nПервое блюдо: овощной салат.\nВторое блюдо: сорбет.\nБанкет забронирован на ", "sorbetWithoutFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "sorbetWithoutFeeVegetable")[0]), keyboards.MenuKeyboard)
	}

	if strings.Contains(update.CallbackData(), "tiramisuWithoutFeeVegetable") {
		userDatabaseAdapter.UpdateHistory(update.CallbackQuery, "banquetHistory",
			"Банкет без пробкового сбора.\nПервое блюдо: овощной салат.\nВторое блюдо: тирамису.\nБанкет забронирован на ", "tiramisuWithoutFeeVegetable")
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Вы выбрали банкет без пробкового сбора, овощным салатом и сорбетом на %s."+
				"\nВ скором времени с Вами свяжется наш менеджер для подтверждения бронирования банкета!",
				strings.Split(update.CallbackQuery.Data, "tiramisuWithoutFeeVegetable")[0]), keyboards.MenuKeyboard)
	}
}
