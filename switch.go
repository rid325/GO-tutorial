package main

import "fmt"

func getDayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "Weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	default:
		return "Unknown"
	}
}

func showSwitch() {
	days := []string{"Monday", "Saturday", "Wednesday", "Sunday"}
	for _, day := range days {
		fmt.Printf("%s is a %s\n", day, getDayType(day))
	}
}
