package pkg

import "flag"

func GetFlags() (string, string, int, string, string, string) {
	flagToken := flag.String("token", "", "a bot token")
	flagIp := flag.String("ip", "", "an IP of clickHouse")
	flagPort := flag.Int("port", 9440, "a port of clickHouse")
	flagPassword := flag.String("password", "", "a password of clickHouse")
	flagUser := flag.String("user", "", "a name of user of clickHouse")
	flagDatabase := flag.String("database", "projects", "a name of database of clickHouse")

	flag.Parse()
	return *flagToken, *flagIp, *flagPort, *flagPassword, *flagUser, *flagDatabase
}
