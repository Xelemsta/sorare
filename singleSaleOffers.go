package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

var mapOffersByPlayer = map[string][]Node{}

//curl 'https://api.sorare.com/graphql' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'Origin: altair://-' --data-binary '{"query":"# Welcome to Altair GraphQL Client.\n# You can send your request using CmdOrCtrl + Enter.\n\n# Enter your graphQL query here.\n\nquery{\n  singleSaleOffers(first:500){\n    nodes{\n      aasmState\n      acceptedAt\n      acceptor{\n        nickname\n      }\n      creditCardFee\n      id\n      endDate\n      price\n      priceInFiat{\n        eur\n      }\n      startDate\n      card{\n        player{\n          slug\n        }\n        id\n        rarity\n        owner{\n          id\n        }\n      }\n    }\n    pageInfo{\n      hasNextPage\n      startCursor\n    }\n  }\n}","variables":{}}' --compressed

func getSingleSalesOffers() {
	graphqlClient := graphql.NewClient(SorareUrl)
	graphqlRequest := graphql.NewRequest(singleSaleOffersFirstRequest)
	graphqlRequest.Header.Add("HTTP_APIKEY", apiKey)

	i := 0
	for {
		i++
		sorareOffers, err := callSorareSingleSaleOffers(graphqlClient, graphqlRequest)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Has next page %v \n", sorareOffers.TransfertMarket.SingleSaleOffers.PageInfo.HasNextPage)
		fmt.Printf("pageInfo %+v \n", sorareOffers.TransfertMarket.SingleSaleOffers.PageInfo)
		if !sorareOffers.TransfertMarket.SingleSaleOffers.PageInfo.HasNextPage {
			break
		}

		graphqlRequest = graphql.NewRequest(fmt.Sprintf(singleSaleOffersSecondRequest, sorareOffers.TransfertMarket.SingleSaleOffers.PageInfo.EndCursor))
		graphqlRequest.Header.Add("HTTP_APIKEY", apiKey)

		fmt.Printf("Request number %d \n", i)
	}

	fmt.Print("Works done, regrouping data by players")
	for i, d := range mapOffersByPlayer {
		fmt.Printf("\nPlayer %s, len %d \n", i, len(d))
	}
	fmt.Printf("%d", len(mapOffersByPlayer))

}

func callSorareSingleSaleOffers(graphqlClient *graphql.Client, graphqlRequest *graphql.Request) (SorareOffers, error) {
	sorareOffers := SorareOffers{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &sorareOffers); err != nil {
		return sorareOffers, err
	}

	fmt.Printf("Number of nodes retrieved %d, total count %d \n", len(sorareOffers.TransfertMarket.SingleSaleOffers.Nodes), sorareOffers.TransfertMarket.SingleSaleOffers.TotalCount)
	fmt.Printf("Gathered edges of len %d \n", len(sorareOffers.TransfertMarket.SingleSaleOffers.Edges))

	fmt.Printf("Current size of mapOfferPlayer %d \n", len(mapOffersByPlayer))

	for _, node := range sorareOffers.TransfertMarket.SingleSaleOffers.Nodes {
		if node.Card.Rarity == "rare" {
			mapOffersByPlayer[node.Card.Player.Slug] = append(mapOffersByPlayer[node.Card.Player.Slug], node)
		}
	}

	return sorareOffers, nil
}
