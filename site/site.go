package site

import (
	"time"
)

type Customer struct {
	id   string
	site Site
}

type Site struct {
	id         string
	customerId string
	plantings  []Planting
}

type Plant struct {
	id string
}

type Seed struct {
	id string
}

type Death struct {
	id     string
	amount int
	reason string
}

type Planting struct {
	id        string
	seedId    string
	createdOn time.Time
	planted   int
	female    int
	male      int
	deaths    []Death
}
