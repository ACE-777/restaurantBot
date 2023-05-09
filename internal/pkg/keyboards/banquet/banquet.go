package banquet

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Готовый", "ready"),
		tgbotapi.NewInlineKeyboardButtonData("Персональный", "personalize"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)

var PersonalizeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)

var ReadyKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Со сбором", "withFee"),
		tgbotapi.NewInlineKeyboardButtonData("Без сбора", "withoutFee"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)

var FeeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Овощной салат", "VegetableFee"),
		tgbotapi.NewInlineKeyboardButtonData("Мясной салат", "MeatFee"),
	),
)

var FeeKeyboardTwoVegetable = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сорбет", "sorbetFeeVegetable"),
		tgbotapi.NewInlineKeyboardButtonData("тирамису", "tiramisuFeeVegetable"),
	),
)

var FeeKeyboardTwoMeat = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сорбет", "sorbetFeeMeat"),
		tgbotapi.NewInlineKeyboardButtonData("тирамису", "tiramisuFeeMeat"),
	),
)

var WithoutFeeKeyboardTwoVegetable = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сорбет", "sorbetWithoutFeeVegetable"),
		tgbotapi.NewInlineKeyboardButtonData("тирамису", "tiramisuWithoutFeeVegetable"),
	),
)

var WithoutFeeKeyboardTwoMeat = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("сорбет", "sorbetWithoutFeeMeat"),
		tgbotapi.NewInlineKeyboardButtonData("тирамису", "tiramisuWithoutFeeMeat"),
	),
)

var WithoutFeeKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Овощной салат", "VegetableWithoutFee"),
		tgbotapi.NewInlineKeyboardButtonData("Мясной салат", "MeatWithoutFee"),
	),
)

var MenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Со сбором", "withFee"),
		tgbotapi.NewInlineKeyboardButtonData("Без сбора", "withoutFee"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Главная", "menu"),
	),
)
