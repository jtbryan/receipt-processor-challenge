package main

import (
	"fmt"
	models "receptWebService/Models"
)

func main() {
	var items = []models.Item{
		{
			ShortDescription: "test1",
			Price:            "1.0",
		},
		{
			ShortDescription: "test2",
			Price:            "1.0",
		},
	}

	fmt.Println(items)
}
