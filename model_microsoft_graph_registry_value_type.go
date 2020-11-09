/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphRegistryValueType the model 'MicrosoftGraphRegistryValueType'
type MicrosoftGraphRegistryValueType string

// List of microsoft.graph.registryValueType
const (
	MICROSOFTGRAPHREGISTRYVALUETYPE_UNKNOWN              MicrosoftGraphRegistryValueType = "unknown"
	MICROSOFTGRAPHREGISTRYVALUETYPE_BINARY               MicrosoftGraphRegistryValueType = "binary"
	MICROSOFTGRAPHREGISTRYVALUETYPE_DWORD                MicrosoftGraphRegistryValueType = "dword"
	MICROSOFTGRAPHREGISTRYVALUETYPE_DWORD_LITTLE_ENDIAN  MicrosoftGraphRegistryValueType = "dwordLittleEndian"
	MICROSOFTGRAPHREGISTRYVALUETYPE_DWORD_BIG_ENDIAN     MicrosoftGraphRegistryValueType = "dwordBigEndian"
	MICROSOFTGRAPHREGISTRYVALUETYPE_EXPAND_SZ            MicrosoftGraphRegistryValueType = "expandSz"
	MICROSOFTGRAPHREGISTRYVALUETYPE_LINK                 MicrosoftGraphRegistryValueType = "link"
	MICROSOFTGRAPHREGISTRYVALUETYPE_MULTI_SZ             MicrosoftGraphRegistryValueType = "multiSz"
	MICROSOFTGRAPHREGISTRYVALUETYPE_NONE                 MicrosoftGraphRegistryValueType = "none"
	MICROSOFTGRAPHREGISTRYVALUETYPE_QWORD                MicrosoftGraphRegistryValueType = "qword"
	MICROSOFTGRAPHREGISTRYVALUETYPE_QWORDLITTLE_ENDIAN   MicrosoftGraphRegistryValueType = "qwordlittleEndian"
	MICROSOFTGRAPHREGISTRYVALUETYPE_SZ                   MicrosoftGraphRegistryValueType = "sz"
	MICROSOFTGRAPHREGISTRYVALUETYPE_UNKNOWN_FUTURE_VALUE MicrosoftGraphRegistryValueType = "unknownFutureValue"
)