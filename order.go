package main

import "time"

type Order struct {
	ID             string         `json:"id"`
	ShortID        string         `json:"shortId"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	ClosedAt       time.Time      `json:"closedAt"`
	CreatedAt      time.Time      `json:"createdAt"`
	LastStatus     string         `json:"lastStatus"`
	Details        Details        `json:"details"`
	Delivery       Delivery       `json:"delivery"`
	Merchant       Merchant       `json:"merchant"`
	Payments       Payments       `json:"payments"`
	Bag            Bag            `json:"bag"`
	Origin         Origin         `json:"origin"`
	DeliveryMethod DeliveryMethod `json:"deliveryMethod"`
}
type Details struct {
	Mode             string `json:"mode"`
	Scheduled        bool   `json:"scheduled"`
	Tippable         bool   `json:"tippable"`
	IndoorTipEnabled bool   `json:"indoorTipEnabled"`
	Trackable        bool   `json:"trackable"`
	Boxable          bool   `json:"boxable"`
	PlacedAtBox      bool   `json:"placedAtBox"`
	Reviewed         bool   `json:"reviewed"`
	DarkKitchen      bool   `json:"darkKitchen"`
}
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Driver struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photoUrl"`
	Modal    string `json:"modal"`
}
type EstimatedTimeOfArrival struct {
	DeliversAt string `json:"deliversAt"`
	UpdatedAt  string `json:"updatedAt"`
}
type Delivery struct {
	Address                Address                `json:"address"`
	Driver                 Driver                 `json:"driver"`
	EstimatedTimeOfArrival EstimatedTimeOfArrival `json:"estimatedTimeOfArrival"`
	ExpectedDeliveryTime   string                 `json:"expectedDeliveryTime"`
	ExpectedDuration       int                    `json:"expectedDuration"`
}
type Address struct {
	Establishment string      `json:"establishment"`
	City          string      `json:"city"`
	Country       string      `json:"country"`
	Neighborhood  string      `json:"neighborhood"`
	State         string      `json:"state"`
	StreetName    string      `json:"streetName"`
	StreetNumber  string      `json:"streetNumber"`
	Coordinates   Coordinates `json:"coordinates"`
	Reference     string      `json:"reference"`
	Complement    string      `json:"complement"`
}
type Merchant struct {
	Address      Address `json:"address"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	PhoneNumber  string  `json:"phoneNumber"`
	Logo         string  `json:"logo"`
	CompanyGroup string  `json:"companyGroup"`
	Type         string  `json:"type"`
}
type Method struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Type struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Brand struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MealVoucher struct {
	CardNumber string `json:"cardNumber"`
}
type Amount struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}
type Transactions struct {
	ID        string        `json:"id"`
	Type      string        `json:"type"`
	Status    string        `json:"status"`
	CreatedAt string        `json:"createdAt"`
	Value     int           `json:"value"`
	Refunds   []interface{} `json:"refunds"`
}
type Methods struct {
	ID           string         `json:"id"`
	Method       Method         `json:"method"`
	Type         Type           `json:"type"`
	Brand        Brand          `json:"brand"`
	MealVoucher  MealVoucher    `json:"mealVoucher"`
	Amount       Amount         `json:"amount"`
	Transactions []Transactions `json:"transactions"`
}
type Payments struct {
	Methods []Methods `json:"methods"`
	Total   Amount    `json:"total"`
}
type Benefits struct {
	Type        string `json:"type"`
	Target      string `json:"target"`
	TargetID    string `json:"targetId"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}
type DeliveryFee struct {
	Value             int `json:"value"`
	ValueWithDiscount int `json:"valueWithDiscount"`
}
type SubItems struct {
	ID                     string        `json:"id"`
	ExternalID             string        `json:"externalId"`
	Name                   string        `json:"name"`
	Quantity               int           `json:"quantity"`
	Tags                   []interface{} `json:"tags"`
	TotalPrice             int           `json:"totalPrice"`
	TotalPriceWithDiscount int           `json:"totalPriceWithDiscount"`
	UnitPrice              int           `json:"unitPrice"`
	UnitPriceWithDiscount  int           `json:"unitPriceWithDiscount"`
}
type Items struct {
	ID                     string     `json:"id"`
	UniqueID               string     `json:"uniqueId"`
	ExternalID             string     `json:"externalId"`
	Name                   string     `json:"name"`
	Description            string     `json:"description"`
	Quantity               int        `json:"quantity"`
	SubItems               []SubItems `json:"subItems"`
	Tags                   []string   `json:"tags"`
	TotalPrice             int        `json:"totalPrice"`
	TotalPriceWithDiscount int        `json:"totalPriceWithDiscount"`
	UnitPrice              int        `json:"unitPrice"`
	UnitPriceWithDiscount  int        `json:"unitPriceWithDiscount"`
	Notes                  string     `json:"notes"`
}
type SubTotal struct {
	Value             int `json:"value"`
	ValueWithDiscount int `json:"valueWithDiscount"`
}
type Total struct {
	Value             int `json:"value"`
	ValueWithDiscount int `json:"valueWithDiscount"`
}
type Bag struct {
	Benefits    []Benefits  `json:"benefits"`
	DeliveryFee DeliveryFee `json:"deliveryFee"`
	Items       []Items     `json:"items"`
	SubTotal    SubTotal    `json:"subTotal"`
	Total       Total       `json:"total"`
	Updated     bool        `json:"updated"`
}
type Origin struct {
	Platform   string `json:"platform"`
	AppName    string `json:"appName"`
	AppVersion string `json:"appVersion"`
}
type DeliveryMethod struct {
	ID   string `json:"id"`
	Mode string `json:"mode"`
}

func filterOrders(orders []*Order, filter func(*Order) bool) []*Order {
	filteredOrders := make([]*Order, 0)
	for _, order := range orders {
		if filter(order) {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}
