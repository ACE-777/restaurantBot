package adapter

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"restaurantBot/internal/pkg/clickHouse"
	"strings"
	"sync"
)

type Client struct {
	database *clickHouse.DBConnect
	mu       sync.Mutex
}

type GetMessage struct {
	ID      string `json:"id" db:"id"`
	Message string `json:"message" db:"message"`
}

type GetBooking struct {
	Time string `json:"time_of_booking" db:"time_of_booking"`
}

type GetBanquet struct {
	Time string `json:"time_of_banquet" db:"time_of_banquet"`
}

type GetHistory struct {
	History string `json:"data_of_booking" db:"data_of_booking"`
}

func CreateAdapter(database *clickHouse.DBConnect) *Client {
	adapter := &Client{database: database}
	return adapter
}

func (adapterDb *Client) Get() []*GetMessage {
	ctx := context.Background()
	rows, err := adapterDb.database.Connection.Query(ctx, "SELECT FROM projects.history")
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows driver.Rows) { _ = rows.Close() }(rows)

	gets := make([]*GetMessage, 0)

	for rows.Next() {
		get := &GetMessage{}
		if err = rows.Scan(&get.ID, &get.Message); err != nil {
			log.Fatal(err)
		}

		gets = append(gets, get)
	}

	return gets
}

func (adapterDb *Client) GetBooking(columnData string) []*GetBooking {
	ctx := context.Background()
	var query string

	if columnData == "skipFirst" {
		query = fmt.Sprintf("SELECT time_of_booking FROM projects.booking WHERE (bar_booking == 1) OR (table_booking == 1)")
	} else if columnData == "skipSecond" {
		query = fmt.Sprintf("SELECT time_of_booking FROM projects.booking WHERE (fireplace_booking == 1) OR (scene_booking == 1)")
	} else {
		query = fmt.Sprintf("SELECT time_of_booking FROM projects.booking WHERE (%s == 1)", columnData+"_booking")
	}

	rows, err := adapterDb.database.Connection.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows driver.Rows) { _ = rows.Close() }(rows)

	gets := make([]*GetBooking, 0)

	for rows.Next() {
		get := &GetBooking{}
		if err = rows.Scan(&get.Time); err != nil {
			log.Fatal(err)
		}

		gets = append(gets, get)
	}

	return gets
}

func (adapterDb *Client) GetBanquet() []*GetBanquet {
	ctx := context.Background()
	var query string

	query = fmt.Sprintf("SELECT time_of_banquet FROM projects.banquet")

	rows, err := adapterDb.database.Connection.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows driver.Rows) { _ = rows.Close() }(rows)

	gets := make([]*GetBanquet, 0)

	for rows.Next() {
		get := &GetBanquet{}
		if err = rows.Scan(&get.Time); err != nil {
			log.Fatal(err)
		}

		gets = append(gets, get)
	}

	return gets
}

func (adapterDb *Client) CheckBanquet() (flagAvailable bool) {
	ctx := context.Background()
	var query string

	query = fmt.Sprintf("SELECT time_of_banquet FROM projects.banquet")

	rows, err := adapterDb.database.Connection.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows driver.Rows) { _ = rows.Close() }(rows)

	gets := make([]*GetBanquet, 0)

	for rows.Next() {
		get := &GetBanquet{}
		if err = rows.Scan(&get.Time); err != nil {
			log.Fatal(err)
		}

		gets = append(gets, get)
	}

	if len(gets) == 0 {
		flagAvailable = true
	}

	return
}

func (adapterDb *Client) UpdateHistory(CallbackQuery *tgbotapi.CallbackQuery, typeBooking, dataOfBooking, time string) {
	ctx := context.Background()
	var query string

	query = fmt.Sprintf("INSERT INTO projects.history VALUES (%v,'%s','%s')", CallbackQuery.Message.Chat.ID,
		typeBooking, dataOfBooking+strings.Split(CallbackQuery.Data, time)[0])

	err := adapterDb.database.Connection.Exec(ctx, query)
	if err != nil {
		log.Printf("Error in executing query to clickHouse:%v", err)
	}

	adapterDb.mu.Lock()
	defer adapterDb.mu.Unlock()
	if typeBooking == "banquetHistory" {
		query = fmt.Sprintf("ALTER TABLE projects.banquet DELETE WHERE time_of_banquet = '%s'", strings.Split(CallbackQuery.Data, time)[0])

		err = adapterDb.database.Connection.Exec(ctx, query)
		if err != nil {
			log.Printf("Error in deleting row from clickHouse: %v", err)
		}
	}
}

func (adapterDb *Client) GetHistory(CallbackQuery *tgbotapi.CallbackQuery, typeOfBooking string) []*GetHistory {
	ctx := context.Background()
	var query string

	query = fmt.Sprintf("SELECT (data_of_booking) FROM projects.history WHERE (user_id = %v) AND (type_of_booking = '%s')",
		CallbackQuery.Message.Chat.ID, typeOfBooking)

	rows, err := adapterDb.database.Connection.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows driver.Rows) { _ = rows.Close() }(rows)

	gets := make([]*GetHistory, 0)

	for rows.Next() {
		get := &GetHistory{}
		if err = rows.Scan(&get.History); err != nil {
			log.Fatal(err)
		}

		gets = append(gets, get)
	}

	return gets
}
