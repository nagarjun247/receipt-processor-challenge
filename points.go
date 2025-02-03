package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt Receipt) int {
	points := 0

	alnumRegex := regexp.MustCompile("[a-zA-Z0-9]")
	points += len(alnumRegex.FindAllString(receipt.Retailer, -1))

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
	}

	timeParts := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(timeParts[0])
	if hour >= 14 && hour < 16 {
		points += 10
	}

	return points
}
