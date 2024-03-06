package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestBulkAliasesGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/get/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/get_aliases.json"))
	})

	_, err := client.BulkAliasesGet([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesEnable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/activate/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/enable_aliases.json"))
	})

	_, err := client.BulkAliasesEnable([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesDisable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/deactivate/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/disable_aliases.json"))
	})

	_, err := client.BulkAliasesDisable([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/delete/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/delete_aliases.json"))
	})

	_, err := client.BulkAliasesDelete([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesRestore(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/restore/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/restore_aliases.json"))
	})

	_, err := client.BulkAliasesRestore([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesForget(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"]}`
	mux.HandleFunc("/api/v1/aliases/forget/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/forget_aliases.json"))
	})

	_, err := client.BulkAliasesForget([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBulkAliasesUpdRecipients(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["50c9e585-e7f5-41c4-9016-9014c15454bc","c549db7d-5fac-4b09-9443-9e47f644d29f"],"recipient_ids":["46eebc50-f7f8-46d7-beb9-c37f04c29a84"]}`
	mux.HandleFunc("/api/v1/aliases/recipients/bulk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("bulk/update_recipients.json"))
	})

	_, err := client.BulkAliasesUpdRecipients([]string{
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
	}, []string{
		"46eebc50-f7f8-46d7-beb9-c37f04c29a84",
	})
	if err != nil {
		t.Fatal(err)
	}
}
