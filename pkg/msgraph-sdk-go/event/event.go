package event

import (
	"../calendar"
	"../queryObject"
	"time"
)

type Event struct {
	Attendees                     interface{}          `json:"attendees"` // TODO: AttendeeCollection
	Body                          interface{}          `json:"body"`      // TODO: ItemBody
	BodyPreview                   string            `json:"bodyPreview"`
	Categories                    []string          `json:"categories"`
	ChangeKey                     string            `json:"changeKey"`
	CreatedDateTime               time.Time         `json:"createdDateTime"`
	End                           time.Time         `json:"end"`
	HasAttachments                bool              `json:"HasAttachments"`
	ICalUId                       string            `json:"iCalUId"`
	Id                            string            `json:"id"`
	Importance                    string            `json:"importance"` // TODO: create enum with ["low", "normal", "high"] options
	IsAllDay                      bool              `json:"isAllDay"`
	IsCancelled                   bool              `json:"isCancelled"`
	IsOrganizer                   bool              `json:"isOrganizer"`
	IsReminderOn                  bool              `json:"isReminderOn"`
	LastModifiedDateTime          time.Time         `json:"lastModifiedDateTime"`
	Location                      interface{}          `json:"location"`  // TODO: create Location
	Locations                     interface{}          `json:"locations"` // TODO create LocationCollection
	OnlineMeetingUrl              string            `json:"onlineMeetingUrl"`
	Organizer                     interface{}          `json:"organizer"` // TODO: create Recipient -- this seems superfluous as it only contains emailAddress
	OriginalEndTimeZone           string            `json:"originalEndTimeZone"`
	OriginalStart                 time.Time         `json:"originalStart"`
	OriginalStartTimeZone         string            `json:"originalStartTimeZone"`
	Recurrence                    interface{}          `json:"recurrence"` // TODO: create PatternedRecurrence
	ReminderMinutesBeforeStart    int               `json:"reminderMinutesBeforeStart"`
	ResponseRequested             bool              `json:"responseRequested"`
	ResponseStatus                interface{}          `json:"responseStatus"` // TODO: create ResponseStatus
	Sensitivity                   string            `json:"sensitivity"`    // ["normal", "private", "personal", "confidential"] options only
	SeriesMasterId                string            `json:"seriesMasterId"`
	ShowAs                        string            `json:"showAs"` // ["free", "tentative", "busy", "oof", "workingElsewhere", "unknown"]
	Start                         time.Time         `json:"start"`
	Subject                       string            `json:"subject"`
	Type                          string            `json:"type"` // ["singleInstance", "occurrence", "exception", "seriesMaster"]
	WebLink                       string            `json:"webLink"`
	Attachments                   interface{}          `json:"attachments,omitempty"`
	Calendar                      calendar.Calendar `json:"calendar"`
	Extensions                    interface{}          `json:"extensions"`                              // TODO: ExtensionCollection
	Instances                     interface{}          `json:"instances"`                               // TODO: EventCollection
	MultiValueExtendedProperties  interface{}          `json:"multiValueExtendedProperties,omitempty"`  // TODO: MultiValueExtendedPropertiesCollection
	SingleValueExtendedProperties interface{}          `json:"singleValueExtendedProperties,omitempty"` // TODO: SingleValueExtendedPropertiesCollection
}

type IEvent interface {
	ListEvents() interface{} // TODO: EventCollection
	CreateEvent(event Event) Event
	GetEvent(userQueryObject queryObject.UserQueryObject, eventId string) Event
	UpdateEvent(userQueryObject queryObject.UserQueryObject, eventId string, attendees interface{}, body interface{},
		categories string, end time.Time, importance string, isAllDay bool, isReminderOn bool, location interface{},
		locations interface{}, recurrence interface{}, reminderMinutesBeforeStart int, responseRequested bool,
		sensitivity string, showAs string, start time.Time, subject string) Event
	Delete(userQueryObject queryObject.UserQueryObject, eventId string)
	Cancel(userQueryObject queryObject.UserQueryObject, eventId string, comment string)
	Accept(userQueryObject queryObject.UserQueryObject, eventId string, comment string, sendResponse bool)
	TentativelyAccept(userQueryObject queryObject.UserQueryObject, eventId string, comment string, sendResponse bool)
	Decline(userQueryObject queryObject.UserQueryObject, eventId string, comment string, sendResponse bool)
	Forward(userQueryObject queryObject.UserQueryObject, eventId string, comment string, toRecipients interface{}) // TODO: RecipientCollection
	Delta(userQueryObject queryObject.UserQueryObject, startDateTime time.Time, endDateTime time.Time, deltaToken string, skipToken string) interface{} // TODO: EventCollection
	DismissReminder(userQueryObject queryObject.UserQueryObject, eventId string)
	SnoozeReminder(userQueryObject queryObject.UserQueryObject, eventId string, newReminderTime time.Time)
	ListInstances(userQueryObject queryObject.UserQueryObject, eventId string, startDateTime time.Time, endDateTime time.Time)
	ListAttachments(userQueryObject queryObject.UserQueryObject, eventId string) interface{} // TODO: AttachmentCollection
	AddAttachment(userQueryObject queryObject.UserQueryObject, attachment interface{}) interface{} // TODO: Attachment (param and return)
	CreateOpenExtension(userQueryObject queryObject.UserQueryObject, resourceId string) interface{} // TODO: OpenTypeExtension
	GetOpenExtension(userQueryObject queryObject.UserQueryObject, resourceId string, extensionId string) interface{} // TODO: OpenTypeExtensionCollection
	CreateSingleValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, singleValueLegacyExtendedPropertyCollection interface{}) Event // TODO: SingleValueLegacyExtendedPropertyCollection
	GetEventWithSingleValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, extendedPropertyId string, propertyValue string) Event
	CreateMultiValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, multiValueLegacyExtendedPropertyCollection interface{}) Event // TODO: MultiValueLegacyExtendedPropertyCollection
	GetEventWithMultiValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, extendedPropertyId string) Event
}
