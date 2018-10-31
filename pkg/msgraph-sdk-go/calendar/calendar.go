package calendar

import "time"

type Calendar struct {
	CanEdit bool `json:"canEdit"`
	CanShare bool `json:"canShare"`
	CanViewPrivateItems bool `json:	"canViewPrivateItems"`
	ChangeKey string `json:"changeKey"`
	Color string `json:"color"`
	HexColor string `json:"hexColor"`
	Id string `json:"id"`
	IsDefaultCalendar string `json:"isDefaultCalendar"`
	IsShared bool `json:"isShared"`
	IsSharedWithMe bool `json:"isSharedWithMe"`
	Name string `json:"name"`
	Owner interface{} `json:"owner"` // TODO: Owner emailAddress
	CalendarView interface{} `json:"calendarView",omitempty` // TODO: CalendarView EventCollection
	Events interface{} `json:"events",omitempty` // TODO: Events EventCollection
	MultiValueLegacyExtendedProperties interface{} `json:"multiValueLegacyExtendedProperties",omitempty` // TODO: MultiValueLegacyExtendedPropertiesCollection
	SingleValueLegacyExtendedProperties interface {} `json:"singleValueLegacyExtendedProperties",omitempty` // TODO: SingleValueLegacyExtendedPropertiesCollection
}

type ICalendar interface {
	ListCalendar() interface{} // TODO: CalendarCollection struct
	CreateCalendar(calendar Calendar) Calendar
	GetCalendar() Calendar
	Update(color string, isDefaultCalendar bool, name string) Calendar
	Delete()
	ListCalendarView(startTime time.Time, endTime time.Time) interface{} // TODO: EventCollection struct
	ListEvents()
	CreateEvent() interface{} // TODO: Event struct
	FindMeetingTimes() interface{} // TODO: MeetingTimeSuggestionsResult struct
	GetSchedule() interface{} // TODO: ScheduleInformationCollection struct and ScheduleInformation struct
	CreateSingleValueExtendedProperty() Calendar
	GetCalendarWithSingleValueExtendedProperty() Calendar
	CreateMultiValueExtendedProperty() Calendar
	GetCalendarWithMultiValueExtendedProperty() Calendar
}