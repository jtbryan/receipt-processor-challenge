package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	model "receiptProcessor/Models"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	invalidReceiptError = "The receipt is invalid."
	timeFormat          = "15:04"
	dateFormat          = "2006-01-02"
)

func ValidateReceipt(c *gin.Context) (model.Receipt, error) {
	body := c.Request.Body
	var receipt model.Receipt

	if body == nil {
		return receipt, errors.New(invalidReceiptError)
	}

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&receipt); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}

	validate := validator.New()
	if err := validate.Struct(receipt); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}
	if err := validateField(receipt.Retailer, `^[\w\s\-\&]+$`); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}
	if _, err := time.Parse(dateFormat, receipt.PurchaseDate); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}
	if _, err := time.Parse(timeFormat, receipt.PurchaseTime); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}
	if err := validateField(receipt.Total, `^\d+\.\d{2}$`); err != nil {
		return receipt, errors.New(invalidReceiptError)
	}

	for _, item := range receipt.Items {
		if err := validateField(item.Price, `^\d+\.\d{2}$`); err != nil {
			return receipt, errors.New(invalidReceiptError)
		}
		if err := validateField(item.ShortDescription, `^[\w\s\-]+$`); err != nil {
			return receipt, errors.New(invalidReceiptError)
		}
	}

	return receipt, nil
}

func validateField(field string, pattern string) error {
	match, err := regexp.MatchString(pattern, field)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if !match {
		return errors.New(invalidReceiptError)
	}
	return nil
}
