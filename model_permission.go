/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

import (
	"time"
)

// Permission struct for Permission
type Permission struct {
	ExpirationDateTime  time.Time     `json:"expirationDateTime,omitempty"`
	GrantedTo           interface{}   `json:"grantedTo,omitempty"`
	GrantedToIdentities []interface{} `json:"grantedToIdentities,omitempty"`
	HasPassword         bool          `json:"hasPassword,omitempty"`
	InheritedFrom       interface{}   `json:"inheritedFrom,omitempty"`
	Invitation          interface{}   `json:"invitation,omitempty"`
	Link                interface{}   `json:"link,omitempty"`
	Roles               []string      `json:"roles,omitempty"`
	ShareId             string        `json:"shareId,omitempty"`
}
