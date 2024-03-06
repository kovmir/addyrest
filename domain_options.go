package addyrest

func (c *Client) DomainGetOpts() (*DomainOptions, error) {
	return get[DomainOptions](c, "api/v1/domain-options")
}
