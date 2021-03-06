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

// MicrosoftGraphCallRecordsMediaStream struct for MicrosoftGraphCallRecordsMediaStream
type MicrosoftGraphCallRecordsMediaStream struct {
	AverageAudioDegradation                  interface{} `json:"averageAudioDegradation,omitempty"`
	AverageAudioNetworkJitter                string      `json:"averageAudioNetworkJitter,omitempty"`
	AverageBandwidthEstimate                 int64       `json:"averageBandwidthEstimate,omitempty"`
	AverageJitter                            string      `json:"averageJitter,omitempty"`
	AveragePacketLossRate                    interface{} `json:"averagePacketLossRate,omitempty"`
	AverageRatioOfConcealedSamples           interface{} `json:"averageRatioOfConcealedSamples,omitempty"`
	AverageReceivedFrameRate                 interface{} `json:"averageReceivedFrameRate,omitempty"`
	AverageRoundTripTime                     string      `json:"averageRoundTripTime,omitempty"`
	AverageVideoFrameLossPercentage          interface{} `json:"averageVideoFrameLossPercentage,omitempty"`
	AverageVideoFrameRate                    interface{} `json:"averageVideoFrameRate,omitempty"`
	AverageVideoPacketLossRate               interface{} `json:"averageVideoPacketLossRate,omitempty"`
	EndDateTime                              time.Time   `json:"endDateTime,omitempty"`
	LowFrameRateRatio                        interface{} `json:"lowFrameRateRatio,omitempty"`
	LowVideoProcessingCapabilityRatio        interface{} `json:"lowVideoProcessingCapabilityRatio,omitempty"`
	MaxAudioNetworkJitter                    string      `json:"maxAudioNetworkJitter,omitempty"`
	MaxJitter                                string      `json:"maxJitter,omitempty"`
	MaxPacketLossRate                        interface{} `json:"maxPacketLossRate,omitempty"`
	MaxRatioOfConcealedSamples               interface{} `json:"maxRatioOfConcealedSamples,omitempty"`
	MaxRoundTripTime                         string      `json:"maxRoundTripTime,omitempty"`
	PacketUtilization                        int64       `json:"packetUtilization,omitempty"`
	PostForwardErrorCorrectionPacketLossRate interface{} `json:"postForwardErrorCorrectionPacketLossRate,omitempty"`
	StartDateTime                            time.Time   `json:"startDateTime,omitempty"`
	StreamDirection                          interface{} `json:"streamDirection,omitempty"`
	StreamId                                 string      `json:"streamId,omitempty"`
	WasMediaBypassed                         bool        `json:"wasMediaBypassed,omitempty"`
}
