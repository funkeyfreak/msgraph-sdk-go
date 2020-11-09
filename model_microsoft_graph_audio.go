/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// MicrosoftGraphAudio struct for MicrosoftGraphAudio
type MicrosoftGraphAudio struct {
	Album             string `json:"album,omitempty"`
	AlbumArtist       string `json:"albumArtist,omitempty"`
	Artist            string `json:"artist,omitempty"`
	Bitrate           int64  `json:"bitrate,omitempty"`
	Composers         string `json:"composers,omitempty"`
	Copyright         string `json:"copyright,omitempty"`
	Disc              int32  `json:"disc,omitempty"`
	DiscCount         int32  `json:"discCount,omitempty"`
	Duration          int64  `json:"duration,omitempty"`
	Genre             string `json:"genre,omitempty"`
	HasDrm            bool   `json:"hasDrm,omitempty"`
	IsVariableBitrate bool   `json:"isVariableBitrate,omitempty"`
	Title             string `json:"title,omitempty"`
	Track             int32  `json:"track,omitempty"`
	TrackCount        int32  `json:"trackCount,omitempty"`
	Year              int32  `json:"year,omitempty"`
}
