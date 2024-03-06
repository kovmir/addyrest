package addyrest

type UserUpdateArgs struct {
	Desc     string `json:"description,omitempty"`
	FromName string `json:"from_name,omitempty"`
}

// https://app.addy.io/docs/#usernames-GETapi-v1-usernames
func (c *Client) UsersGet() (*UsersWrap, error) {
	return get[UsersWrap](c, "api/v1/usernames")
}

// https://app.addy.io/docs/#usernames-GETapi-v1-usernames--id-
func (c *Client) UserGet(id string) (*UserWrap, error) {
	return get[UserWrap](c, "api/v1/usernames/"+id)
}

// https://app.addy.io/docs/#usernames-POSTapi-v1-usernames
func (c *Client) UserNew(username string) (*UserWrap, error) {
	data := struct {
		Username string `json:"username"`
	}{Username: username}
	return post[UserWrap](c, "api/v1/usernames", data)
}

// https://app.addy.io/docs/#usernames-POSTapi-v1-active-usernames
func (c *Client) UserEnable(id string) (*UserWrap, error) {
	return post[UserWrap](c, "api/v1/active-usernames", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#usernames-POSTapi-v1-catch-all-usernames
func (c *Client) UserEnabCatchAll(id string) (*UserWrap, error) {
	return post[UserWrap](c, "api/v1/catch-all-usernames", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#usernames-POSTapi-v1-loginable-usernames
func (c *Client) UserAllowLogin(id string) (*UserWrap, error) {
	return post[UserWrap](c, "api/v1/loginable-usernames", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#usernames-DELETEapi-v1-usernames--id-
func (c *Client) UserDelete(id string) error {
	_, err := delete[any](c, "api/v1/usernames/"+id)
	return err
}

// https://app.addy.io/docs/#usernames-DELETEapi-v1-active-usernames--id-
func (c *Client) UserDisable(id string) error {
	_, err := delete[any](c, "api/v1/active-usernames/"+id)
	return err
}

// https://app.addy.io/docs/#usernames-DELETEapi-v1-catch-all-usernames--id-
func (c *Client) UserDisabCatchAll(id string) error {
	_, err := delete[any](c, "api/v1/catch-all-usernames/"+id)
	return err
}

// https://app.addy.io/docs/#usernames-DELETEapi-v1-catch-all-usernames--id-
func (c *Client) UserDisallowLogin(id string) error {
	_, err := delete[any](c, "api/v1/loginable-usernames/"+id)
	return err
}

// https://app.addy.io/docs/#usernames-PATCHapi-v1-usernames--id-
func (c *Client) UserUpdate(id string, data *UserUpdateArgs) (*UserWrap, error) {
	return patch[UserWrap](c, "api/v1/usernames/"+id, data)
}

// https://app.addy.io/docs/#usernames-PATCHapi-v1-usernames--id--default-recipient
func (c *Client) UserUpdRecipient(id, recipent string) (*UserWrap, error) {
	data := struct {
		DefaultRecipient string `json:"default_recipient"`
	}{DefaultRecipient: recipent}
	return patch[UserWrap](c, "api/v1/usernames/"+id+"/default-recipient", data)
}
