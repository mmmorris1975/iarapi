package iarapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

type LoginRequest struct {
	MemberLogin     bool   `json:"memberLogin"`
	Agency          string `json:"agencyName"`
	User            string `json:"memberfname"`
	Password        string `json:"memberpwd"`
	UrlTo           string `json:"urlTo"`
	RememberMe      bool   `json:"rememberPwd"`
	OverrideSession bool   `json:"overrideSession"`
}

type LoginReply struct {
	Message string `json:"d"`
}

type SubscriberInfo struct {
	Id              int    `json:"subscriberId"`
	StatusID        int    `json:"statusID"`
	Name            string `json:"name"`
	TimeZone        int    `json:"timeZone"`
	AssignedPhone   string `json:"assignedPhone"`
	AdditionalPhone string `json:"additionalPhone"`
	Location        struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
	EnableEmailInput         bool   `json:"enableEmailInput"`
	LogoImage                string `json:"logoImage"`
	ScreenName               string `json:"screenName"`
	Country                  string `json:"country"`
	CountryCode              string `json:"countryCode"`
	TelephoneKeyEntries1     string `json:"telephoneKeyEntries1"`
	TelephoneKeyEntries2     string `json:"telephoneKeyEntries2"`
	TelephoneKeyEntries3     string `json:"telephoneKeyEntries3"`
	TelephoneKeyEntries4     string `json:"telephoneKeyEntries4"`
	TelephoneKeyEntries5     string `json:"telephoneKeyEntries5"`
	TelephoneKeyEntries6     string `json:"telephoneKeyEntries6"`
	TelephoneKeyEntries7     string `json:"telephoneKeyEntries7"`
	TelephoneKeyEntries8     string `json:"telephoneKeyEntries8"`
	TelephoneKeyEntries9     string `json:"telephoneKeyEntries9"`
	TelephoneKeyEntriesDef   string `json:"telephoneKeyEntriesDef"`
	AutoClear                bool   `json:"autoClear"`
	MinutesToClearEtaExpired int    `json:"minutesToClearEtaExpired"`
	MaxTimeInToggleView      int    `json:"maxTimeInToggleView"`
	City                     string `json:"city"`
	State                    string `json:"state"`
	EnableDigitalDashboard   bool   `json:"enableDigitalDashboard"`
	CurrentDate              string `json:"currentDate"`
	CurrentTime              string `json:"currentTime"`
	ToggleInDashboard        bool   `json:"toggleInDashboard"`
	TimeZoneId               int    `json:"timeZoneId"`
	IsAffectedByDstChange    bool   `json:"isAffectedByDstChange"`
	OldTimeZoneId            int    `json:"oldTimeZoneId"`
	OldIsAffectedByDstChange bool   `json:"oldIsAffectedByDstChange"`
	EnableLegacyDashboard    bool   `json:"enableLegacyDashboard"`
	Emailaddr                string `json:"emailaddr"`
	IsActive                 bool   `json:"isActive"`
	NameForDispatcherUse     string `json:"nameForDispatcherUse"`
	TtdToggleInDashboard     bool   `json:"ttdToggleInDashboard"`
	AllowSpecialShifts       bool   `json:"allowSpecialShifts"`
}

type MemberInfo struct {
	Id                                 int         `json:"memberId"`
	SubscriberId                       int         `json:"subscriberId"`
	FirstName                          string      `json:"firstName"`
	LastName                           string      `json:"lastName"`
	ProfileImage                       string      `json:"profileImage"`
	MaxTimeEmergencyMode               int         `json:"maxTimeEmergencyMode"`
	ClearNow                           bool        `json:"clearNow"`
	ColorBorder                        int         `json:"colorBorder"`
	CanEditOwnSchedule                 bool        `json:"canEditOwnSchedule"`
	CanEditAllSchedules                bool        `json:"canEditAllSchedules"`
	AllowOwnPCFScheduling              bool        `json:"allowOwnPCFScheduling"`
	AllowOwnCFScheduling               bool        `json:"allowOwnCFScheduling"`
	CanManageEvents                    bool        `json:"canManageEvents"`
	AllowManageHydrants                bool        `json:"allowManageHydrants"`
	AllowDeleteHydrants                bool        `json:"allowDeleteHydrants"`
	AllowManageMarkers                 bool        `json:"allowManageMarkers"`
	AllowDeleteMarkers                 bool        `json:"allowDeleteMarkers"`
	PermittedToVerifyIncidentAddresses bool        `json:"permittedToVerifyIncidentAddresses"`
	PermittedToCreateGeofence          bool        `json:"permittedToCreateGeofence"`
	DefaultRespondNow                  string      `json:"defaultRespondNow"`
	Position                           string      `json:"position"`
	ReminderShifts                     string      `json:"reminderShifts"`
	PositionId                         int         `json:"positionId"`
	AllowToggleEmergencyDD             bool        `json:"allowToggleEmergencyDD"`
	AllowEditOwnProfile                bool        `json:"allowEditOwnProfile"`
	PermittedChangePage6               bool        `json:"permittedChangePage6"`
	AllowEditScrollMessage             bool        `json:"allowEditScrollMessage"`
	DefaultLocation                    string      `json:"defaultLocation"`
	DefaultCategory                    int         `json:"defaultCategory"`
	DefaultShiftDuration               int         `json:"defaultShiftDuration"`
	DoNotDisturb                       int         `json:"doNotDisturb"`
	DoNotDisturbStartTime              string      `json:"doNotDisturbStartTime"`
	DoNotDisturbFinishTime             string      `json:"doNotDisturbFinishTime"`
	TextMessageAddressDND              bool        `json:"textMessageAddressDND"`
	AppPushNotificationsDND            bool        `json:"appPushNotificationsDND"`
	DoNotDisturbDevice                 bool        `json:"doNotDisturbDevice"`
	DeviceToken                        interface{} `json:"deviceToken"`
	DeviceActive                       bool        `json:"deviceActive"`
	MemberEmail                        string      `json:"memberEmail"`
	SecondaryEmail                     string      `json:"secondaryEmail"`
	TextMemberAddress                  string      `json:"textMemberAddress"`
}

