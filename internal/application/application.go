package application

import (
	"fmt"
	"log"
	"os"
	"time"

	server_http "github.com/vandi37/parse-ru-time-duration-go/internal/http"
)

type Config struct {
	server_http.Handler
}

type Application struct {
	Duration  time.Duration
	IsService bool
	Config    Config
}

func NewService() *Application {
	return &Application{}
}

func New(d time.Duration) *Application {
	return &Application{
		Duration: d,
	}
}

func (a *Application) Run() error {
	// Exiting in duration
	defer log.Printf("application stopped before timeout")
	go a.ExitTimeOut()

	// The program
	log.Println("the program is working")
	a.Config.Start()
	// The program end

	// Returning without error
	return nil
}

func (a *Application) ExitTimeOut() {
	// Checking service mod
	if a.IsService {
		return
	}

	// Waiting duration seconds
	time.Sleep(a.Duration)

	// Exiting after timeout
	fmt.Println("")
	log.Printf("timeout %s has passed. Ending the program", a.Duration)
	os.Exit(418)
}
