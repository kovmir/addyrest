package addyrest

// https://app.addy.io/docs/#api-token-GETapi-v1-api-token-details
func (c *Client) TokenGetAPIDetails() (*TokenAPIDetails, error) {
	return get[TokenAPIDetails](c, "api/v1/api-token-details")
}
