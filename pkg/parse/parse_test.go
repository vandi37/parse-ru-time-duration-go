package parse_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/vandi37/parse-ru-time-duration-go/pkg/parse"
	"github.com/vandi37/vanerrors"
)

func TestParse(t *testing.T) {
	var testCases = []struct {
		input    string
		output   time.Duration
		hasError bool
		httpCode int
		errName  string
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
			input:  "2 недели 1 день",
			output: (time.Hour * 24 * 7 * 2) + (time.Hour * 24 * 1),
		},
		{
			input:  "1 месяц",
			output: time.Hour * 24 * 30,
		},
		{
			input:  "2 года",
			output: time.Hour * 24 * 365 * 2,
		},
		{
			input:  "1 год 2 месяца 3 дня",
			output: (time.Hour * 24 * 365) + (time.Hour * 24 * 30 * 2) + (time.Hour * 24 * 3),
		},
		{
			input:    "123",
			hasError: true,
			httpCode: http.StatusBadRequest,
			errName:  parse.NoTypeAfterNumber,
		},
		{
			input:    "123 abc",
			hasError: true,
			httpCode: http.StatusBadRequest,
			errName:  parse.UnknownWord,
		},
		{
			input:    "1 мин 2",
			hasError: true,
			httpCode: http.StatusBadRequest,
			errName:  parse.NoTypeAfterNumber,
		},
		{
			input:    "abc",
			hasError: true,
			httpCode: http.StatusBadRequest,
			errName:  parse.UnknownWord,
		},
		{
			input:    "1 мин abc",
			hasError: true,
			httpCode: http.StatusBadRequest,
			errName:  parse.UnknownWord,
		},
		{
			input:  "1 с 2 м 3 ч 4 д 5 н 6 мес 7 г",
			output: time.Second + (time.Minute * 2) + (time.Hour * 3) + (time.Hour * 24 * 4) + (time.Hour * 24 * 7 * 5) + (time.Hour * 24 * 30 * 6) + (time.Hour * 24 * 365 * 7),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			duration, err := parse.Parser(tc.input)
			if tc.hasError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
				if vanerrors.GetCode(err) != tc.httpCode {
					t.Errorf("Expected HTTP status code %d, but got %d", tc.httpCode, vanerrors.GetCode(err))
				}
				if vanerrors.GetName(err) != tc.errName {
					t.Errorf("Expected error message '%s', but got '%s'", tc.errName, vanerrors.GetName(err))
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if duration != tc.output {
					t.Errorf("Expected duration %v, but got %v", tc.output, duration)
				}
			}
		})
	}
}