type Incident struct {
	Id                      int         `json:"id"`
	ArrivedOn               string      `json:"arrivedOn"`
	MessageBody             string      `json:"messageBody"`
	DestinationEmailAddress string      `json:"destinationEmailAddress"`
	OriginationEmailAddress string      `json:"originationEmailAddress"`
	SubscriberId            int         `json:"subscriberId"`
	VerifiedAddressStatus   int         `json:"verifiedAddressStatus"`
	ArrivedOnString         string      `json:"arrivedOnString"`
	Index                   int         `json:"index"`
	Address                 string      `json:"address"`
	Location                string      `json:"location"`
	Direction               interface{} `json:"direction"`
	VerifiedStreetNumber    string      `json:"verifiedStreetNumber"`
	VerifiedStreetName      string      `json:"verifiedStreetName"`
	VerifiedCity            string      `json:"verifiedCity"`
	VerifiedState           string      `json:"verifiedState"`
	VerifiedCountry         string      `json:"verifiedCountry"`
	LongDirection           string      `json:"longDirection"`
	HasCoordinatesInBoddy   bool        `json:"hasCoordinatesInBoddy"`
	IsVerifiedAndActive     bool        `json:"isVerifiedAndActive"`
	AddedBy                 interface{} `json:"addedBy"`
	AddedOn                 string      `json:"addedOn"`
	LastUpdatedBy           interface{} `json:"lastUpdatedBy"`
	UpdatedOn               string      `json:"updatedOn"`
	TimeZoneId              int         `json:"timeZoneId"`
	IsDst                   bool        `json:"isDst"`
}
type IncidentList []*Incident

type Message struct {
	Id           string    `json:"id"`
	MessageId    int       `json:"messageId"`
	SubscriberId int       `json:"subscriberId"`
	Message      string    `json:"message"`
	CreatedDate  time.Time `json:"createdDate"`
}
type MessageList []*Message

type Dispatcher struct {
	CentralizedId  int    `json:"centralizedId"`
	DispatcherId   int    `json:"dispatcherId"`
	DispatcherName string `json:"dispatcherName"`
}

type Dispatchers struct {
	SubscriberId int           `json:"subscriberId"`
	Dispatchers  []*Dispatcher `json:"dispatchers"`
}

type ResponderCode struct {
	Id                int    `json:"id"`
	SubscriberId      int    `json:"subscriberId"`
	KeyEntry          string `json:"keyEntry"`
	StatusForTracking bool   `json:"statusForTracking"`
	IsTelephoneKey    bool   `json:"isTelephoneKey"`
	IsDefaultKey      bool   `json:"isDefaultKey"`
	CustomSortOrder   int    `json:"customSortOrder"`
}

type ResponderCodes struct {
	ResponseCodes []*ResponderCode `json:"responseCodes"`
	TelephoneKeys []*ResponderCode `json:"telephoneKeys"`
}

type Responder struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Position       string    `json:"position"`
	RespondingTo   string    `json:"respondingTo"`
	CalledAt       time.Time `json:"calledAt"`
	EtaBefore      time.Time `json:"etaBefore"`
	LastName       string    `json:"lastName"`
	IsMutualAid    bool      `json:"isMutualAid"`
	SubscriberName string    `json:"subscriberName"`
	SubscriberId   int       `json:"subscriberId"`
	Order          int       `json:"order"`
	ImageUrl       string    `json:"imageUrl"`
	ColorBorder    int       `json:"colorBorder"`
	MemberId       int       `json:"memberId"`
	Expired        time.Time `json:"expired"`
	TimeZoneId     int       `json:"timeZoneId"`
	IsDst          bool      `json:"isDst"`
	ResponseCodeId int       `json:"responseCodeId"`
}
type ResponderList []*Responder

// TODO - we don't know what this looks like
type Apparatus struct{}
type ApparatusList []*Apparatus

type OnDutyAtCode struct {
	Id           string `json:"id"`
	SubscriberId int    `json:"subscriberId"`
	KeyEntry     string `json:"keyEntry"`
}
type OnDutyAtCodeList []*OnDutyAtCode

type IncidentSearchRequest struct {
	StartTime time.Time
	EndTime   time.Time
	PageSize  int
	Page      int
}

func (r IncidentSearchRequest) MarshalJSON() ([]byte, error) {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.PageSize < 1 {
		r.PageSize = 100
	}

	sYr, sMo, sDa := r.StartTime.Date()
	eYr, eMo, eDa := r.EndTime.Date()
	dateFmt := "%04d-%02d-%02d"

	m := map[string]interface{}{
		"startDate": fmt.Sprintf(dateFmt, sYr, sMo, sDa),
		"endDate":   fmt.Sprintf(dateFmt, eYr, eMo, eDa),
		"loading":   false,
		"submit":    true,
		"page":      r.Page,
		"pageSize":  r.PageSize,
	}

	return json.Marshal(m)
}
