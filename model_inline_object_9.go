/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	Prompts                        []interface{} `json:"prompts,omitempty"`
	BargeInAllowed                 bool          `json:"bargeInAllowed,omitempty"`
	InitialSilenceTimeoutInSeconds int32         `json:"initialSilenceTimeoutInSeconds,omitempty"`
	MaxSilenceTimeoutInSeconds     int32         `json:"maxSilenceTimeoutInSeconds,omitempty"`
	MaxRecordDurationInSeconds     int32         `json:"maxRecordDurationInSeconds,omitempty"`
	PlayBeep                       bool          `json:"playBeep,omitempty"`
	StopTones                      []string      `json:"stopTones,omitempty"`
	ClientContext                  string        `json:"clientContext,omitempty"`
}
