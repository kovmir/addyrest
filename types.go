package addyrest

import (
	"strings"
	"time"
)

// Time format returned by Addy.
type UnixTime struct {
	time.Time
}

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
	}
	ut.Time = t
	return nil
}

type AccountDetailsWrap struct {
	Data AccountDetails `json:"data"`
}
type AccountDetails struct {
	ID                           string    `json:"id"`
	Username                     string    `json:"username"`
	FromName                     string    `json:"from_name"`
	EmailSubject                 string    `json:"email_subject"`
	BannerLocation               string    `json:"banner_location"`
	Bandwidth                    uint      `json:"bandwidth"`
	UsernameCount                uint      `json:"username_count"`
	UsernameLimit                uint      `json:"username_limit"`
	DefaultRecipientID           string    `json:"default_recipient_id"`
	DefaultAliasDomain           string    `json:"default_alias_domain"`
	DefaultAliasFormat           string    `json:"default_alias_format"`
	Subscription                 string    `json:"subscription"`
	SubscriptionEndsAt           *UnixTime `json:"subscription_ends_at"`
	BandwidthLimit               uint      `json:"bandwidth_limit"`
	RecipientCount               uint      `json:"recipient_count"`
	RecipientLimit               uint      `json:"recipient_limit"`
	ActiveDomainCount            uint      `json:"active_domain_count"`
	ActiveDomainLimit            uint      `json:"active_domain_limit"`
	ActiveSharedDomainAliasCount uint      `json:"active_shared_domain_alias_count"`
	ActiveSharedDomainAliasLimit uint      `json:"active_shared_domain_alias_limit"`
	TotalEmailsForwarded         uint      `json:"total_emails_forwarded"`
	TotalEmailsBlocked           uint      `json:"total_emails_blocked"`
	TotalEmailsReplied           uint      `json:"total_emails_replied"`
	TotalEmailsSent              uint      `json:"total_emails_sent"`
	CreatedAt                    UnixTime  `json:"created_at"`
	UpdatedAt                    *UnixTime `json:"updated_at"`
}

