// Code generated by goa v3.7.14, DO NOT EDIT.
//
// api_key_service HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package client

import (
	apikeyservice "goa.design/examples/security/hierarchy/gen/api_key_service"
)

// BuildDefaultPayload builds the payload for the api_key_service default
// endpoint from CLI flags.
func BuildDefaultPayload(apiKeyServiceDefaultKey string) (*apikeyservice.DefaultPayload, error) {
	var key string
	{
		key = apiKeyServiceDefaultKey
	}
	v := &apikeyservice.DefaultPayload{}
	v.Key = key

	return v, nil
}

// BuildSecurePayload builds the payload for the api_key_service secure
// endpoint from CLI flags.
func BuildSecurePayload(apiKeyServiceSecureToken string) (*apikeyservice.SecurePayload, error) {
	var token string
	{
		token = apiKeyServiceSecureToken
	}
	v := &apikeyservice.SecurePayload{}
	v.Token = token

	return v, nil
}
