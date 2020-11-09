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

// MicrosoftGraphContract struct for MicrosoftGraphContract
type MicrosoftGraphContract struct {
	Id                string    `json:"id,omitempty"`
	DeletedDateTime   time.Time `json:"deletedDateTime,omitempty"`
	ContractType      string    `json:"contractType,omitempty"`
	CustomerId        string    `json:"customerId,omitempty"`
	DefaultDomainName string    `json:"defaultDomainName,omitempty"`
	DisplayName       string    `json:"displayName,omitempty"`
}
