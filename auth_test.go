package conoha

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	var token string
	var endpoint string
	var ar *AuthResponse

	ar = Authenticate()
	token = GetToken(ar)
	endpoint = GetEndpoint("object-store", ar)

	fmt.Printf("token: %s\n", token)
	fmt.Printf("endpoint: %s\n", endpoint)
}
