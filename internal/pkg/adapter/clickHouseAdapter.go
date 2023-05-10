package adapter

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"log"
	"restaurantBot/internal/pkg/clickHouse"
)

type Client struct {
	database *clickHouse.DBConnect
}

type GetMessage struct {
	ID      string `json:"id" db:"id"`
	Message string `json:"message" db:"message"`
}

func CreateAdapter(database *clickHouse.DBConnect) *Client {
	adapter := &Client{database: database}
	return adapter
}

func (adapterDb *Client) Get() []*GetMessage {
	ctx := context.Background()
	rows, err := adapterDb.database.Connection.Query(ctx, "SELECT * FROM projects.history")
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
