package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)

func main() {
	startingDate, _ := time.Parse(timeFormat, "01-01-0001")
	endDate, _ := time.Parse(timeFormat, "01-01-3001")

	if len(os.Args) == 3 {
		var err error
		startingDate, err = time.Parse(timeFormat, os.Args[1])
		if err != nil {
			panic(wrongUsageMessage)
		}
		endDate, err = time.Parse(timeFormat, os.Args[2])
		if err != nil {
			panic(wrongUsageMessage)
		}
	}

	orders := getOrders()

	filter := func(o *Order) bool {
		return o.LastStatus == "CONCLUDED" &&
			o.CreatedAt.After(startingDate) &&
			o.CreatedAt.Before(endDate)
	}

	filteredOrders := filterOrders(orders, filter)
	printTabulated(filteredOrders, "\t")
}

func getOrders() []*Order {
	totalOrders := make([]*Order, 0)
	numberOfPages := int(math.Ceil(float64(historyLen) / float64(maxItensPerPage)))

	printProgress(0, numberOfPages)

	for i := 0; i < numberOfPages; i++ {
		itensPerPage := getRequestPaginationParam(i, numberOfPages)
		response := requestOrderHistory(i, itensPerPage)

		orders := make([]*Order, itensPerPage)
		err := json.Unmarshal(response, &orders)
		if err != nil {
			panic(err.Error())
		}

		printProgress(i+1, numberOfPages)
		hasNoMorePagesOfHistory := len(orders) == 0
		if hasNoMorePagesOfHistory {
			break
		}

		totalOrders = append(totalOrders, orders...)
	}

	fmt.Println()
	return totalOrders
}

func getRequestPaginationParam(currentPageIndex, numberOfPages int) int {
	itensPerPage := maxItensPerPage
	if isLastPage := currentPageIndex == numberOfPages-1; isLastPage {
		itensPerPage = historyLen - ((numberOfPages - 1) * maxItensPerPage)
	}
	return itensPerPage
}

func requestOrderHistory(page, size int) []byte {
	url := fmt.Sprintf("https://marketplace.ifood.com.br/v4/customers/me/orders?page=%d&size=%d", page, size)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	bufferSize := int64(1024 * 10 * size)
	response, err := ioutil.ReadAll(io.LimitReader(resp.Body, bufferSize))
	if err != nil {
		panic(err.Error())
	}

	return response
}
