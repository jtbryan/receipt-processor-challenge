package points

import (
	models "receiptProcessor/Models"

	"math"
	"strconv"
	"strings"
	"unicode"
)

func CalculatePoints(receipt models.Receipt) string {
	totalPoints := 0

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			totalPoints += 1
		}
	}

	change, err := strconv.Atoi(receipt.Total[len(receipt.Total)-2:])

	if err != nil {
		panic(err)
	}

	// 50 points if the total is a round dollar amount with no cents.
	if change == 0 {
		totalPoints += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if change%25 == 0 {
		totalPoints += 25
	}

	// 5 points for every two items on the receipt.
	totalPoints += ((len(receipt.Items) / 2) * 5)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)

			if err != nil {
				panic(err)
			}

			totalPoints += int(math.Ceil(price * 0.2))
		}
	}

	day, err := strconv.Atoi(receipt.PurchaseDate[len(receipt.PurchaseDate)-2:])

	if err != nil {
		panic(err)
	}

	// 6 points if the day in the purchase date is odd.
	if day%2 != 0 {
		totalPoints += 6
	}

	hr, err := strconv.Atoi(receipt.PurchaseTime[:2])

	if err != nil {
		panic(err)
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if hr >= 14 && hr < 16 {
		totalPoints += 10
	}

	return strconv.Itoa(totalPoints)
}
