package main

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
)

var INDEXER_ENDPOINTS = [2]string{
	"https://indexer-mainnet-v1.opti.domains",
	"https://indexer-testnet-v1.opti.domains",
}

type DomainHasName struct {
	Name  string `json:"name"`
	Node  string `json:"node"`
	Owner string `json:"owner"`
}

func getDomainNameFromId(id string) (string, error) {
	idBig := new(big.Int)
	idBig, ok := idBig.SetString(id, 10)
	if !ok {
		return "", errors.New("id is not a number")
	}

	namehash := "0x" + idBig.Text(16)

	for i := 0; i < len(INDEXER_ENDPOINTS); i++ {
		// Send GET request
		resp, err := http.Get(INDEXER_ENDPOINTS[i] + "/node/" + namehash)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Decode JSON response
		var domains []DomainHasName
		err = json.NewDecoder(resp.Body).Decode(&domains)
		if err != nil {
			continue
		}

		if len(domains) > 0 {
			return domains[0].Name, nil
		}
	}

	return "", errors.New("not found")
}
