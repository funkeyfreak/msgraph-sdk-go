/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphRecurrencePatternType the model 'MicrosoftGraphRecurrencePatternType'
type MicrosoftGraphRecurrencePatternType string

// List of microsoft.graph.recurrencePatternType
const (
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_DAILY            MicrosoftGraphRecurrencePatternType = "daily"
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_WEEKLY           MicrosoftGraphRecurrencePatternType = "weekly"
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_ABSOLUTE_MONTHLY MicrosoftGraphRecurrencePatternType = "absoluteMonthly"
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_RELATIVE_MONTHLY MicrosoftGraphRecurrencePatternType = "relativeMonthly"
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_ABSOLUTE_YEARLY  MicrosoftGraphRecurrencePatternType = "absoluteYearly"
	MICROSOFTGRAPHRECURRENCEPATTERNTYPE_RELATIVE_YEARLY  MicrosoftGraphRecurrencePatternType = "relativeYearly"
)