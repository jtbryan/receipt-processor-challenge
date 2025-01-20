package points

import (
	models "receiptProcessor/Models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	receipt := new(models.Receipt)

	receipt.Retailer = "Target"
	receipt.PurchaseDate = "2022-01-01"
	receipt.PurchaseTime = "15:15"

	items := []models.Item{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
	}
	receipt.Items = items

	receipt.Total = "35.35"

	actual := CalculatePoints(*receipt)
	expected := "38"

	if actual != expected {
		t.Errorf("got %s, wanted %s", actual, expected)
	}
}

func TestReceiptTotalError(t *testing.T) {
	defer func() { recover() }()
	receipt := new(models.Receipt)
	receipt.Total = "35.ab"

	CalculatePoints(*receipt)
}

func TestReceiptItemPriceError(t *testing.T) {
	defer func() { recover() }()

	receipt := new(models.Receipt)
	receipt.Total = "35.00"
	items := []models.Item{
		{ShortDescription: "Emils Cheese Pizza", Price: "6.ab"},
	}
	receipt.Items = items

	CalculatePoints(*receipt)
}

func TestReceiptPurchaseDateError(t *testing.T) {
	defer func() { recover() }()

	receipt := new(models.Receipt)
	receipt.PurchaseDate = "2020-01-ab"
	receipt.Total = "35.00"
	items := []models.Item{
		{ShortDescription: "Emils Cheese Pizza", Price: "6.20"},
	}
	receipt.Items = items

	CalculatePoints(*receipt)
}

func TestReceiptPurchaseTimeError(t *testing.T) {
	defer func() { recover() }()

	receipt := new(models.Receipt)
	receipt.PurchaseDate = "2020-01-01"
	receipt.PurchaseTime = "ab:12"
	receipt.Total = "35.00"
	items := []models.Item{
		{ShortDescription: "Emils Cheese Pizza", Price: "6.20"},
	}
	receipt.Items = items

	CalculatePoints(*receipt)
}
