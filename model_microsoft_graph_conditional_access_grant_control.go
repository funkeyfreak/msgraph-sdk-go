/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConditionalAccessGrantControl the model 'MicrosoftGraphConditionalAccessGrantControl'
type MicrosoftGraphConditionalAccessGrantControl string

// List of microsoft.graph.conditionalAccessGrantControl
const (
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_BLOCK                 MicrosoftGraphConditionalAccessGrantControl = "block"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_MFA                   MicrosoftGraphConditionalAccessGrantControl = "mfa"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_COMPLIANT_DEVICE      MicrosoftGraphConditionalAccessGrantControl = "compliantDevice"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_DOMAIN_JOINED_DEVICE  MicrosoftGraphConditionalAccessGrantControl = "domainJoinedDevice"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_APPROVED_APPLICATION  MicrosoftGraphConditionalAccessGrantControl = "approvedApplication"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_COMPLIANT_APPLICATION MicrosoftGraphConditionalAccessGrantControl = "compliantApplication"
	MICROSOFTGRAPHCONDITIONALACCESSGRANTCONTROL_UNKNOWN_FUTURE_VALUE  MicrosoftGraphConditionalAccessGrantControl = "unknownFutureValue"
)