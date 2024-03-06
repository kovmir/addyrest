package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRecipientsGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("recipients/get_all.json"))
	})

	_, err := client.RecipientsGet(&RecipientsGetArgs{
		Filter: map[string]string{"verified": "true"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("recipients/get_specific.json"))
	})

	_, err := client.RecipientGet("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientNew(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"email":"me@example.com"}`
	mux.HandleFunc("/api/v1/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/new.json"))
	})

	_, err := client.RecipientNew("me@example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientResendEmail(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/recipients/email/resend", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, "{}")
	})

	err := client.RecipientResendEmail("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientEnabEnc(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/encrypted-recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/enable_encryption.json"))
	})

	_, err := client.RecipientEnabEnc("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientEnabEncInl(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/inline-encrypted-recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/enable_inline_encryption.json"))
	})

	_, err := client.RecipientEnabEncInl("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientEnabProtHead(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/protected-headers-recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/enable_protected_headers.json"))
	})

	_, err := client.RecipientEnabProtHead("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientEnabReplSend(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/allowed-recipients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/enable_reply_send.json"))
	})

	_, err := client.RecipientEnabReplSend("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDelete("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDelPubKey(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/recipient-keys/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDelPubKey("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDisabEnc(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/encrypted-recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDisabEnc("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDisabInlEnc(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/inline-encrypted-recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDisabInlEnc("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDisabProtHeads(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/protected-headers-recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDisabProtHeads("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientDisabReplSend(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/allowed-recipients/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RecipientDisabReplSend("46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRecipientAddPubKey(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"key_data":"-----BEGIN PGP PUBLIC KEY BLOCK-----"}`
	mux.HandleFunc("/api/v1/recipient-keys/46eebc50-f7f8-46d7-beb9-c37f04c29a84", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("recipients/add_public_key.json"))
	})

	_, err := client.RecipientAddPubKey("46eebc50-f7f8-46d7-beb9-c37f04c29a84", "-----BEGIN PGP PUBLIC KEY BLOCK-----")
	if err != nil {
		t.Fatal(err)
	}
}
