package parse

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/VandiKond/vanerrors"
)

// Error constants
const (
	NoTypeAfterNumber = "no type after number"
	UnknownWord       = "unknown word"
)

// Parses russian expression to time.Duration type
func Parser(data string) (time.Duration, error) {
	// Setting the result
	var result time.Duration

	// Setting helper variables
	var currentNum time.Duration
	var lastWasNum bool

	// Creating a slice of words
	var parts = strings.Split(data, " ")

	// Going in every word
	for _, p := range parts {
		// If the slice part is a number
		var num, err = strconv.Atoi(p)
		if err == nil {
			// Two numbers in a row not allowed
			if lastWasNum {
				return 0, vanerrors.NewHTTP(NoTypeAfterNumber, http.StatusBadRequest, nil)
			}
			// Setting helper variables
			currentNum = time.Duration(num)
			lastWasNum = true

			continue
		}
		// Two non numbers in a row not allowed
		if !lastWasNum {
			currentNum = 1
		}
		// Finding the type
		var durType = FindDurationType(p)

		// If the type is not valid
		if durType < 0 || durType > 7 {
			return 0, vanerrors.NewHTTP(UnknownWord, 400, nil)
		}
		// Adding the type and current num to the result
		result += Types[durType] * currentNum

		// Setting handler variable
		lastWasNum = false
	}
	if lastWasNum {
		return 0, vanerrors.NewHTTP(NoTypeAfterNumber, http.StatusBadRequest, nil)
	}

	// Returning the result
	return result, nil
}

// Finds out is the string matches with russian time words
// -1 - not found
// 0 - second
// 1 - minute
// 2 - hour
// 3 - day
// 4 - week
// 5 - month
// 6 - year
//
// Your can use parse.Types[] for getting the time.Duration data
func FindDurationType(str string) int {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	switch str {
	case "с", "сек", "секунда", "секунду", "секундой", "секунде", "секунд", "секундам", "секунды", "секундами", "секундах":
		return 0
	case "м", "мин", "минута", "минуту", "минутой", "минуте", "минут", "минутам", "минуты", "минутами", "минутах":
		return 1
	case "ч", "час", "часу", "часом", "часе", "часа", "часов", "часам", "часы", "часами", "часах":
		return 2
	case "д", "дня", "дню", "день", "днём", "дне", "дней", "дням", "дни", "днями", "днях":
		return 3
	case "н", "нед", "неделя", "недели", "неделе", "неделю", "неделей", "недель", "неделям", "неделями", "неделях":
		return 4
	case "мес", "месяц", "месяца", "месяцу", "месяцем", "месяце", "месяцы", "месяцев", "месяцам", "месяцами", "месяцах":
		return 5
	case "г", "л", "год", "года", "году", "годом", "лет", "годам", "годы", "годами", "годах":
		return 6
	default:
		return -1
	}
}

// The duration types
var Types [7]time.Duration = [7]time.Duration{
	time.Second,          // second (0)
	time.Minute,          // minute (1)
	time.Hour,            // hour (2)
	time.Hour * 24,       // day (3)
	time.Hour * 24 * 7,   // week (4)
	time.Hour * 24 * 30,  // month (5)
	time.Hour * 24 * 365, // year (6)
}
