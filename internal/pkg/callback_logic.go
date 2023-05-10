package pkg

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"restaurantBot/internal/pkg/adapter"
	"restaurantBot/internal/pkg/clickHouse"
	"restaurantBot/internal/pkg/keyboards"
	"restaurantBot/internal/pkg/keyboards/banquet"
	"strconv"
)

func makeKeyboardBooking(db *clickHouse.DBConnect, update tgbotapi.Update) (keyboard tgbotapi.InlineKeyboardMarkup) {
	userDatabaseAdapter := adapter.CreateAdapter(db)
	JSONData := userDatabaseAdapter.GetBooking(update.CallbackQuery.Data)

	usersJSON, err := json.Marshal(JSONData)
	if err != nil {
		log.Printf("Error in marshalalling func GetBooking: %v", err)
	}

	var bookingData []adapter.GetBooking

	err = json.Unmarshal(usersJSON, &bookingData)
	if err != nil {
		log.Printf("Error in unmarshalling JSON: %v", err)
	}

	keyboard = tgbotapi.NewInlineKeyboardMarkup()
	for i := range bookingData {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(bookingData[i].Time, bookingData[i].Time+update.CallbackQuery.Data+"Booking"),
		)

		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	return
}

func makeKeyboardBanquet(db *clickHouse.DBConnect, update tgbotapi.Update) (keyboard tgbotapi.InlineKeyboardMarkup) {
	userDatabaseAdapter := adapter.CreateAdapter(db)
	JSONData := userDatabaseAdapter.GetBanquet()

	usersJSON, err := json.Marshal(JSONData)
	if err != nil {
		log.Printf("Error in marshalalling func GetBooking: %v", err)
	}

	var bookingData []adapter.GetBanquet

	err = json.Unmarshal(usersJSON, &bookingData)
	if err != nil {
		log.Printf("Error in unmarshalling JSON: %v", err)
	}

	keyboard = tgbotapi.NewInlineKeyboardMarkup()
	for i := range bookingData {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(bookingData[i].Time, bookingData[i].Time+update.CallbackQuery.Data),
		)

		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}

	return
}

func checkAvailabilityBanquet(update tgbotapi.Update, bot *tgbotapi.BotAPI, db *clickHouse.DBConnect) {
	userDatabaseAdapter := adapter.CreateAdapter(db)
	flagAvailable := userDatabaseAdapter.CheckBanquet()
	if flagAvailable {
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"К сожалению на данную неделю больше нет возможности забронировать банкет", keyboards.MenuKeyboard)
	} else {
		sendMessageWithKeyboard(bot, update.CallbackQuery.Message.Chat.ID,
			"Выберите тип пакета:", banquet.MainKeyboard)
	}

}

func GetUsersHistory(userDatabaseAdapter *adapter.Client, update tgbotapi.Update, typeOfBooking string) string {
	arrayOfBooking := userDatabaseAdapter.GetHistory(update.CallbackQuery, typeOfBooking)
	if len(arrayOfBooking) == 0 {
		return "У Вас пока что нет еще бронирований."
	}

	usersJSON, err := json.Marshal(arrayOfBooking)
	if err != nil {
		log.Printf("Error in marshalalling func GetALLTours: %v", err)
	}

	var bookingData []adapter.GetHistory

	err = json.Unmarshal(usersJSON, &bookingData)
	if err != nil {
		log.Printf("Error in unmarshalling JSON: %v", err)
	}

	var output string
	for i := range bookingData {
		output += strconv.Itoa(i+1) + "." + bookingData[i].History + "\n"
	}

	return output
}
