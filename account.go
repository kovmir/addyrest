package addyrest

// https://app.addy.io/docs/#account-details-GETapi-v1-account-details
func (c *Client) AccountGetDetails() (*AccountDetailsWrap, error) {
	return get[AccountDetailsWrap](c, "api/v1/account-details")
}
