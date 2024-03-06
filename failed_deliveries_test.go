package addyrest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFailedDeliveriesGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/failed-deliveries", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("failed_deliveries/get_all.json"))
	})

	_, err := client.FailedDeliveriesGet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFailedDeliveryGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/failed-deliveries/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("failed_deliveries/get_specific.json"))
	})

	_, err := client.FailedDeliveryGet("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFailedDeliveryDel(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/failed-deliveries/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.FailedDeliveryDel("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}
