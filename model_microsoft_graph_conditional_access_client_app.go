/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphConditionalAccessClientApp the model 'MicrosoftGraphConditionalAccessClientApp'
type MicrosoftGraphConditionalAccessClientApp string

// List of microsoft.graph.conditionalAccessClientApp
const (
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_ALL                             MicrosoftGraphConditionalAccessClientApp = "all"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_BROWSER                         MicrosoftGraphConditionalAccessClientApp = "browser"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_MOBILE_APPS_AND_DESKTOP_CLIENTS MicrosoftGraphConditionalAccessClientApp = "mobileAppsAndDesktopClients"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_EXCHANGE_ACTIVE_SYNC            MicrosoftGraphConditionalAccessClientApp = "exchangeActiveSync"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_EAS_SUPPORTED                   MicrosoftGraphConditionalAccessClientApp = "easSupported"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_OTHER                           MicrosoftGraphConditionalAccessClientApp = "other"
	MICROSOFTGRAPHCONDITIONALACCESSCLIENTAPP_UNKNOWN_FUTURE_VALUE            MicrosoftGraphConditionalAccessClientApp = "unknownFutureValue"
)