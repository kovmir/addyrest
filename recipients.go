package addyrest

type RecipientsGetArgs struct {
	Filter map[string]string
}

// https://app.addy.io/docs/#recipients-GETapi-v1-recipients
func (c *Client) RecipientsGet(params *RecipientsGetArgs) (*RecipientsWrap, error) {
	queryParams := make([]string, 10)
	for k, v := range params.Filter {
		queryParams = append(queryParams, "filter["+k+"]="+v)
	}
	return getWithParams[RecipientsWrap](c, "api/v1/recipients", queryParams)
}

// https://app.addy.io/docs/#recipients-GETapi-v1-recipients--id-
func (c *Client) RecipientGet(id string) (*RecipientWrap, error) {
	return get[RecipientWrap](c, "api/v1/recipients/"+id)
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-recipients
func (c *Client) RecipientNew(email string) (*RecipientWrap, error) {
	data := struct {
		Email string `json:"email"`
	}{Email: email}

	return post[RecipientWrap](c, "api/v1/recipients", data)
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-recipients-email-resend
func (c *Client) RecipientResendEmail(id string) error {
	_, err := post[any](c, "api/v1/recipients/email/resend", IDGeneric{ID: id})
	return err
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-encrypted-recipients
func (c *Client) RecipientEnabEnc(id string) (*RecipientWrap, error) {
	return post[RecipientWrap](c, "api/v1/encrypted-recipients", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-inline-encrypted-recipients
func (c *Client) RecipientEnabEncInl(id string) (*RecipientWrap, error) {
	return post[RecipientWrap](c, "api/v1/inline-encrypted-recipients", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-protected-headers-recipients
func (c *Client) RecipientEnabProtHead(id string) (*RecipientWrap, error) {
	return post[RecipientWrap](c, "api/v1/protected-headers-recipients", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#recipients-POSTapi-v1-allowed-recipients
func (c *Client) RecipientEnabReplSend(id string) (*RecipientWrap, error) {
	return post[RecipientWrap](c, "api/v1/allowed-recipients", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-recipients--id-
func (c *Client) RecipientDelete(id string) error {
	_, err := delete[any](c, "api/v1/recipients/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-recipient-keys--id-
func (c *Client) RecipientDelPubKey(id string) error {
	_, err := delete[any](c, "api/v1/recipient-keys/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-encrypted-recipients--id-
func (c *Client) RecipientDisabEnc(id string) error {
	_, err := delete[any](c, "api/v1/encrypted-recipients/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-inline-encrypted-recipients--id-
func (c *Client) RecipientDisabInlEnc(id string) error {
	_, err := delete[any](c, "api/v1/inline-encrypted-recipients/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-protected-headers-recipients--id-
func (c *Client) RecipientDisabProtHeads(id string) error {
	_, err := delete[any](c, "api/v1/protected-headers-recipients/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-DELETEapi-v1-allowed-recipients--id-
func (c *Client) RecipientDisabReplSend(id string) error {
	_, err := delete[any](c, "api/v1/allowed-recipients/"+id)
	return err
}

// https://app.addy.io/docs/#recipients-PATCHapi-v1-recipient-keys--id-
func (c *Client) RecipientAddPubKey(id, pubKey string) (*RecipientWrap, error) {
	data := struct {
		PubKey string `json:"key_data"`
	}{PubKey: pubKey}
	return patch[RecipientWrap](c, "api/v1/recipient-keys/"+id, data)
}
