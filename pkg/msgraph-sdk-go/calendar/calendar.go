package calendar

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
