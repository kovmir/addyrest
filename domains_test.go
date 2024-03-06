package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDomainsGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/domains", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("domains/get_all.json"))
	})

	_, err := client.DomainsGet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("domains/get_specific.json"))
	})

	_, err := client.DomainGet("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainNew(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"domain":"example.com"}`
	mux.HandleFunc("/api/v1/domains", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("domains/new.json"))
	})

	_, err := client.DomainNew("example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainEnable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"0ad7a75a-1517-4b86-bb8a-9443d4965e60"}`
	mux.HandleFunc("/api/v1/active-domains", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("domains/enable.json"))
	})

	_, err := client.DomainEnable("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainEnabCatchAll(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"0ad7a75a-1517-4b86-bb8a-9443d4965e60"}`
	mux.HandleFunc("/api/v1/catch-all-domains", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("domains/enable_catchall.json"))
	})

	_, err := client.DomainEnabCatchAll("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.DomainDelete("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainDisable(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/active-domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.DomainDisable("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainDisabCatchAll(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/catch-all-domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.DomainDisabCatchAll("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainUpdate(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"description":"New description","from_name":"Mr Example"}`
	mux.HandleFunc("/api/v1/domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("domains/update.json"))
	})

	_, err := client.DomainUpdate("0ad7a75a-1517-4b86-bb8a-9443d4965e60", &DomainUpdateArgs{
		Desc:     "New description",
		FromName: "Mr Example",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDomainUpdDefRecipient(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"default_recipient":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/domains/0ad7a75a-1517-4b86-bb8a-9443d4965e60/default-recipient", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("domains/update_default_recipient.json"))
	})

	_, err := client.DomainUpdDefRecipient("0ad7a75a-1517-4b86-bb8a-9443d4965e60",
		"46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}
