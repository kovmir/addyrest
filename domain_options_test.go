package addyrest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDomainGetOpts(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/domain-options", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("domain_options/get.json"))
	})

	_, err := client.DomainGetOpts()
	if err != nil {
		t.Fatal(err)
	}
}
