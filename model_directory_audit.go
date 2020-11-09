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

// DirectoryAudit struct for DirectoryAudit
type DirectoryAudit struct {
	ActivityDateTime    time.Time                            `json:"activityDateTime,omitempty"`
	ActivityDisplayName string                               `json:"activityDisplayName,omitempty"`
	AdditionalDetails   []interface{}                        `json:"additionalDetails,omitempty"`
	Category            string                               `json:"category,omitempty"`
	CorrelationId       string                               `json:"correlationId,omitempty"`
	InitiatedBy         MicrosoftGraphAuditActivityInitiator `json:"initiatedBy,omitempty"`
	LoggedByService     string                               `json:"loggedByService,omitempty"`
	OperationType       string                               `json:"operationType,omitempty"`
	Result              interface{}                          `json:"result,omitempty"`
	ResultReason        string                               `json:"resultReason,omitempty"`
	TargetResources     []interface{}                        `json:"targetResources,omitempty"`
}
