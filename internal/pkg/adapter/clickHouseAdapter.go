package adapter

import (
	"context"
	"fmt"
	"log"
	"restaurantBot/internal/pkg/clickHouse"
	"time"
)

type Client struct {
	database *clickHouse.DBConnect
}

func CreateAdapter(database *clickHouse.DBConnect) *Client {
	adapter := &Client{database: database}
	return adapter
}

func (adapterDb *Client) Get() {
	ctx := context.Background()
	rows, err := adapterDb.database.Connection.Query(ctx, "SELECT countries, mountain, sea, excursion, health FROM tour.countries")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var version string
		var version2 string
		var version3 string
		var version4 time.Time
		var version5 int64
		if err := rows.Scan(&version, &version2, &version3, &version4, &version5); err != nil {
			log.Fatal(err)
		}
		fmt.Println(version, version2, version3, version4, version5)
		fmt.Println("====================================")
	}

	return
}
