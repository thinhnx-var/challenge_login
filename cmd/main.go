package main

import (
	"github.com/thinhnx-var/challenge_login/service/router"
)

func main() {
	host := "127.0.0.1:8088"
	router.StartH2CServer(host)
}
