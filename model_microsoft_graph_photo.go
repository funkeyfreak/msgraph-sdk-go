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

// MicrosoftGraphPhoto struct for MicrosoftGraphPhoto
type MicrosoftGraphPhoto struct {
	CameraMake          string      `json:"cameraMake,omitempty"`
	CameraModel         string      `json:"cameraModel,omitempty"`
	ExposureDenominator interface{} `json:"exposureDenominator,omitempty"`
	ExposureNumerator   interface{} `json:"exposureNumerator,omitempty"`
	FNumber             interface{} `json:"fNumber,omitempty"`
	FocalLength         interface{} `json:"focalLength,omitempty"`
	Iso                 int32       `json:"iso,omitempty"`
	Orientation         int32       `json:"orientation,omitempty"`
	TakenDateTime       time.Time   `json:"takenDateTime,omitempty"`
}
