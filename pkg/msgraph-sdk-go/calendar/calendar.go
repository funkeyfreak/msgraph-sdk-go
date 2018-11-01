package calendar

import (
	"../queryObject"
	"time"
)

// Calendar object
type Calendar struct {
	CanEdit bool `json:"canEdit"`
	CanShare bool `json:"canShare"`
	CanViewPrivateItems bool `json:"canViewPrivateItems"`
	ChangeKey string `json:"changeKey"`
	Color string `json:"color"`
	HexColor string `json:"hexColor"`
	Id string `json:"id"`
	IsDefaultCalendar string `json:"isDefaultCalendar"`
	IsShared bool `json:"isShared"`
	IsSharedWithMe bool `json:"isSharedWithMe"`
	Name string `json:"name"`
	Owner interface{} `json:"owner"` // TODO: Owner emailAddress
	CalendarView interface{} `json:"calendarView,omitempty"` // TODO: CalendarView EventCollection
	Events interface{} `json:"events,omitempty"` // TODO: Events EventCollection
	MultiValueLegacyExtendedProperties interface{} `json:"multiValueLegacyExtendedProperties,omitempty"` // TODO: MultiValueLegacyExtendedPropertiesCollection
	SingleValueLegacyExtendedProperties interface {} `json:"singleValueLegacyExtendedProperties,omitempty"` // TODO: SingleValueLegacyExtendedPropertiesCollection
}

type ICalendar interface {
	ListCalendar() interface{} // TODO: CalendarCollection struct
	CreateCalendar(calendar Calendar) Calendar
	GetCalendar(userQueryObject queryObject.UserQueryObject, eventId string) Calendar
	Update(userQueryObject queryObject.UserQueryObject, eventId string, color string, isDefaultCalendar bool, name string) Calendar
	Delete(userQueryObject queryObject.UserQueryObject, eventId string)
	ListCalendarView(startTime time.Time, endTime time.Time) interface{} // TODO: EventCollection struct
	ListEvents()
	CreateEvent() interface{} // TODO: Event struct
	FindMeetingTimes() interface{} // TODO: MeetingTimeSuggestionsResult struct
	GetSchedule() interface{} // TODO: ScheduleInformationCollection struct and ScheduleInformation struct
	CreateSingleValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, singleValueLegacyExtendedPropertyCollection interface{}) Calendar // TODO: singleValueLegacyExtendedPropertyCollection
	GetCalendarWithSingleValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, extendedPropertyId string, propertyValue string) Calendar
	CreateMultiValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, multiValueLegacyExtendedPropertyCollection interface{}) Calendar // TODO: MultiValueLegacyExtendedPropertyCollection
	GetCalendarWithMultiValueExtendedProperty(userQueryObject queryObject.UserQueryObject, resourceId string, extendedPropertyId string) Calendar
}