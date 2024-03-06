package addyrest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAccountGetDetails(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/account-details", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("token/details.json"))
	})

	_, err := client.AccountGetDetails()
	if err != nil {
		t.Fatal(err)
	}
}
