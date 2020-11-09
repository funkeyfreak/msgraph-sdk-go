/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphRiskDetail the model 'MicrosoftGraphRiskDetail'
type MicrosoftGraphRiskDetail string

// List of microsoft.graph.riskDetail
const (
	MICROSOFTGRAPHRISKDETAIL_NONE                                        MicrosoftGraphRiskDetail = "none"
	MICROSOFTGRAPHRISKDETAIL_ADMIN_GENERATED_TEMPORARY_PASSWORD          MicrosoftGraphRiskDetail = "adminGeneratedTemporaryPassword"
	MICROSOFTGRAPHRISKDETAIL_USER_PERFORMED_SECURED_PASSWORD_CHANGE      MicrosoftGraphRiskDetail = "userPerformedSecuredPasswordChange"
	MICROSOFTGRAPHRISKDETAIL_USER_PERFORMED_SECURED_PASSWORD_RESET       MicrosoftGraphRiskDetail = "userPerformedSecuredPasswordReset"
	MICROSOFTGRAPHRISKDETAIL_ADMIN_CONFIRMED_SIGNIN_SAFE                 MicrosoftGraphRiskDetail = "adminConfirmedSigninSafe"
	MICROSOFTGRAPHRISKDETAIL_AI_CONFIRMED_SIGNIN_SAFE                    MicrosoftGraphRiskDetail = "aiConfirmedSigninSafe"
	MICROSOFTGRAPHRISKDETAIL_USER_PASSED_MFA_DRIVEN_BY_RISK_BASED_POLICY MicrosoftGraphRiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	MICROSOFTGRAPHRISKDETAIL_ADMIN_DISMISSED_ALL_RISK_FOR_USER           MicrosoftGraphRiskDetail = "adminDismissedAllRiskForUser"
	MICROSOFTGRAPHRISKDETAIL_ADMIN_CONFIRMED_SIGNIN_COMPROMISED          MicrosoftGraphRiskDetail = "adminConfirmedSigninCompromised"
	MICROSOFTGRAPHRISKDETAIL_HIDDEN                                      MicrosoftGraphRiskDetail = "hidden"
	MICROSOFTGRAPHRISKDETAIL_ADMIN_CONFIRMED_USER_COMPROMISED            MicrosoftGraphRiskDetail = "adminConfirmedUserCompromised"
	MICROSOFTGRAPHRISKDETAIL_UNKNOWN_FUTURE_VALUE                        MicrosoftGraphRiskDetail = "unknownFutureValue"
)
