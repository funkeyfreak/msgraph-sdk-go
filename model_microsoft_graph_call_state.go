/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphCallState the model 'MicrosoftGraphCallState'
type MicrosoftGraphCallState string

// List of microsoft.graph.callState
const (
	MICROSOFTGRAPHCALLSTATE_INCOMING             MicrosoftGraphCallState = "incoming"
	MICROSOFTGRAPHCALLSTATE_ESTABLISHING         MicrosoftGraphCallState = "establishing"
	MICROSOFTGRAPHCALLSTATE_ESTABLISHED          MicrosoftGraphCallState = "established"
	MICROSOFTGRAPHCALLSTATE_HOLD                 MicrosoftGraphCallState = "hold"
	MICROSOFTGRAPHCALLSTATE_TRANSFERRING         MicrosoftGraphCallState = "transferring"
	MICROSOFTGRAPHCALLSTATE_TRANSFER_ACCEPTED    MicrosoftGraphCallState = "transferAccepted"
	MICROSOFTGRAPHCALLSTATE_REDIRECTING          MicrosoftGraphCallState = "redirecting"
	MICROSOFTGRAPHCALLSTATE_TERMINATING          MicrosoftGraphCallState = "terminating"
	MICROSOFTGRAPHCALLSTATE_TERMINATED           MicrosoftGraphCallState = "terminated"
	MICROSOFTGRAPHCALLSTATE_UNKNOWN_FUTURE_VALUE MicrosoftGraphCallState = "unknownFutureValue"
)
