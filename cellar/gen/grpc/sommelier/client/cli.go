// Code generated by goa v3.11.1, DO NOT EDIT.
//
// sommelier gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"encoding/json"
	"fmt"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
)

// BuildPickPayload builds the payload for the sommelier pick endpoint from CLI
// flags.
func BuildPickPayload(sommelierPickMessage string) (*sommelier.Criteria, error) {
	var err error
	var message sommelierpb.PickRequest
	{
		if sommelierPickMessage != "" {
			err = json.Unmarshal([]byte(sommelierPickMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
			}
		}
	}
	v := &sommelier.Criteria{
		Name:   message.Name,
		Winery: message.Winery,
	}
	if message.Varietal != nil {
		v.Varietal = make([]string, len(message.Varietal))
		for i, val := range message.Varietal {
			v.Varietal[i] = val
		}
	}

	return v, nil
}
