package privateapi

/*
Copyright © Telnyx LLC

Package privateapi implements routines for interacting with private-api service.

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/spf13/viper"
)

type UserQuery struct {
	Id             string `url:"user_id"`
	Email          string `url:"user_email"`
	OrganizationId string `url:"organization_id"`
}

type ConnectionQuery struct {
	Id     string `url:"connection_id"`
	UserId string `url:"user_id"`
}

type UserShort struct {
	Id                         string `json:"user_id"`
	Email                      string `json:"email"`
	OrganizationName           string `json:"organization_name"`
	OrganizationOwnerId        string `json:"organization_owner_id"`
	OrganizationId             string `json:"organization_id"`
	CreatedAt                  string `json:"created_at"`
	BillingCurrency            string `json:"billing_currency"`
	FirstName                  string `json:"first_name"`
	LastName                   string `json:"last_name"`
	AccountCancelled           bool   `json:"account_cancelled"`
	AccountCancelReason        string `json:"account_cancel_reason"`
	CancelledAt                string `json:"cancelled_at"`
	GlobalOutboundChannelLimit int    `json:"global_outbound_channel_limit"`
	AccountBlockedAt           string `json:"account_blocked_at"`
	AccountBlocked             bool   `json:"account_blocked"`
	AutoRechargeEnabled        bool   `json:"auto_recharge_enabled"`
	VerifiedByTelnyx           string `json:"verified_by_telnyx"`
	E911Accepted               bool   `json:"e911_accepted"`
}

type User struct {
	ApiUserType                      int    `json:"api_user_type"`
	LastName                         string `json:"last_name"`
	UserType                         string `json:"user_type"`
	AccountBlockReason               string `json:"account_block_reason"`
	AutoRechargeEnabled              bool   `json:"auto_recharge_enabled"`
	AccountCancelled                 bool   `json:"account_cancelled"`
	GlobalOutboundChannelLimit       int    `json:"global_outbound_channel_limit"`
	AccountBlockedAt                 string `json:"account_blocked_at"`
	PortingUserInfoStrategy          string `json:"porting_user_info_strategy"`
	VerifiedByTelnyx                 string `json:"verified_by_telnyx"`
	City                             string `json:"city"`
	FirstName                        string `json:"first_name"`
	Id                               string `json:"user_id"`
	Zip                              string `json:"zip"`
	AccountBlocked                   bool   `json:"account_blocked"`
	VerifiedPhoneNumberAt            string `json:"verified_phone_number_at"`
	AccountCancelReason              string `json:"account_cancel_reason"`
	AccountBlockedBy                 string `json:"account_blocked_by"`
	OrganizationName                 string `json:"organization_name"`
	State                            string `json:"state"`
	ManagedAccountAllowCustomPricing bool   `json:"managed_account_allow_custom_pricing"`
	CancelledAt                      string `json:"cancelled_at"`
	VerifiedPhoneNumber              string `json:"verified_phone_number"`
	Email                            string `json:"email"`
	PhoneNumber                      string `json:"phone_number"`
	BillingCurrency                  string `json:"billing_currency"`
	ProfileType                      string `json:"profile_type"`
	AccountCancelledBy               string `json:"account_cancelled_by"`
	OrganizationId                   string `json:"organization_id"`
	LastSignInAt                     string `json:"last_sign_in_at"`
	BusinessName                     string `json:"business_name"`
	E911Accepted                     bool   `json:"e911_accepted"`
	LastActualUserSignInAt           string `json:"last_actual_user_sign_in_at"`
	EmailConfirmedAt                 string `json:"email_confirmed_at"`
	OrganizationOwnerId              string `json:"organization_owner_id"`
	DefaultNumbersExternalPin        string `json:"default_numbers_external_pin"`
	CreatedAt                        string `json:"created_at"`
	AccountManagerId                 string `json:"account_manager_id"`
	LastActualUserSignInIp           string `json:"last_actual_user_sign_in_ip"`
	Address1                         string `json:"address_1"`
	Address2                         string `json:"address_2"`
	NpacUser                         bool   `json:"npac_user"`
	Channels                         int    `json:"channels"`
}

type ConnectionShort struct {
	ConnectionId                     int                 `json:"connection_id"`
	ConnectionName                   string              `json:"connection_name"`
	UserId                           string              `json:"user_id"`
	CreatedAt                        string              `json:"created_at"`
	Active                           bool                `json:"active"`
	AuthenticationType               string              `json:"authentication"`
	AnchorsiteOverride               string              `json:"anchorsite_override"`
	WebhookApiVersion                string              `json:"webhook_api_version"`
	Username                         string              `json:"username"`
	WebhoolUrl                       string              `json:"webhook_url"`
	FailoverUrl                      string              `json:"failover_url"`
	WebhookTimeoutSeconds            int                 `json:"webhook_timeout_seconds"`
	XmlRequestUrl                    string              `json:"xml_request_url"`
	XmlRequestMethod                 string              `json:"xml_request_method"`
	XmlFailoverUrl                   string              `json:"xml_failover_url"`
	DtmfType                         string              `json:"dtmf_type"`
	Codecs                           []string            `json:"codecs"`
	InboundChannelLimit              int                 `json:"inbound_channel_limit"`
	OutboundChannelLimit             int                 `json:"outbound_channel_limit"`
	TechPrefix                       string              `json:"tech_prefix"`
	RtcpPort                         string              `json:"rtcp_port"`
	OutboundIpAuthToken              string              `json:"outbound_ip_auth_token"`
	DefaultOnHoldConfortNoiseEnabled bool                `json:"default_on_hold_comfort_noise_enabled"`
	SipSubdomain                     string              `json:"sip_subdomain"`
	IpPorts                          []*ConnectionIpPort `json:"ips_ports"`
	Fqdns                            []string            `json:"fqdns"`
}

type Connection struct {
	XmlFailoverUrl                   string              `json:"xml_failover_url"`
	InboundChannelLimit              int                 `json:"inbound_channel_limit"`
	IsupHeadersEnabled               bool                `json:"isup_headers_enabled"`
	OutboundIpAuthMethod             string              `json:"outbound_ip_auth_method"`
	UserType                         string              `json:"user_type"`
	IpPorts                          []*ConnectionIpPort `json:"ips_ports"`
	XmlStatusCallbackMethod          string              `json:"xml_status_callback_method"`
	PrackEnabled                     bool                `json:"prack_enabled"`
	SipSubdomainReceiveSettings      string              `json:"sip_subdomain_receive_settings"`
	SipRegion                        string              `json:"sip_region"`
	Timeout2xx                       string              `json:"timeout_2xx"`
	AndroidPushCredentialsId         string              `json:"android_push_credential_id"`
	RtcpReportEnabled                bool                `json:"rtcp_report_enabled"`
	TransportProtocol                string              `json:"transport_protocol"`
	RtcpPort                         string              `json:"rtcp_port"`
	Password                         string              `json:"password"`
	UserId                           string              `json:"user_id"`
	DtmfType                         string              `json:"dtmf_type"`
	WebhookTimeoutSeconds            int                 `json:"webhook_timeout_seconds"`
	T38ReinviteSource                string              `json:"t38_reinvite_source"`
	InboundInstantRingbackEnabled    bool                `json:"inbound_instant_ringback_enabled"`
	TechPrefix                       string              `json:"tech_prefix"`
	OutboundLocalization             string              `json:"outbound_localization"`
	AuthenticationType               string              `json:"authentication"`
	DnisNumberFormat                 string              `json:"dnis_number_format"`
	HangupOnTimeout                  bool                `json:"hangup_on_timeout"`
	MediaEncryptionEnabled           bool                `json:"media_encryption_enabled"`
	RtcpReporteFrequencySeconds      int                 `json:"rtcp_report_frequency_seconds"`
	Active                           bool                `json:"active"`
	InstantRingbackEnabled           bool                `json:"instant_ringback_enabled"`
	Username                         string              `json:"username"`
	OutboundIpAuthToken              string              `json:"outbound_ip_auth_token"`
	DefaultOnHoldConfortNoiseEnabled bool                `json:"default_on_hold_comfort_noise_enabled"`
	CallControlTimeout               int                 `json:"call_control_timeout"`
	FqdnOutboundAuthentication       string              `json:"fqdn_outbound_authentication"`
	AniOverride                      string              `json:"ani_override"`
	WebhookApiVersion                string              `json:"webhook_api_version"`
	OnnetT38PassthroughEnabled       bool                `json:"onnet_t38_passthrough_enabled"`
	ConnectionId                     int                 `json:"connection_id"`
	SipSubdomain                     string              `json:"sip_subdomain"`
	Fqdns                            []string            `json:"fqdns"`
	InboundGenerateRingbackTone      bool                `json:"inbound_generate_ringback_tone"`
	AuthenticationPassword           string              `json:"authentication_password"`
	OutboundProfileId                string              `json:"outbound_profile_id"`
	AnchorsiteOverride               string              `json:"anchorsite_override"`
	Codecs                           []string            `json:"codecs"`
	FailoverUrl                      string              `json:"failover_url"`
	XmlRequestMethod                 string              `json:"xml_request_method"`
	RtcpCaptureEnabled               bool                `json:"rtcp_capture_enabled"`
	MicrosoftTeamsSbc                bool                `json:"microsoft_teams_sbc"`
	SipCompactHeadersEnabled         bool                `json:"sip_compact_headers_enabled"`
	XmlStatusCallbackUrl             string              `json:"xml_status_callback_url"`
	IosPushCredentialsId             string              `json:"ios_push_credential_id"`
	OutboundCallParkingEnabled       bool                `json:"outbound_call_parking_enabled"`
	EncryptedMedia                   string              `json:"encrypted_media"`
	Timeout1xx                       string              `json:"timeout_1xx"`
	Ips                              []string            `json:"ips"`
	SipUrlCallingPreference          string              `json:"sip_uri_calling_preference"`
	CreatedAt                        string              `json:"created_at"`
	ConnectionName                   string              `json:"connection_name"`
	OutboundGenerateRingbackTone     bool                `json:"outbound_generate_ringback_tone"`
	PrivacyZoneEnabled               bool                `json:"privacy_zone_enabled"`
	AniNumberFormat                  string              `json:"ani_number_format"`
	EncodeContactHeaderEnabled       bool                `json:"encode_contact_header_enabled"`
	XmlRequestUrl                    string              `json:"xml_request_url"`
	ThirdPartyControlEnabled         bool                `json:"third_party_control_enabled"`
	AniOverrideType                  string              `json:"ani_override_type"`
	WebhoolUrl                       string              `json:"webhook_url"`
	OutboundChannelLimit             int                 `json:"outbound_channel_limit"`
	CallFailureCheckEnabled          bool                `json:"call_failure_check_enabled"`
}

type ConnectionIpPort struct {
	Ip             string `json:"ip"`
	Port           int    `json:"port"`
	AuthorizedIpId string `json:"authorized_ip_id"`
}

func (u *User) ToShort() *UserShort {
	return &UserShort{
		Id:                         u.Id,
		OrganizationId:             u.OrganizationId,
		Email:                      u.Email,
		CreatedAt:                  u.CreatedAt,
		OrganizationOwnerId:        u.OrganizationOwnerId,
		OrganizationName:           u.OrganizationName,
		BillingCurrency:            u.BillingCurrency,
		FirstName:                  u.FirstName,
		LastName:                   u.LastName,
		AccountCancelled:           u.AccountCancelled,
		AccountCancelReason:        u.AccountCancelReason,
		CancelledAt:                u.CancelledAt,
		GlobalOutboundChannelLimit: u.GlobalOutboundChannelLimit,
		AccountBlockedAt:           u.AccountBlockedAt,
		AccountBlocked:             u.AccountBlocked,
		AutoRechargeEnabled:        u.AutoRechargeEnabled,
		VerifiedByTelnyx:           u.VerifiedByTelnyx,
		E911Accepted:               u.E911Accepted,
	}
}

func (conn *Connection) ToShort() *ConnectionShort {
	return &ConnectionShort{
		ConnectionId:                     conn.ConnectionId,
		ConnectionName:                   conn.ConnectionName,
		UserId:                           conn.UserId,
		CreatedAt:                        conn.CreatedAt,
		Active:                           conn.Active,
		AuthenticationType:               conn.AuthenticationType,
		AnchorsiteOverride:               conn.AnchorsiteOverride,
		WebhookApiVersion:                conn.WebhookApiVersion,
		Username:                         conn.Username,
		WebhoolUrl:                       conn.WebhoolUrl,
		FailoverUrl:                      conn.FailoverUrl,
		WebhookTimeoutSeconds:            conn.WebhookTimeoutSeconds,
		XmlRequestUrl:                    conn.XmlRequestUrl,
		XmlRequestMethod:                 conn.XmlRequestMethod,
		XmlFailoverUrl:                   conn.XmlFailoverUrl,
		DtmfType:                         conn.DtmfType,
		Codecs:                           conn.Codecs,
		OutboundChannelLimit:             conn.OutboundChannelLimit,
		TechPrefix:                       conn.TechPrefix,
		RtcpPort:                         conn.RtcpPort,
		OutboundIpAuthToken:              conn.OutboundIpAuthToken,
		DefaultOnHoldConfortNoiseEnabled: conn.DefaultOnHoldConfortNoiseEnabled,
		SipSubdomain:                     conn.SipSubdomain,
		IpPorts:                          conn.IpPorts,
		Fqdns:                            conn.Fqdns,
		InboundChannelLimit:              conn.InboundChannelLimit,
	}
}

func FetchUsers(env string, q *UserQuery) ([]*User, error) {
	e, err := privateApiUrlForEnv(env)
	if err != nil {
		return nil, err
	}

	v, _ := query.Values(q)
	url := fmt.Sprintf("%s/accounts?%s", e, v.Encode())

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "telnyx-cli")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var response []*User
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func FetchConnections(env string, q *ConnectionQuery) ([]*Connection, error) {
	e, err := privateApiUrlForEnv(env)
	if err != nil {
		return nil, err
	}

	v, _ := query.Values(q)
	url := fmt.Sprintf("%s/connections?%s", e, v.Encode())

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "telnyx-cli")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var response []*Connection
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func privateApiUrlForEnv(env string) (string, error) {
	var privateApiUrl string

	if env == "dev" {
		privateApiUrl = viper.GetString("private_api_dev_url")
	} else if env == "prod" {
		privateApiUrl = viper.GetString("private_api_prod_url")
	} else {
		return "", fmt.Errorf("invalid env: %s. Please use dev|prod", env)
	}

	return privateApiUrl, nil
}
