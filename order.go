package main

import "time"

type Order struct {
	ClosedAt       time.Time `json:"closedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Estabilishment *Merchant `json:"merchant"`
	Bag            *Bag      `json:"bag"`
	Payment        *Payment  `json:"payment"`
	LastStatus     string    `json:"lastStatus"`
}

type Merchant struct {
	Name string `json:"name"`
}

type Bag struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Name string `json:"name"`
}

type Payment struct {
	Values *Values `json:"values"`
}

type Values struct {
	Bag int `json:"bag"`
}
