package parse_test

import (
	"testing"
	"time"

	"github.com/VandiKond/parse-ru-time-duration-go/pkg/parse"
)

func TestParse(t *testing.T) {
	var testCases = []struct {
		input  string
		output time.Duration
	}{
		{
			input:  "1 мин",
			output: time.Minute,
		},
		{
			input:  "1 час 6 минут",
			output: time.Hour + (time.Minute * 6),
		},
		{
			input:  "2 часа 15 минут 30 секунд",
			output: (time.Hour * 2) + (time.Minute * 15) + (time.Second * 30),
		},
		{
			input:  "5 минут 10 секунд",
			output: (time.Minute * 5) + (time.Second * 10),
		},
		{
			input:  "1 час",
			output: time.Hour,
		},
		{
			input:  "30 секунд",
			output: time.Second * 30,
		},
		{
			input:  "1 день",
			output: time.Hour * 24,
		},
		{
			input:  "2 дня 3 часа",
			output: (time.Hour * 24 * 2) + (time.Hour * 3),
		},
		{
			input:  "1 неделя",
			output: time.Hour * 24 * 7,
		},
		{
			input:  "2 недели 1 день 5 часов",
			output: (time.Hour * 24 * 7 * 2) + (time.Hour * 24) + (time.Hour * 5),
		},
		{
			input:  "10 сек",
			output: time.Second * 10,
		},
		{
			input:  "1 м",
			output: time.Minute,
		},
		{
			input:  "2 ч",
			output: time.Hour * 2,
		},
		{
			input:  "3 д",
			output: time.Hour * 24 * 3,
		},
		{
			input:  "1 н",
			output: time.Hour * 24 * 7,
		},
		{
			input:  "0 мин",
			output: time.Duration(0),
		},
	}

	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			var got, err = parse.Parser(c.input)
			if err != nil {
				t.Fatalf("got error %s expected no errors", err.Error())
			}
			if got != c.output {
				t.Fatalf("got %s, expected %s", got, c.output)
			}
		})
	}
}
