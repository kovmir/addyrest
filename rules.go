package addyrest

type RuleNewParams struct {
	Name       string      `json:"name"`
	Conditions []Condition `json:"conditions"`
	Actions    []Action    `json:"actions"`
	Operator   string      `json:"operator,omitempty"`
	Forwards   bool        `json:"forwards,omitempty"`
	Replies    bool        `json:"replies,omitempty"`
	Sends      bool        `json:"sends,omitempty"`
}
type Condition struct {
	Type   string   `json:"type"`
	Match  string   `json:"match,omitempty"`
	Values []string `json:"values"`
}
type Action struct {
	Type  string `json:"action"`
	Value string `json:"value"`
}

type RuleUpdateArgs struct {
	Name       string       `json:"name"`
	Conditions []Condition  `json:"conditions"`
	Actions    []Action     `json:"actions"`
	Operator   AddyOperator `json:"operator,omitempty"`
	Forwards   bool         `json:"forwards,omitempty"`
	Replies    bool         `json:"replies,omitempty"`
	Sends      bool         `json:"sends,omitempty"`
}

type AddyOperator string

const (
	AddyANDCond AddyOperator = "AND"
	AddyORCond  AddyOperator = "OR"
)

// https://app.addy.io/docs/#rules-GETapi-v1-rules
func (c *Client) RulesGet() (*RulesWrap, error) {
	return get[RulesWrap](c, "api/v1/rules")
}

// https://app.addy.io/docs/#rules-GETapi-v1-rules--id-
func (c *Client) RuleGet(id string) (*RuleWrap, error) {
	return get[RuleWrap](c, "api/v1/rules/"+id)
}

// https://app.addy.io/docs/#rules-POSTapi-v1-rules
func (c *Client) RuleNew(data *RuleNewParams) (*RuleWrap, error) {
	return post[RuleWrap](c, "api/v1/rules", data)
}

// https://app.addy.io/docs/#rules-POSTapi-v1-reorder-rules
func (c *Client) RulesUpdOrder(ids []string) error {
	_, err := post[any](c, "api/v1/reorder-rules", IDsGeneric{IDs: ids})
	return err
}

// https://app.addy.io/docs/#rules-POSTapi-v1-active-rules
func (c *Client) RuleEnable(id string) (*RuleWrap, error) {
	return post[RuleWrap](c, "api/v1/active-rules", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#rules-DELETEapi-v1-rules--id-
func (c *Client) RuleDelete(id string) error {
	_, err := delete[any](c, "api/v1/rules/"+id)
	return err
}

// https://app.addy.io/docs/#rules-DELETEapi-v1-active-rules--id-
func (c *Client) RuleDisable(id string) error {
	_, err := delete[any](c, "api/v1/active-rules/"+id)
	return err
}

// https://app.addy.io/docs/#rules-PATCHapi-v1-rules--id-
func (c *Client) RuleUpdate(id string, data *RuleUpdateArgs) (*RuleWrap, error) {
	return patch[RuleWrap](c, "api/v1/rules/"+id, data)
}
