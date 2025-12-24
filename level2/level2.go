package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func TimeNow() time.Time {
	return TimeNow()
}

func currentDayOfTheWeek() string {
	switch time.Time.Weekday(TimeNow()) {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресенье"
	}
	return ""
}

func dayOrNight() string {
	var time int = TimeNow().Hour()
	if time >= 10 && time <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	var currentDay = time.Time.Weekday(TimeNow())
	return (5 - int(currentDay) + 7) % 7

}

func CheckCurrentDayOfTheWeek(answer string) bool {
	var str string = strings.ToLower(answer)
	switch time.Time.Weekday(TimeNow()) {
	case time.Monday:
		return str == "понедельник"
	case time.Tuesday:
		return str == "вторник"
	case time.Wednesday:
		return str == "среда"
	case time.Thursday:
		return str == "четверг"
	case time.Friday:
		return str == "пятница"
	case time.Saturday:
		return str == "суббота"
	case time.Sunday:
		return str == "воскресенье"
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	var str string = strings.Trim(answer, " ")
	str = strings.ToLower(answer)
	if utf8.RuneCountInString(str) != 4 || (str != "день" && str != "ночь") {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}

	var time int = TimeNow().Hour()
	if str == "день" {
		if time >= 10 && time <= 22 {
			return true, nil
		}
		return false, nil
	} else {
		if time < 10 || time > 22 {
			return true, nil
		}
		return false, nil
	}

}
