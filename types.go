package main

import "time"

type SorareOffers struct {
	TransfertMarket struct {
		SingleSaleOffers struct {
			Edges    []Edge `json:"edges"`
			Nodes    []Node `json:"nodes"`
			PageInfo struct {
				HasNextPage bool   `json:"hasNextPage"`
				EndCursor   string `json:"endCursor"`
				StartCursor string `json:"startCursor"`
			} `json:"pageInfo"`
			TotalCount int64 `json:"totalCount"`
		} `json:"singleSaleOffers"`
	} `json:"transferMarket"`
}

type Node struct {
	Acceptor    interface{} `json:"acceptor"`
	Open        bool        `json:"open"`
	AasmState   string      `json:"aasmState"`
	AcceptedAt  time.Time   `json:"acceptedAt"`
	CancelledAt time.Time   `json:"cancelledAt"`
	Price       string      `json:"price"`
	PriceInFiat struct {
		Eur float64 `json:"eur"`
	} `json:"priceInFiat"`
	StartDate time.Time `json:"startDate"`
	Card      struct {
		Age       int64     `json:"age"`
		CanBuy    bool      `json:"canBuy"`
		CanSell   bool      `json:"canSell"`
		CreatedAt time.Time `json:"createdAt"`
		Player    struct {
			Slug string `json:"slug"`
		} `json:"player"`
		ID        string `json:"id"`
		OnSale    bool   `json:"onSale"`
		Rarity    string `json:"rarity"`
		Xp        int64  `json:"xp"`
		Grade     int64  `json:"grade"`
		CardPrint struct {
			CardEdition struct {
				DisplayName string `json:"displayName"`
				Name        string `json:"name"`
			}
		} `json:"cardPrint"`
		Owner struct {
			ID string `json:"id"`
		} `json:"owner"`
	} `json:"card"`
	Contract struct {
		Name string `json:"name"`
	} `json:"contract"`
	Deal struct {
		DealId                string `json:"dealId"`
		MinReceiveAmountInWei string `json:"minReceiveAmountInWei"`
	} `json:"deal"`
}

type Edge struct {
	Cursor string `json:"cursor"`
	Node   Node   `json:"node"`
}
