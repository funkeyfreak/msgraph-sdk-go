/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPlannerGroup struct for MicrosoftGraphPlannerGroup
type MicrosoftGraphPlannerGroup struct {
	Id    string                      `json:"id,omitempty"`
	Plans []MicrosoftGraphPlannerPlan `json:"plans,omitempty"`
}