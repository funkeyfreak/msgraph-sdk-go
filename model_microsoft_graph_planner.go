/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphPlanner struct for MicrosoftGraphPlanner
type MicrosoftGraphPlanner struct {
	Id      string                        `json:"id,omitempty"`
	Buckets []MicrosoftGraphPlannerBucket `json:"buckets,omitempty"`
	Plans   []MicrosoftGraphPlannerPlan   `json:"plans,omitempty"`
	Tasks   []MicrosoftGraphPlannerTask   `json:"tasks,omitempty"`
}