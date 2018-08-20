package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("\nToday: ", loc, "Time: ", now)

	pastDate := time.Date(2018, time.August, 11, 11, 22, 20, 100, time.UTC)
	fmt.Println("\nToday: ", loc, "Time: ", pastDate)

	difference := now.Sub(pastDate)
	days := difference.Hours() / 24
	hours := difference.Hours()
	minutes := difference.Minutes()
	seconds := difference.Seconds()

	fmt.Println("Difference in\nDays: %d", days, "\nHours: %d", hours, "\nMinutes", minutes, "\nSeconds", seconds)
}
