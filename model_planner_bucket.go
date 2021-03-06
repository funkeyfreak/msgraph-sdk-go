/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// PlannerBucket struct for PlannerBucket
type PlannerBucket struct {
	Name      string                      `json:"name,omitempty"`
	OrderHint string                      `json:"orderHint,omitempty"`
	PlanId    string                      `json:"planId,omitempty"`
	Tasks     []MicrosoftGraphPlannerTask `json:"tasks,omitempty"`
}
