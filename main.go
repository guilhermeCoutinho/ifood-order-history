package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const inputSize int = 2000
const wrongUsageMessage = "Usage: make run STARTING_DATE='dd-mm-yyyy' END_DATE='dd-mm-yyyy' "
const timeFormat = "02-01-2006" // dd-mm-yyyy

func getRequests() []string {
	req, ok := os.LookupEnv("CURL_REQUEST")
	if !ok {
		panic("Envar not set")
	}
	return []string{string(req)}
}

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

	filter := func(o *Order) bool {
		return o.LastStatus == "CONCLUDED" &&
			o.CreatedAt.After(startingDate) &&
			o.CreatedAt.Before(endDate)
	}
	orders := getOrders(filter)
	printTabulated(orders, "\t")
}

func getOrders(filter func(*Order) bool) []*Order {
	allOrders := []*Order{}
	for _, request := range getRequests() {
		req := strings.ReplaceAll(request, "\n", "")
		req = strings.ReplaceAll(req, "\\", "")
		response := curlToReq(req)

		orders := make([]*Order, inputSize)
		err := json.Unmarshal(response, &orders)
		if err != nil {
			panic(err.Error())
		}
		allOrders = append(allOrders, orders...)
	}

	filteredOrders := make([]*Order, 0)
	for _, order := range allOrders {
		if filter(order) {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func curlToReq(curl string) []byte {
	args := strings.Split(curl, "-H")
	url := strings.Split(args[0], "'")[1]
	url = strings.Replace(url, "size=12", fmt.Sprintf("size=%d", inputSize), 1)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	for i := 1; i < len(args); i++ {
		args[i] = strings.ReplaceAll(args[i], "'", "")
		args[i] = strings.Trim(args[i], " ")
		args := strings.Split(args[i], ":")
		arg := args[0]
		param := args[1]
		req.Header.Set(arg, param)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(io.LimitReader(resp.Body, int64(1024*10*inputSize)))
	if err != nil {
		panic(err.Error())
	}
	return response
}

func printTabulated(orders []*Order, separator string) {
	for _, order := range orders {
		year, month, day := order.CreatedAt.Date()
		name := order.Merchant.Name
		price := fmt.Sprintf("%.2f", float64(order.Payment.Values.Bag)/float64(100))

		orderSummary := ""
		for _, item := range order.Bag.Items {
			orderSummary += strings.ReplaceAll(item.Name, separator, " ") + ";"
		}

		output := []string{
			fmt.Sprintf("%v-%v-%v", day, month, year),
			name,
			string(price),
			orderSummary,
		}
		fmt.Println(strings.Join(output, separator))
	}
}

func print(v interface{}) {
	s, _ := json.MarshalIndent(v, "", " ")
	fmt.Println(string(s))
}
