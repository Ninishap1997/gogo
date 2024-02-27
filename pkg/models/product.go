package main

import "github.com/shopspring/decimal"

type Product struct {
	Id     int             `json:"id"`
	Name   string          `json:"name"`
	Weight int             `json:"weight"`
	Cost   decimal.Decimal `json:"cost"`
}
