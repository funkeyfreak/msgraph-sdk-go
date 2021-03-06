/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphCallRecordsDeviceInfo struct for MicrosoftGraphCallRecordsDeviceInfo
type MicrosoftGraphCallRecordsDeviceInfo struct {
	CaptureDeviceDriver              string      `json:"captureDeviceDriver,omitempty"`
	CaptureDeviceName                string      `json:"captureDeviceName,omitempty"`
	CaptureNotFunctioningEventRatio  interface{} `json:"captureNotFunctioningEventRatio,omitempty"`
	CpuInsufficentEventRatio         interface{} `json:"cpuInsufficentEventRatio,omitempty"`
	DeviceClippingEventRatio         interface{} `json:"deviceClippingEventRatio,omitempty"`
	DeviceGlitchEventRatio           interface{} `json:"deviceGlitchEventRatio,omitempty"`
	HowlingEventCount                int32       `json:"howlingEventCount,omitempty"`
	InitialSignalLevelRootMeanSquare interface{} `json:"initialSignalLevelRootMeanSquare,omitempty"`
	LowSpeechLevelEventRatio         interface{} `json:"lowSpeechLevelEventRatio,omitempty"`
	LowSpeechToNoiseEventRatio       interface{} `json:"lowSpeechToNoiseEventRatio,omitempty"`
	MicGlitchRate                    interface{} `json:"micGlitchRate,omitempty"`
	ReceivedNoiseLevel               int32       `json:"receivedNoiseLevel,omitempty"`
	ReceivedSignalLevel              int32       `json:"receivedSignalLevel,omitempty"`
	RenderDeviceDriver               string      `json:"renderDeviceDriver,omitempty"`
	RenderDeviceName                 string      `json:"renderDeviceName,omitempty"`
	RenderMuteEventRatio             interface{} `json:"renderMuteEventRatio,omitempty"`
	RenderNotFunctioningEventRatio   interface{} `json:"renderNotFunctioningEventRatio,omitempty"`
	RenderZeroVolumeEventRatio       interface{} `json:"renderZeroVolumeEventRatio,omitempty"`
	SentNoiseLevel                   int32       `json:"sentNoiseLevel,omitempty"`
	SentSignalLevel                  int32       `json:"sentSignalLevel,omitempty"`
	SpeakerGlitchRate                interface{} `json:"speakerGlitchRate,omitempty"`
}
