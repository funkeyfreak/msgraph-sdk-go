/*
 * Partial Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph

// DriveItem struct for DriveItem
type DriveItem struct {
	Audio             interface{}                      `json:"audio,omitempty"`
	Content           string                           `json:"content,omitempty"`
	CTag              string                           `json:"cTag,omitempty"`
	Deleted           interface{}                      `json:"deleted,omitempty"`
	File              interface{}                      `json:"file,omitempty"`
	FileSystemInfo    interface{}                      `json:"fileSystemInfo,omitempty"`
	Folder            interface{}                      `json:"folder,omitempty"`
	Image             interface{}                      `json:"image,omitempty"`
	Location          interface{}                      `json:"location,omitempty"`
	Package           interface{}                      `json:"package,omitempty"`
	PendingOperations interface{}                      `json:"pendingOperations,omitempty"`
	Photo             interface{}                      `json:"photo,omitempty"`
	Publication       interface{}                      `json:"publication,omitempty"`
	RemoteItem        interface{}                      `json:"remoteItem,omitempty"`
	Root              interface{}                      `json:"root,omitempty"`
	SearchResult      interface{}                      `json:"searchResult,omitempty"`
	Shared            interface{}                      `json:"shared,omitempty"`
	SharepointIds     interface{}                      `json:"sharepointIds,omitempty"`
	Size              int64                            `json:"size,omitempty"`
	SpecialFolder     interface{}                      `json:"specialFolder,omitempty"`
	Video             interface{}                      `json:"video,omitempty"`
	WebDavUrl         string                           `json:"webDavUrl,omitempty"`
	Workbook          interface{}                      `json:"workbook,omitempty"`
	Analytics         interface{}                      `json:"analytics,omitempty"`
	Children          []MicrosoftGraphDriveItem        `json:"children,omitempty"`
	ListItem          interface{}                      `json:"listItem,omitempty"`
	Permissions       []MicrosoftGraphPermission       `json:"permissions,omitempty"`
	Subscriptions     []MicrosoftGraphSubscription     `json:"subscriptions,omitempty"`
	Thumbnails        []MicrosoftGraphThumbnailSet     `json:"thumbnails,omitempty"`
	Versions          []MicrosoftGraphDriveItemVersion `json:"versions,omitempty"`
}