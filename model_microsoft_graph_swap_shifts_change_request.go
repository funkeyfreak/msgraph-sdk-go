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

// MicrosoftGraphSwapShiftsChangeRequest struct for MicrosoftGraphSwapShiftsChangeRequest
type MicrosoftGraphSwapShiftsChangeRequest struct {
	Id                      string      `json:"id,omitempty"`
	CreatedDateTime         time.Time   `json:"createdDateTime,omitempty"`
	LastModifiedBy          interface{} `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    time.Time   `json:"lastModifiedDateTime,omitempty"`
	AssignedTo              interface{} `json:"assignedTo,omitempty"`
	ManagerActionDateTime   time.Time   `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage    string      `json:"managerActionMessage,omitempty"`
	ManagerUserId           string      `json:"managerUserId,omitempty"`
	SenderDateTime          time.Time   `json:"senderDateTime,omitempty"`
	SenderMessage           string      `json:"senderMessage,omitempty"`
	SenderUserId            string      `json:"senderUserId,omitempty"`
	State                   interface{} `json:"state,omitempty"`
	RecipientActionDateTime time.Time   `json:"recipientActionDateTime,omitempty"`
	RecipientActionMessage  string      `json:"recipientActionMessage,omitempty"`
	RecipientUserId         string      `json:"recipientUserId,omitempty"`
	SenderShiftId           string      `json:"senderShiftId,omitempty"`
	RecipientShiftId        string      `json:"recipientShiftId,omitempty"`
}