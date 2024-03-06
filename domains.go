package addyrest

type DomainUpdateArgs struct {
	Desc     string `json:"description,omitempty"`
	FromName string `json:"from_name,omitempty"`
}

// https://app.addy.io/docs/#domains-GETapi-v1-domains
func (c *Client) DomainsGet() (*DomainsWrap, error) {
	return get[DomainsWrap](c, "api/v1/domains")
}

// https://app.addy.io/docs/#domains-GETapi-v1-domains--id-
func (c *Client) DomainGet(id string) (*DomainWrap, error) {
	return get[DomainWrap](c, "api/v1/domains/"+id)
}

// https://app.addy.io/docs/#domains-POSTapi-v1-domains
func (c *Client) DomainNew(domain string) (*DomainWrap, error) {
	data := struct {
		Domain string `json:"domain"`
	}{Domain: domain}
	return post[DomainWrap](c, "api/v1/domains", data)
}

// https://app.addy.io/docs/#domains-POSTapi-v1-active-domains
func (c *Client) DomainEnable(id string) (*DomainWrap, error) {
	return post[DomainWrap](c, "api/v1/active-domains", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#domains-POSTapi-v1-catch-all-domains
func (c *Client) DomainEnabCatchAll(id string) (*DomainWrap, error) {
	return post[DomainWrap](c, "api/v1/catch-all-domains", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#domains-DELETEapi-v1-domains--id-
func (c *Client) DomainDelete(id string) error {
	_, err := delete[any](c, "api/v1/domains/"+id)
	return err
}

// https://app.addy.io/docs/#domains-DELETEapi-v1-active-domains--id-
func (c *Client) DomainDisable(id string) error {
	_, err := delete[any](c, "api/v1/active-domains/"+id)
	return err
}

// https://app.addy.io/docs/#domains-DELETEapi-v1-catch-all-domains--id-
func (c *Client) DomainDisabCatchAll(id string) error {
	_, err := delete[any](c, "api/v1/catch-all-domains/"+id)
	return err
}

// https://app.addy.io/docs/#domains-PATCHapi-v1-domains--id-
func (c *Client) DomainUpdate(id string, data *DomainUpdateArgs) (*DomainWrap, error) {
	return patch[DomainWrap](c, "api/v1/domains/"+id, data)
}

// https://app.addy.io/docs/#domains-PATCHapi-v1-domains--id--default-recipient
func (c *Client) DomainUpdDefRecipient(id, recipient string) (*DomainWrap, error) {
	data := struct {
		DefaultRecipient string `json:"default_recipient,omitempty"`
	}{DefaultRecipient: recipient}
	return patch[DomainWrap](c, "api/v1/domains/"+id+"/default-recipient", data)
}
