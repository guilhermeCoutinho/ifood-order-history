package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func printTabulated(orders []*Order, separator string) {
	for _, order := range orders {
		date := order.CreatedAt.Format("02/Jan/06 15:04")

		name := order.Merchant.Name
		price := fmt.Sprintf("%.2f", float64(order.Payments.Total.Value)/float64(100))

		orderSummary := ""
		for _, item := range order.Bag.Items {
			orderSummary += strings.ReplaceAll(item.Name, separator, " ") + ";"
		}

		output := []string{
			date,
			name,
			string(price),
			orderSummary,
		}
		fmt.Println(strings.Join(output, separator))
	}
}

func printProgress(current, total int) {
	barSegmentCount := 20
	currentNormalized := int(float64(barSegmentCount) * float64(current) / float64(total))

	progress := ""
	for i := 0; i < barSegmentCount; i++ {
		if i < currentNormalized {
			progress += "="
		} else {
			progress += "."
		}
	}

	fmt.Printf("\r%d/%d [%s]", current, total, progress)
}

func print(v interface{}) {
	//	s, _ := json.MarshalIndent(v, "", " ")
	s, _ := json.Marshal(v)
	fmt.Println(string(s))
}
