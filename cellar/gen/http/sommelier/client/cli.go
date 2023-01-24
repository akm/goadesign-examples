// Code generated by goa v3.11.0, DO NOT EDIT.
//
// sommelier HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"encoding/json"
	"fmt"

	sommelier "goa.design/examples/cellar/gen/sommelier"
)

// BuildPickPayload builds the payload for the sommelier pick endpoint from CLI
// flags.
func BuildPickPayload(sommelierPickBody string) (*sommelier.Criteria, error) {
	var err error
	var body PickRequestBody
	{
		err = json.Unmarshal([]byte(sommelierPickBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"varietal\": [\n         \"pinot noir\",\n         \"merlot\",\n         \"cabernet franc\"\n      ],\n      \"winery\": \"longoria\"\n   }'")
		}
	}
	v := &sommelier.Criteria{
		Name:   body.Name,
		Winery: body.Winery,
	}
	if body.Varietal != nil {
		v.Varietal = make([]string, len(body.Varietal))
		for i, val := range body.Varietal {
			v.Varietal[i] = val
		}
	}

	return v, nil
}
