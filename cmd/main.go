package main

import (
	"time"

	"github.com/VandiKond/parse-ru-time-duration-go/internal/application"
)

func main() {
	app := application.New(time.Second * 10)
	app.Run()
}
