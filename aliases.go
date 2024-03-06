package addyrest

import "fmt"

type AliasSortCond string
type AliasFormat string

const (
	AliasSortLocalPart  AliasSortCond = "local_part"
	AliasSortDomain     AliasSortCond = "domain"
	AliasSortEmail      AliasSortCond = "email"
	AliasSortEmailsFwd  AliasSortCond = "emails_forwarded"
	AliasSortEmailsBlkd AliasSortCond = "emails_blocked"
	AliasSortEmailsRepl AliasSortCond = "emails_replied"
	AliasSortEmailsSent AliasSortCond = "emails_sent"
	AliasSortActive     AliasSortCond = "active"
	AliasSortCreatedAt  AliasSortCond = "created_at"
	AliasSortUpdatedAt  AliasSortCond = "updated_at"
	AliasSortDeletedAt  AliasSortCond = "deleted_at"

	AliasFmtRndChars AliasFormat = "random_characters"
	AliasFmtUUID     AliasFormat = "uuid"
	AliasFmtRndWords AliasFormat = "random_words"
	AliasFmtCustom   AliasFormat = "custom"
)

type AliasesGetArgs struct {
	// Include deleted or only deleted:
	// Filter[deleted] = with | only
	//
	// Choose to return active or unactive aliases:
	// Filter[active] = true | false
	//
	// Search aliases by email and description.
	// Filter[search] = <search_string>
	Filter map[string]string
	// Paginate the alias results, default 100, min 1 max 100.
	PageSize uint
	// Paginate the alias results; what page number do you want?
	PageNumber uint
	// Sort aliases based on this given condition.
	SortCond AliasSortCond
	// Descending sort?
	SortDesc bool
	// Return aliases with recipients?
	WithRecipients bool
	// Return aliases using the recipient with the specified ID.
	Recipient string
	// Return aliases using the custom domain with the specified ID.
	Domain string
	// Return aliases using the username with the specified ID.
	Username string
}

type AliasNewArgs struct {
	// The domain of the alias.
	Domain string `json:"domain"`
	// The description of the alias
	Desc string `json:"description,omitempty"`
	// The chosen format for the alias.
	Format AliasFormat `json:"format,omitempty"`
	// The chosen local part for the alias (only required if you have the
	// format as custom)
	LocalPart string `json:"local_part,omitempty"`
	// An array of recipient ids to add (the default recipient will be used
	// if none provided)
	Recipients []string `json:"recipient_ids,omitempty"`
}

type AliasRecipientArgs struct {
	AliasID      string   `json:"alias_id"`
	RecipientIDs []string `json:"recipient_ids"`
}

type AliasUpdateArgs struct {
	Desc     string `json:"description,omitempty"`
	FromName string `json:"from_name,omitempty"`
}

// https://app.addy.io/docs/#aliases-GETapi-v1-aliases
func (c *Client) AliasesGet(params *AliasesGetArgs) (*AliasesWrap, error) {
	queryParams := make([]string, 10)
	for k, v := range params.Filter {
		queryParams = append(queryParams, "filter["+k+"]="+v)
	}
	if params.SortCond != "" {
		if params.SortDesc {
			params.SortCond = "-" + params.SortCond
		}
		queryParams = append(queryParams, string(params.SortCond))
	}
	if params.PageNumber > 0 {
		numStr := fmt.Sprintf("%v", params.PageNumber)
		queryParams = append(queryParams, "page[number]="+numStr)
	}
	if params.PageSize > 0 {
		numStr := fmt.Sprintf("%v", params.PageSize)
		queryParams = append(queryParams, "page[size]="+numStr)
	}
	return getWithParams[AliasesWrap](c, "api/v1/aliases", queryParams)
}

// https://app.addy.io/docs/#aliases-GETapi-v1-aliases--id-
func (c *Client) AliasGet(id string) (*AliasWrap, error) {
	return get[AliasWrap](c, "api/v1/aliases/"+id)
}

// https://app.addy.io/docs/#aliases-POSTapi-v1-aliases
func (c *Client) AliasNew(params *AliasNewArgs) (*AliasWrap, error) {
	return post[AliasWrap](c, "api/v1/aliases", params)
}

// https://app.addy.io/docs/#aliases-POSTapi-v1-alias-recipients
func (c *Client) AliasUpdRecipients(data *AliasRecipientArgs) (*AliasWrap, error) {
	return post[AliasWrap](c, "api/v1/alias-recipients", data)
}

// https://app.addy.io/docs/#aliases-POSTapi-v1-active-aliases
func (c *Client) AliasEnable(id string) (*AliasWrap, error) {
	return post[AliasWrap](c, "api/v1/active-aliases", IDGeneric{ID: id})
}

// https://app.addy.io/docs/#aliases-DELETEapi-v1-aliases--id-
func (c *Client) AliasDelete(id string) error {
	_, err := delete[any](c, "api/v1/aliases/"+id)
	return err
}

// https://app.addy.io/docs/#aliases-DELETEapi-v1-aliases--id--forget
func (c *Client) AliasForget(id string) error {
	_, err := delete[any](c, "api/v1/aliases/"+id+"/forget")
	return err
}

// https://app.addy.io/docs/#aliases-DELETEapi-v1-active-aliases--id-
func (c *Client) AliasDisable(id string) error {
	_, err := delete[any](c, "api/v1/active-aliases/"+id)
	return err
}

// https://app.addy.io/docs/#aliases-PATCHapi-v1-aliases--id-
func (c *Client) AliasUpdate(id string, data *AliasUpdateArgs) (*AliasWrap, error) {
	return patch[AliasWrap](c, "api/v1/aliases/"+id, data)
}

// https://app.addy.io/docs/#aliases-PATCHapi-v1-aliases--id--restore
func (c *Client) AliasRestore(id string) (*AliasWrap, error) {
	return patch[AliasWrap](c, "api/v1/aliases/"+id+"/restore", nil)
}
