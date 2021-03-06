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

// OnenotePage struct for OnenotePage
type OnenotePage struct {
	Content              string      `json:"content,omitempty"`
	ContentUrl           string      `json:"contentUrl,omitempty"`
	CreatedByAppId       string      `json:"createdByAppId,omitempty"`
	LastModifiedDateTime time.Time   `json:"lastModifiedDateTime,omitempty"`
	Level                int32       `json:"level,omitempty"`
	Links                interface{} `json:"links,omitempty"`
	Order                int32       `json:"order,omitempty"`
	Title                string      `json:"title,omitempty"`
	UserTags             []string    `json:"userTags,omitempty"`
	ParentNotebook       interface{} `json:"parentNotebook,omitempty"`
	ParentSection        interface{} `json:"parentSection,omitempty"`
}
