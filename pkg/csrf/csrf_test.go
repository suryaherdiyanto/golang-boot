package csrf

import (
	"fmt"
	"testing"
)

func TestCSRF(t *testing.T) {
	var token = New("1234")

	if token == "" {
		t.Error("Token is empty")
	}

	fmt.Println(token)
}
