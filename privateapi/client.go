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
