package server_http

import (
	"io"
	"net/http"

	"github.com/VandiKond/vanerrors"
	"github.com/vandi37/parse-ru-time-duration-go/pkg/parse"
)

// The url
type Url string

func (u Url) GetUrl() Url {
	return u
}

// Starting the server
func (h ParseHandler) Start() error {
	err := http.ListenAndServe(string(h.GetUrl()), h)
	return err
}

// Serving the server
func (h ParseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "CouldNotReadBody", http.StatusInternalServerError)
		return
	}
	result, err := parse.Parser(string(body))
	if err != nil {
		errName := vanerrors.GetName(err)
		if errName == "" {
			errName = "InternalServerError"
		}
		errCode := vanerrors.GetCode(err)
		http.Error(w, errName, errCode)
		return
	}
	w.Write([]byte(result.String()))
}

// The handler struct
type ParseHandler struct {
	Url
}

// The interface for all handlers
type Handler interface {
	// Starts the http service
	Start() error
	// Gets the url
	GetUrl() Url
}
