package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestUsersGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/usernames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("usernames/get_all.json"))
	})

	_, err := client.UsersGet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("usernames/get_specific.json"))
	})

	_, err := client.UserGet("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserNew(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"username":"johndoe"}`
	mux.HandleFunc("/api/v1/usernames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/new.json"))
	})

	_, err := client.UserNew("johndoe")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserEnable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"2777dee6-1721-45c0-8d01-698b6be2335f"}`
	mux.HandleFunc("/api/v1/active-usernames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/enable.json"))
	})

	_, err := client.UserEnable("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserEnabCatchAll(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"2777dee6-1721-45c0-8d01-698b6be2335f"}`
	mux.HandleFunc("/api/v1/catch-all-usernames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/enable_catchall.json"))
	})

	_, err := client.UserEnabCatchAll("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserAllowLogin(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"2777dee6-1721-45c0-8d01-698b6be2335f"}`
	mux.HandleFunc("/api/v1/loginable-usernames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/allow_login.json"))
	})

	_, err := client.UserAllowLogin("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.UserDelete("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDisable(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/active-usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.UserDisable("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDisabCatchAll(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/catch-all-usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.UserDisabCatchAll("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDisallowLogin(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/loginable-usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.UserDisallowLogin("2777dee6-1721-45c0-8d01-698b6be2335f")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUpdate(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"description":"New description","from_name":"John"}`
	mux.HandleFunc("/api/v1/usernames/2777dee6-1721-45c0-8d01-698b6be2335f", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/update.json"))
	})

	_, err := client.UserUpdate("2777dee6-1721-45c0-8d01-698b6be2335f", &UserUpdateArgs{
		Desc:     "New description",
		FromName: "John",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserUpdRecipient(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"default_recipient":"46eebc50-f7f8-46d7-beb9-c37f04c29a84"}`
	mux.HandleFunc("/api/v1/usernames/2777dee6-1721-45c0-8d01-698b6be2335f/default-recipient", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("usernames/update.json"))
	})

	_, err := client.UserUpdRecipient("2777dee6-1721-45c0-8d01-698b6be2335f", "46eebc50-f7f8-46d7-beb9-c37f04c29a84")
	if err != nil {
		t.Fatal(err)
	}
}
