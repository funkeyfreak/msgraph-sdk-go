/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphOfficeGraphInsights struct for MicrosoftGraphOfficeGraphInsights
type MicrosoftGraphOfficeGraphInsights struct {
	Id       string                        `json:"id,omitempty"`
	Shared   []MicrosoftGraphSharedInsight `json:"shared,omitempty"`
	Trending []MicrosoftGraphTrending      `json:"trending,omitempty"`
	Used     []MicrosoftGraphUsedInsight   `json:"used,omitempty"`
}