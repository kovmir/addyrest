package addyrest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAppGetVersion(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/app-version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("app_version/get_version.json"))
	})

	_, err := client.AppGetVersion()
	if err != nil {
		t.Fatal(err)
	}
}
