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

// MicrosoftGraphGroupSettingTemplate struct for MicrosoftGraphGroupSettingTemplate
type MicrosoftGraphGroupSettingTemplate struct {
	Id              string                               `json:"id,omitempty"`
	DeletedDateTime time.Time                            `json:"deletedDateTime,omitempty"`
	Description     string                               `json:"description,omitempty"`
	DisplayName     string                               `json:"displayName,omitempty"`
	Values          []MicrosoftGraphSettingTemplateValue `json:"values,omitempty"`
}
