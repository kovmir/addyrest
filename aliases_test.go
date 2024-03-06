package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestAliasesGet(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/aliases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("aliases/get_all.json"))
	})

	_, err := client.AliasesGet(&AliasesGetArgs{
		Filter: map[string]string{
			"deleted": "with",
			"active":  "true",
			"search":  "johndoe",
		},
		SortCond:       AliasSortEmail,
		SortDesc:       true,
		PageNumber:     1,
		PageSize:       10,
		WithRecipients: true,
		Recipient:      "46eebc50-f7f8-46d7-beb9-c37f04c29a84",
		Domain:         "0ad7a75a-1517-4b86-bb8a-9443d4965e60",
		Username:       "46eebc50-f7f8-46d7-beb9-c37f04c29a84",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasGet(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/aliases/50c9e585-e7f5-41c4-9016-9014c15454bc",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("aliases/get_specific.json"))
		})

	_, err := client.AliasGet("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasNew(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"domain":"anonaddy.me",` +
		`"description":"For example.com",` +
		`"format":"uuid",` +
		`"local_part":"hello",` +
		`"recipient_ids":["46eebc50-f7f8-46d7-beb9-c37f04c29a84"]}`

	mux.HandleFunc("/api/v1/aliases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("aliases/create_new.json"))
	})

	_, err := client.AliasNew(&AliasNewArgs{
		Domain:     "anonaddy.me",
		Desc:       "For example.com",
		Format:     AliasFmtUUID,
		LocalPart:  "hello",
		Recipients: []string{"46eebc50-f7f8-46d7-beb9-c37f04c29a84"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasUpdRecipients(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"alias_id":"50c9e585-e7f5-41c4-9016-9014c15454bc","recipient_ids":["46eebc50-f7f8-46d7-beb9-c37f04c29a84"]}`
	mux.HandleFunc("/api/v1/alias-recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("aliases/update_recipients.json"))
	})

	_, err := client.AliasUpdRecipients(&AliasRecipientArgs{
		AliasID:      "50c9e585-e7f5-41c4-9016-9014c15454bc",
		RecipientIDs: []string{"46eebc50-f7f8-46d7-beb9-c37f04c29a84"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasEnable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"50c9e585-e7f5-41c4-9016-9014c15454bc"}`
	mux.HandleFunc("/api/v1/active-aliases",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != bodyExpected {
				t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
			}
			fmt.Fprint(w, fixture("aliases/enable.json"))
		})

	_, err := client.AliasEnable("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/aliases/50c9e585-e7f5-41c4-9016-9014c15454bc",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		})

	err := client.AliasDelete("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasForget(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/aliases/50c9e585-e7f5-41c4-9016-9014c15454bc/forget",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		})

	err := client.AliasForget("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAliasDisable(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/active-aliases/50c9e585-e7f5-41c4-9016-9014c15454bc",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		})

	err := client.AliasDisable("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}

}

func TestAliasUpdate(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"description":"New description","from_name":"John Doe"}`
	mux.HandleFunc("/api/v1/aliases/50c9e585-e7f5-41c4-9016-9014c15454bc",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != bodyExpected {
				t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
			}
			fmt.Fprint(w, fixture("aliases/update.json"))
		})

	_, err := client.AliasUpdate("50c9e585-e7f5-41c4-9016-9014c15454bc",
		&AliasUpdateArgs{
			Desc:     "New description",
			FromName: "John Doe",
		})
	if err != nil {
		t.Fatal(err)
	}

}

func TestAliasRestore(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/aliases/50c9e585-e7f5-41c4-9016-9014c15454bc/restore",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, fixture("aliases/restore.json"))
		})

	_, err := client.AliasRestore("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}

}
