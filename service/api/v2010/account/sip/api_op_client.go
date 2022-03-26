// Package sip contains auto-generated files. DO NOT MODIFY
package sip

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_list"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/credential_lists"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domain"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/domains"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_list"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/sip/ip_access_control_lists"
)

// Client for managing Session Initiation Protocol (SIP) resources
type Client struct {
	client *client.Client

	accountSid string

	CredentialList       func(string) *credential_list.Client
	CredentialLists      *credential_lists.Client
	Domain               func(string) *domain.Client
	Domains              *domains.Client
	IpAccessControlList  func(string) *ip_access_control_list.Client
	IpAccessControlLists *ip_access_control_lists.Client
}

// ClientProperties are the properties required to manage the sip resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the sip client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,

		CredentialList: func(credentialListSid string) *credential_list.Client {
			return credential_list.New(client, credential_list.ClientProperties{
				AccountSid: properties.AccountSid,
				Sid:        credentialListSid,
			})
		},
		CredentialLists: credential_lists.New(client, credential_lists.ClientProperties{
			AccountSid: properties.AccountSid,
		}),
		Domain: func(domainSid string) *domain.Client {
			return domain.New(client, domain.ClientProperties{
				AccountSid: properties.AccountSid,
				Sid:        domainSid,
			})
		},
		Domains: domains.New(client, domains.ClientProperties{
			AccountSid: properties.AccountSid,
		}),
		IpAccessControlList: func(ipAccessControlListSid string) *ip_access_control_list.Client {
			return ip_access_control_list.New(client, ip_access_control_list.ClientProperties{
				AccountSid: properties.AccountSid,
				Sid:        ipAccessControlListSid,
			})
		},
		IpAccessControlLists: ip_access_control_lists.New(client, ip_access_control_lists.ClientProperties{
			AccountSid: properties.AccountSid,
		}),
	}
}
