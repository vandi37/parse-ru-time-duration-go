package main

import (
	"time"

	"github.com/VandiKond/parse-ru-time-duration-go/internal/application"
)

func main() {
	app := application.New(time.Hour)
	app.Run()
}
