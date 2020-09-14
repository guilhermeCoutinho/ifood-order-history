package main

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID             string          `json:"id"`
	ShortID        string          `json:"shortId"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	ClosedAt       time.Time       `json:"closedAt"`
	LastStatus     string          `json:"lastStatus"`
	Details        *Details        `json:"details"`
	Delivery       *Delivery       `json:"delivery"`
	Merchant       *Merchant       `json:"merchant"`
	Payment        *Payment        `json:"payment"`
	Payments       *Payments       `json:"payments"`
	Bag            *Bag            `json:"bag"`
	Origin         *Origin         `json:"origin"`
	DeliveryMethod *DeliveryMethod `json:"deliveryMethod"`
	Customer       *Customer       `json:"customer"`
}
type Details struct {
	Mode        string `json:"mode"`
	Scheduled   bool   `json:"scheduled"`
	Tippable    bool   `json:"tippable"`
	Trackable   bool   `json:"trackable"`
	Boxable     bool   `json:"boxable"`
	PlacedAtBox bool   `json:"placedAtBox"`
	Reviewed    bool   `json:"reviewed"`
	DarkKitchen bool   `json:"darkKitchen"`
	Cancelable  bool   `json:"cancelable"`
}
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Address struct {
	City         string       `json:"city"`
	Country      string       `json:"country"`
	Neighborhood string       `json:"neighborhood"`
	State        string       `json:"state"`
	StreetName   string       `json:"streetName"`
	PostalCode   string       `json:"postalCode"`
	StreetNumber string       `json:"streetNumber"`
	Coordinates  *Coordinates `json:"coordinates"`
	Reference    string       `json:"reference"`
	Complement   string       `json:"complement"`
}
type Driver struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photoUrl"`
	Modal    string `json:"modal"`
}
type EstimatedTimeOfArrival struct {
	DeliversAt time.Time `json:"deliversAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
type Delivery struct {
	Address                *Address                `json:"address"`
	Driver                 *Driver                 `json:"driver"`
	EstimatedTimeOfArrival *EstimatedTimeOfArrival `json:"estimatedTimeOfArrival"`
	ExpectedDeliveryTime   string                  `json:"expectedDeliveryTime"`
	ExpectedDuration       int                     `json:"expectedDuration"`
}

type Merchant struct {
	Address      *Address  `json:"address"`
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phoneNumber"`
	Logo         string    `json:"logo"`
	CompanyGroup string    `json:"companyGroup"`
	Type         string    `json:"type"`
}
type Values struct {
	Bag int `json:"bag"`
}
type Authorizations struct {
	ID        string        `json:"id"`
	Type      string        `json:"type"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"createdAt"`
	Value     int           `json:"value"`
	Refunds   []interface{} `json:"refunds"`
}
type Payment struct {
	Mode                 string            `json:"mode"`
	Currency             string            `json:"currency"`
	Description          string            `json:"description"`
	MethodImagePath      string            `json:"methodImagePath"`
	CreditCardNumber     string            `json:"creditCardNumber"`
	CreditCardHolderName string            `json:"creditCardHolderName"`
	Values               *Values           `json:"values"`
	Authorizations       []*Authorizations `json:"authorizations"`
}
type TotalValue struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}
type Amount struct {
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}
type Credit struct {
	Brand      string `json:"brand"`
	Gateway    string `json:"gateway"`
	Acquirer   string `json:"acquirer"`
	CardNumber string `json:"cardNumber"`
}
type Methods struct {
	ID        string  `json:"id"`
	Method    string  `json:"method"`
	Type      string  `json:"type"`
	Liability string  `json:"liability"`
	Amount    *Amount `json:"amount"`
	Credit    *Credit `json:"credit"`
}
type Payments struct {
	Total          *TotalValue       `json:"total"`
	Methods        []*Methods        `json:"methods"`
	Authorizations []*Authorizations `json:"authorizations"`
}
type DeliveryFee struct {
	Value             int `json:"value"`
	ValueWithDiscount int `json:"valueWithDiscount"`
}
type Items struct {
	ExternalID             string   `json:"externalId"`
	Name                   string   `json:"name"`
	Quantity               int      `json:"quantity"`
	SubItems               []*Items `json:"subItems"`
	Tags                   []string `json:"tags"`
	TotalPrice             int      `json:"totalPrice"`
	TotalPriceWithDiscount int      `json:"totalPriceWithDiscount"`
	UnitPrice              int      `json:"unitPrice"`
	UnitPriceWithDiscount  int      `json:"unitPriceWithDiscount"`
}

type TotalValueWithDiscount struct {
	Value             int `json:"value"`
	ValueWithDiscount int `json:"valueWithDiscount"`
}
type Bag struct {
	Benefits    []interface{}           `json:"benefits"`
	DeliveryFee DeliveryFee             `json:"deliveryFee"`
	Items       []Items                 `json:"items"`
	SubTotal    *TotalValueWithDiscount `json:"subTotal"`
	Total       *TotalValueWithDiscount `json:"total"`
	Updated     bool                    `json:"updated"`
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
type Customer struct {
	ID string `json:"id"`
}
