package main

import (
	"strconv"
	"strings"
)

func FormatDuration(seconds int64) string {
	currentTime := seconds
	var sec int64
	var mins int64
	var hours int64
	var days int64
	var years int64

	if currentTime == 0 {
		return "now"
	}

	for currentTime > 0 {
		if currentTime/31536000 > 0 {
			years += seconds / 31536000
			currentTime -= seconds / 31536000 * 31536000
		} else if currentTime/86400 > 0 {
			days += currentTime / 86400
			currentTime -= (currentTime / 86400) * 86400
		} else if currentTime/3600 > 0 {
			hours += currentTime / 3600
			currentTime -= (currentTime / 3600) * 3600
		} else if currentTime/60 > 0 {
			mins += currentTime / 60
			currentTime -= (currentTime / 60) * 60
		} else {
			sec += currentTime
			currentTime = 0
		}
	}

	var listSentence []string

	if years > 1 {
		listSentence = append(listSentence, strconv.FormatInt(years, 10)+" years")
	} else if years == 1 {
		listSentence = append(listSentence, "1 year")
	}

	if days > 1 {
		listSentence = append(listSentence, strconv.FormatInt(days, 10)+" days")
	} else if days == 1 {
		listSentence = append(listSentence, "1 day")
	}

	if hours > 1 {
		listSentence = append(listSentence, strconv.FormatInt(hours, 10)+" hours")
	} else if hours == 1 {
		listSentence = append(listSentence, "1 hour")
	}

	if mins > 1 {
		listSentence = append(listSentence, strconv.FormatInt(mins, 10)+" minutes")
	} else if mins == 1 {
		listSentence = append(listSentence, "1 minute")
	}

	if sec > 1 {
		listSentence = append(listSentence, strconv.FormatInt(sec, 10)+" seconds")
	} else if sec == 1 {
		listSentence = append(listSentence, "1 second")
	}

	newString := strings.Join(listSentence[:len(listSentence)-1], ", ")

	if len(listSentence) == 1 {
		newString = listSentence[0]
	} else {
		newString = newString + " and " + listSentence[len(listSentence)-1]
	}

	return strconv.FormatInt(currentTime, 10)
}
