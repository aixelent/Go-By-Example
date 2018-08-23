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
	days := int(difference.Hours() / 24)
	hours := int(difference.Hours())
	minutes := int(difference.Minutes())
	seconds := int(difference.Seconds())

	// fmt.Printf("Difference in\nDays: %d", days, "\nHours: %d", hours, "\nMinutes", minutes, "\nSeconds", seconds)
	fmt.Print("Difference in:\n")
	fmt.Printf("Days: %d\n", days)
	fmt.Printf("Hours: %d\n", hours)
	fmt.Printf("Minutes: %d\n", minutes)
	fmt.Printf("Seconds: %d\n", seconds)
}
