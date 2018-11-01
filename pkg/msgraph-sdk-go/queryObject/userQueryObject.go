package queryObject

// A generic query object for keeping a UPN and an objectId (calendarId, eventId, etc.)
type UserQueryObject struct {
	UserPrincipalName string `json:"userPrincipalName"`
	AuthenticatedUserObjectId string `json:"authenticatedUserObjectId"`
}
