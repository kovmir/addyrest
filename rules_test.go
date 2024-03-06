package addyrest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestRulesGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/rules", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("rules/get_all.json"))
	})

	_, err := client.RulesGet()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleGet(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/rules/50c9e585-e7f5-41c4-9016-9014c15454bc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("rules/get_specific.json"))
	})

	_, err := client.RuleGet("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleNew(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"name":"First Rule","conditions":[{"type":"sender","match":"is exactly","values":["will@anonaddy.com"]}],"actions":[{"action":"subject","value":"New Subject!"}],"operator":"AND","forwards":true}`
	mux.HandleFunc("/api/v1/rules", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("rules/new.json"))
	})

	_, err := client.RuleNew(&RuleNewParams{
		Name: "First Rule",
		Conditions: []Condition{{
			Type:   "sender",
			Match:  "is exactly",
			Values: []string{"will@anonaddy.com"},
		}},
		Actions: []Action{{
			Type:  "subject",
			Value: "New Subject!",
		}},
		Operator: "AND",
		Forwards: true,
		Replies:  false,
		Sends:    false,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRulesUpdOrder(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"ids":["c549db7d-5fac-4b09-9443-9e47f644d29f","50c9e585-e7f5-41c4-9016-9014c15454bc"]}`
	mux.HandleFunc("/api/v1/reorder-rules", func(w http.ResponseWriter, r *http.Request) {
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

	err := client.RulesUpdOrder([]string{
		"c549db7d-5fac-4b09-9443-9e47f644d29f",
		"50c9e585-e7f5-41c4-9016-9014c15454bc",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleEnable(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"id":"0ad7a75a-1517-4b86-bb8a-9443d4965e60"}`
	mux.HandleFunc("/api/v1/active-rules", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != bodyExpected {
			t.Fatalf("Unexpected incoming JSON.\nHave `%s`;\nWant `%s`\n", body, bodyExpected)
		}
		fmt.Fprint(w, fixture("rules/enable.json"))
	})

	_, err := client.RuleEnable("0ad7a75a-1517-4b86-bb8a-9443d4965e60")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleDelete(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/rules/50c9e585-e7f5-41c4-9016-9014c15454bc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RuleDelete("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleDisable(t *testing.T) {
	teardown := setup()
	defer teardown()
	mux.HandleFunc("/api/v1/active-rules/50c9e585-e7f5-41c4-9016-9014c15454bc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	err := client.RuleDisable("50c9e585-e7f5-41c4-9016-9014c15454bc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuleUpdate(t *testing.T) {
	teardown := setup()
	defer teardown()
	bodyExpected := `{"name":"New Name","conditions":[{"type":"sender","match":"is exactly","values":["will@anonaddy.com"]}],"actions":[{"action":"subject","value":"New Subject!"}],"operator":"AND","forwards":true}`
	mux.HandleFunc("/api/v1/rules/50c9e585-e7f5-41c4-9016-9014c15454bc", func(w http.ResponseWriter, r *http.Request) {
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

	_, err := client.RuleUpdate("50c9e585-e7f5-41c4-9016-9014c15454bc", &RuleUpdateArgs{
		Name: "New Name",
		Conditions: []Condition{{
			Type:   "sender",
			Match:  "is exactly",
			Values: []string{"will@anonaddy.com"},
		}},
		Actions: []Action{{
			Type:  "subject",
			Value: "New Subject!",
		}},
		Operator: "AND",
		Forwards: true,
		Replies:  false,
		Sends:    false,
	})
	if err != nil {
		t.Fatal(err)
	}
}
