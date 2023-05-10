package clickHouse

import (
	"crypto/tls"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
)

type DBConnect struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`

	Connection clickhouse.Conn
}

func (client *DBConnect) Open() error {
	tlsConfig := &tls.Config{}

	db, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", client.Ip, client.Port)},
		Auth: clickhouse.Auth{
			Database: client.Database,
			Username: client.User,
			Password: client.Password,
		},
		TLS: tlsConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.Connection = db
	return nil
}

func (client *DBConnect) Close() {
	err := client.Connection.Close()
	if err != nil {
		log.Fatal(err)
	}
}
