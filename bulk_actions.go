package addyrest

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-get-bulk
func (c *Client) BulkAliasesGet(ids []string) (*AliasesWrap, error) {
	return post[AliasesWrap](c, "api/v1/aliases/get/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-activate-bulk
func (c *Client) BulkAliasesEnable(ids []string) (*BulkWrap, error) {
	return post[BulkWrap](c, "api/v1/aliases/activate/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-deactivate-bulk
func (c *Client) BulkAliasesDisable(ids []string) (*BulkWrap, error) {
	return post[BulkWrap](c, "api/v1/aliases/deactivate/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-delete-bulk
func (c *Client) BulkAliasesDelete(ids []string) (*BulkWrap, error) {
	return post[BulkWrap](c, "api/v1/aliases/delete/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-restore-bulk
func (c *Client) BulkAliasesRestore(ids []string) (*BulkWrap, error) {
	return post[BulkWrap](c, "api/v1/aliases/restore/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-forget-bulk
func (c *Client) BulkAliasesForget(ids []string) (*BulkWrap, error) {
	return post[BulkWrap](c, "api/v1/aliases/forget/bulk",
		IDsGeneric{IDs: ids})
}

// https://app.addy.io/docs/#alias-bulk-actions-POSTapi-v1-aliases-recipients-bulk
func (c *Client) BulkAliasesUpdRecipients(aliasIDs, recipientIDs []string) (*BulkWrap, error) {
	data := struct {
		AliasIDs     []string `json:"ids"`
		RecipientIDs []string `json:"recipient_ids"`
	}{AliasIDs: aliasIDs, RecipientIDs: recipientIDs}
	return post[BulkWrap](c, "api/v1/aliases/recipients/bulk", data)
}