type AliasWrap struct {
	Data Alias `json:"data"`
}
type AliasesWrap struct {
	Data  []Alias      `json:"data"`
	Links AliasesLinks `json:"links"`
	Meta  AliasesMeta  `json:"meta"`
}
type AliasesLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}
type AliasesMeta struct {
	CurrentPage int                `json:"current_page"`
	From        int                `json:"from"`
	LastPage    int                `json:"last_page"`
	Links       []AliasesMetaLinks `json:"links"`
	Path        string             `json:"path"`
	PerPage     int                `json:"per_page"`
	To          int                `json:"to"`
	Total       int                `json:"total"`
}
type AliasesMetaLinks struct {
	URL    any    `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
type Alias struct {
	ID              string      `json:"id"`
	UserID          string      `json:"user_id"`
	AliasableID     string      `json:"aliasable_id"`
	AliasableType   string      `json:"aliasable_type"`
	LocalPart       string      `json:"local_part"`
	Extension       string      `json:"extension"`
	Domain          string      `json:"domain"`
	Email           string      `json:"email"`
	Active          bool        `json:"active"`
	Description     string      `json:"description"`
	EmailsForwarded uint        `json:"emails_forwarded"`
	EmailsBlocked   uint        `json:"emails_blocked"`
	EmailsReplied   uint        `json:"emails_replied"`
	EmailsSent      uint        `json:"emails_sent"`
	Recipients      []Recipient `json:"recipients"`
	CreatedAt       UnixTime    `json:"created_at"`
	UpdatedAt       *UnixTime   `json:"updated_at"`
	DeletedAt       *UnixTime   `json:"deleted_at"`
}

type RecipientWrap struct {
	Data Recipient `json:"data"`
}
type RecipientsWrap struct {
	Data []Recipient `json:"data"`
}
type Recipient struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	Email            string    `json:"email"`
	CanReplySend     bool      `json:"can_reply_send"`
	ShouldEncrypt    bool      `json:"should_encrypt"`
	InlineEncryption bool      `json:"inline_encryption"`
	ProtectedHeaders bool      `json:"protected_headers"`
	Fingerprint      string    `json:"fingerprint"`
	EmailVerifiedAt  *UnixTime `json:"email_verified_at"`
	Aliases          []Alias   `json:"aliases"`
	CreatedAt        UnixTime  `json:"created_at"`
	UpdatedAt        *UnixTime `json:"updated_at"`
}

type DomainWrap struct {
	Data Domain `json:"data"`
}
type DomainsWrap struct {
	Data []Domain `json:"data"`
}
type Domain struct {
	ID                      string    `json:"id"`
	UserID                  string    `json:"user_id"`
	Domain                  string    `json:"domain"`
	Description             string    `json:"description"`
	Aliases                 []Alias   `json:"aliases"`
	DefaultRecipient        Recipient `json:"default_recipient"`
	Active                  bool      `json:"active"`
	CatchAll                bool      `json:"catch_all"`
	DomainVerifiedAt        *UnixTime `json:"domain_verified_at"`
	DomainMxValidatedAt     *UnixTime `json:"domain_mx_validated_at"`
	DomainSendingVerifiedAt *UnixTime `json:"domain_sending_verified_at"`
	CreatedAt               UnixTime  `json:"created_at"`
	UpdatedAt               *UnixTime `json:"updated_at"`
}

type FailedDeliveryWrap struct {
	Data FailedDelivery `json:"data"`
}
type FailedDeliveriesWrap struct {
	Data []FailedDelivery `json:"data"`
}
type FailedDelivery struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	RecipientID    string    `json:"recipient_id"`
	RecipientEmail string    `json:"recipient_email"`
	AliasID        string    `json:"alias_id"`
	AliasEmail     string    `json:"alias_email"`
	BounceType     string    `json:"bounce_type"`
	RemoteMta      string    `json:"remote_mta"`
	Sender         string    `json:"sender"`
	EmailType      string    `json:"email_type"`
	Status         string    `json:"status"`
	Code           string    `json:"code"`
	AttemptedAt    UnixTime  `json:"attempted_at"`
	CreatedAt      *UnixTime `json:"created_at"`
	UpdatedAt      *UnixTime `json:"updated_at"`
}

type RuleWrap struct {
	Data Rule `json:"data"`
}
type RulesWrap struct {
	Data []Rule `json:"data"`
}
type Rule struct {
	ID         string       `json:"id"`
	UserID     string       `json:"user_id"`
	Name       string       `json:"name"`
	Order      int          `json:"order"`
	Conditions []Conditions `json:"conditions"`
	Actions    []Actions    `json:"actions"`
	Operator   string       `json:"operator"`
	Forwards   bool         `json:"forwards"`
	Replies    bool         `json:"replies"`
	Sends      bool         `json:"sends"`
	Active     bool         `json:"active"`
	CreatedAt  UnixTime     `json:"created_at"`
	UpdatedAt  *UnixTime    `json:"updated_at"`
}
type Conditions struct {
	Type   string   `json:"type"`
	Match  string   `json:"match"`
	Values []string `json:"values"`
}
type Actions struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TokenAPIDetails struct {
	Name      string    `json:"name"`
	CreatedAt UnixTime  `json:"created_at"`
	ExpiresAt *UnixTime `json:"expires_at"`
}

type AppVersion struct {
	Version string `json:"version"`
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Patch   int    `json:"patch"`
}

type DomainOptions struct {
	Data               []string `json:"data"`
	DefaultAliasDomain string   `json:"defaultAliasDomain"`
	DefaultAliasFormat string   `json:"defaultAliasFormat"`
}

type UserWrap struct {
	Data User `json:"data"`
}
type UsersWrap struct {
	Data []User `json:"data"`
}
type User struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	Username         string    `json:"username"`
	Description      string    `json:"description"`
	Aliases          []Alias   `json:"aliases"`
	DefaultRecipient string    `json:"default_recipient"`
	Active           bool      `json:"active"`
	CatchAll         bool      `json:"catch_all"`
	CanLogin         bool      `json:"can_login"`
	CreatedAt        UnixTime  `json:"created_at"`
	UpdatedAt        *UnixTime `json:"updated_at"`
}

// Multiple methods require either ID or an array of IDs as arguments.
type IDGeneric struct {
	ID string `json:"id"`
}
type IDsGeneric struct {
	IDs []string `json:"ids"`
}

// Responce from bulk actions.
type BulkWrap struct {
	Data Bulk `json:"data"`
}
type Bulk struct {
	Message string   `json:"message"`
	IDs     []string `json:"ids"`
}
