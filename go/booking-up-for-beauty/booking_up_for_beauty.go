package booking

import (
	"fmt"
	"log"
	"time"
)

var openDay time.Time = time.Date(2012, 9, 15, 0, 0, 0, 0, time.UTC)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dateFormat := "1/2/2006 15:04:05"
	dateString, error := time.Parse(dateFormat, date)
	if error != nil {
		log.Fatalln(error)
	}
	return dateString
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateString, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatalln(err)
	}
	return dateString.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateString, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatalln(err)
	}
	return dateString.Hour() >= 12 && dateString.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateString := Schedule(date).Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %v.", dateString)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), openDay.Month(), openDay.Day(), 0, 0, 0, 0, time.UTC)
}
