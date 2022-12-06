package walletpass

import "time"

type BarcodeFormat string
type DataDetectorType string
type DateStyle string
type NumberStyle string
type TextAlignment string
type TimeStyle string
type TransitType string

const (
	BarcodeFormatQR      BarcodeFormat = "PKBarcodeFormatQR"
	BarcodeFormatPDF417  BarcodeFormat = "PKBarcodeFormatPDF417"
	BarcodeFormatAztec   BarcodeFormat = "PKBarcodeFormatAztec"
	BarcodeFormatCode128 BarcodeFormat = "PKBarcodeFormatCode128"

	DataDetectorTypePhoneNumber   DataDetectorType = "PKDataDetectorTypePhoneNumber"
	DataDetectorTypeLink          DataDetectorType = "PKDataDetectorTypeLink"
	DataDetectorTypeAddress       DataDetectorType = "PKDataDetectorTypeAddress"
	DataDetectorTypeCalendarEvent DataDetectorType = "PKDataDetectorTypeCalendarEvent"

	DateStyleNone   DateStyle = "PKDateStyleNone"
	DateStyleShort  DateStyle = "PKDateStyleShort"
	DateStyleMedium DateStyle = "PKDateStyleMedium"
	DateStyleLong   DateStyle = "PKDateStyleLong"
	DateStyleFull   DateStyle = "PKDateStyleFull"

	NumberStyleDecimal    NumberStyle = "PKNumberStyleDecimal"
	NumberStylePercent    NumberStyle = "PKNumberStylePercent"
	NumberStyleScientific NumberStyle = "PKNumberStyleScientific"
	NumberStyleSpellOut   NumberStyle = "PKNumberStyleSpellOut"

	TextAlignmentLeft    TextAlignment = "PKTextAlignmentLeft"
	TextAlignmentCenter  TextAlignment = "PKTextAlignmentCenter"
	TextAlignmentRight   TextAlignment = "PKTextAlignmentRight"
	TextAlignmentNatural TextAlignment = "PKTextAlignmentNatural"

	TimeStyleNone   TimeStyle = "PKTimeStyleNone"
	TimeStyleShort  TimeStyle = "PKTimeStyleShort"
	TimeStyleMedium TimeStyle = "PKTimeStyleMedium"
	TimeStyleLong   TimeStyle = "PKTimeStyleLong"
	TimeStyleFull   TimeStyle = "PKTimeStyleFull"

	TransitTypeAir     TransitType = "PKTransitTypeAir"
	TransitTypeBoat    TransitType = "PKTransitTypeBoat"
	TransitTypeBus     TransitType = "PKTransitTypeBus"
	TransitTypeGeneric TransitType = "PKTransitTypeGeneric"
	TransitTypeTrain   TransitType = "PKTransitTypeTrain"
)

type Pass struct {
	AppLaunchURL               string        `json:"appLaunchURL,omitempty"`
	AssociatedStoreIdentifiers []int64       `json:"associatedStoreIdentifiers,omitempty"`
	AuthenticationToken        string        `json:"authenticationToken,omitempty"`
	BackgroundColor            string        `json:"backgroundColor,omitempty"`
	Barcodes                   []Barcodes    `json:"barcodes,omitempty"`
	Beacons                    []Beacons     `json:"beacons,omitempty"`
	BoardingPass               *BoardingPass `json:"boardingPass,omitempty"`
	Coupon                     *Coupon       `json:"coupon,omitempty"`
	Description                string        `json:"description,omitempty"`
	EventTicket                *EventTicket  `json:"eventTicket,omitempty"`
	ExpirationDate             *time.Time    `json:"expirationDate,omitempty"`
	ForegroundColor            string        `json:"foregroundColor,omitempty"`
	FormatVersion              int64         `json:"formatVersion"`
	Generic                    *Generic      `json:"generic,omitempty"`
	GroupingIdentifier         string        `json:"groupingIdentifier,omitempty"`
	LabelColor                 string        `json:"labelColor,omitempty"`
	Locations                  []Locations   `json:"locations,omitempty"`
	LogoText                   string        `json:"logoText,omitempty"`
	MaxDistance                int64         `json:"maxDistance,omitempty"`
	NFC                        *NFC          `json:"nfc,omitempty"`
	OrganizationName           string        `json:"organizationName"`
	PassTypeIdentifier         string        `json:"passTypeIdentifier"`
	RelevantDate               *time.Time    `json:"relevantDate,omitempty"`
	Semantics                  *SemanticTags `json:"semantics,omitempty"`
	SerialNumber               string        `json:"serialNumber"`
	SharingProhibited          bool          `json:"sharingProhibited,omitempty"`
	StoreCard                  *StoreCard    `json:"storeCard,omitempty"`
	SuppressStripShine         bool          `json:"suppressStripShine,omitempty"`
	TeamIdentifier             string        `json:"teamIdentifier"`
	UserInfo                   interface{}   `json:"userInfo,omitempty"`
	Voided                     bool          `json:"voided,omitempty"`
	WebServiceURL              string        `json:"webServiceURL,omitempty"`
}

type Barcodes struct {
	AltText         string        `json:"altText,omitempty"`
	Format          BarcodeFormat `json:"format,omitempty"`
	Message         string        `json:"message,omitempty"`
	MessageEncoding string        `json:"messageEncoding,omitempty"`
}

type Beacons struct {
	Major         uint16 `json:"major,omitempty"`
	Minor         uint16 `json:"minor,omitempty"`
	ProximityUUID string `json:"proximityUUID,omitempty"`
	RelevantText  string `json:"relevantText,omitempty"`
}

type BoardingPass struct {
	*PassFields
	TransitType TransitType `json:"transitType,omitempty"`
}

type Coupon struct {
	*PassFields
}

type EventTicket struct {
	*PassFields
}

type Generic struct {
	*PassFields
}

type PassFields struct {
	AuxiliaryFields []PassFieldContent `json:"auxiliaryFields,omitempty"`
	BackFields      []PassFieldContent `json:"backFields,omitempty"`
	HeaderFields    []PassFieldContent `json:"headerFields,omitempty"`
	PrimaryFields   []PassFieldContent `json:"primaryFields,omitempty"`
	SecondaryFields []PassFieldContent `json:"secondaryFields,omitempty"`
}

type PassFieldContent struct {
	AttributedValue   string             `json:"attributedValue,omitempty"`
	ChangeMessage     string             `json:"changeMessage,omitempty"`
	CurrencyCode      string             `json:"currencyCode,omitempty"`
	DataDetectorTypes []DataDetectorType `json:"dataDetectorTypes,omitempty"`
	DateStyle         DateStyle          `json:"dateStyle,omitempty"`
	IgnoresTimeZone   bool               `json:"ignoresTimeZone,omitempty"`
	IsRelative        bool               `json:"isRelative,omitempty"`
	Key               string             `json:"key,omitempty"`
	Label             string             `json:"label,omitempty"`
	NumberStyle       NumberStyle        `json:"numberStyle,omitempty"`
	TextAlignment     TextAlignment      `json:"textAlignment,omitempty"`
	TimeStyle         TimeStyle          `json:"timeStyle,omitempty"`
	Value             string             `json:"value,omitempty"`
}

type Locations struct {
	Altitude     float64 `json:"altitude,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	RelevantText string  `json:"relevantText,omitempty"`
}

type NFC struct {
	EncryptionPublicKey    string `json:"encryptionPublicKey,omitempty"`
	Message                string `json:"message,omitempty"`
	RequiresAuthentication bool   `json:"requiresAuthentication,omitempty"`
}

type SemanticTags struct {
	AirlineCode                    string                `json:"airlineCode,omitempty"`
	ArtistIDs                      string                `json:"artistIDs,omitempty"`
	AwayTeamAbbreviation           string                `json:"awayTeamAbbreviation,omitempty"`
	AwayTeamLocation               string                `json:"awayTeamLocation,omitempty"`
	AwayTeamName                   string                `json:"awayTeamName,omitempty"`
	Balance                        *CurrencyAmount       `json:"balance,omitempty"`
	BoardingGroup                  string                `json:"boardingGroup,omitempty"`
	BoardingSequenceNumber         string                `json:"boardingSequenceNumber,omitempty"`
	CarNumber                      string                `json:"carNumber,omitempty"`
	ConfirmationNumber             string                `json:"confirmationNumber,omitempty"`
	CurrentArrivalDate             *time.Time            `json:"currentArrivalDate,omitempty"`
	CurrentBoardingDate            *time.Time            `json:"currentBoardingDate,omitempty"`
	CurrentDepartureDate           *time.Time            `json:"currentDepartureDate,omitempty"`
	DepartureAirportCode           string                `json:"departureAirportCode,omitempty"`
	DepartureAirportName           string                `json:"departureAirportName,omitempty"`
	DepartureGate                  string                `json:"departureGate,omitempty"`
	DepartureLocation              *Location             `json:"departureLocation,omitempty"`
	DepartureLocationDescription   string                `json:"departureLocationDescription,omitempty"`
	DeparturePlatform              string                `json:"departurePlatform,omitempty"`
	DepartureStationName           string                `json:"departureStationName,omitempty"`
	DepartureTerminal              string                `json:"departureTerminal,omitempty"`
	DestinationAirportCode         string                `json:"destinationAirportCode,omitempty"`
	DestinationAirportName         string                `json:"destinationAirportName,omitempty"`
	DestinationGate                string                `json:"destinationGate,omitempty"`
	DestinationLocation            *Location             `json:"destinationLocation,omitempty"`
	DestinationLocationDescription string                `json:"destinationLocationDescription,omitempty"`
	DestinationPlatform            string                `json:"destinationPlatform,omitempty"`
	DestinationStationName         string                `json:"destinationStationName,omitempty"`
	DestinationTerminal            string                `json:"destinationTerminal,omitempty"`
	Duration                       int64                 `json:"duration,omitempty"`
	EventEndDate                   *time.Time            `json:"eventEndDate,omitempty"`
	EventName                      string                `json:"eventName,omitempty"`
	EventStartDate                 *time.Time            `json:"eventStartDate,omitempty"`
	EventType                      string                `json:"eventType,omitempty"`
	FlightCode                     string                `json:"flightCode,omitempty"`
	FlightNumber                   int64                 `json:"flightNumber,omitempty"`
	Genre                          string                `json:"genre,omitempty"`
	HomeTeamAbbreviation           string                `json:"homeTeamAbbreviation,omitempty"`
	HomeTeamLocation               string                `json:"homeTeamLocation,omitempty"`
	HomeTeamName                   string                `json:"homeTeamName,omitempty"`
	LeagueAbbreviation             string                `json:"leagueAbbreviation,omitempty"`
	LeagueName                     string                `json:"leagueName,omitempty"`
	MembershipProgramName          string                `json:"membershipProgramName,omitempty"`
	MembershipProgramNumber        string                `json:"membershipProgramNumber,omitempty"`
	OriginalArrivalDate            *time.Time            `json:"originalArrivalDate,omitempty"`
	OriginalBoardingDate           *time.Time            `json:"originalBoardingDate,omitempty"`
	OriginalDepartureDate          *time.Time            `json:"originalDepartureDate,omitempty"`
	PassengerName                  *PersonNameComponents `json:"passengerName,omitempty"`
	PerformerNames                 []string              `json:"performerNames,omitempty"`
	PriorityStatus                 string                `json:"priorityStatus,omitempty"`
	Seats                          *Seat                 `json:"seats,omitempty"`
	SecurityScreening              string                `json:"securityScreening,omitempty"`
	SilenceRequested               bool                  `json:"silenceRequested,omitempty"`
	SportName                      string                `json:"sportName,omitempty"`
	TotalPrice                     *CurrencyAmount       `json:"totalPrice,omitempty"`
	TransitProvider                string                `json:"transitProvider,omitempty"`
	TransitStatus                  string                `json:"transitStatus,omitempty"`
	TransitStatusReason            string                `json:"transitStatusReason,omitempty"`
	VehicleName                    string                `json:"vehicleName,omitempty"`
	VehicleNumber                  string                `json:"vehicleNumber,omitempty"`
	VehicleType                    string                `json:"vehicleType,omitempty"`
	VenueEntrance                  string                `json:"venueEntrance,omitempty"`
	VenueLocation                  *Location             `json:"venueLocation,omitempty"`
	VenueName                      string                `json:"venueName,omitempty"`
	VenuePhoneNumber               string                `json:"venuePhoneNumber,omitempty"`
	VenueRoom                      string                `json:"venueRoom,omitempty"`
	WifiAccess                     *WifiNetwork          `json:"wifiAccess,omitempty"`
}

type CurrencyAmount struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currencyCode,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type PersonNameComponents struct {
	FamilyName             string `json:"familyName,omitempty"`
	GivenName              string `json:"givenName,omitempty"`
	MiddleName             string `json:"middleName,omitempty"`
	NamePrefix             string `json:"namePrefix,omitempty"`
	NameSuffix             string `json:"nameSuffix,omitempty"`
	Nickname               string `json:"nickname,omitempty"`
	PhoneticRepresentation string `json:"phoneticRepresentation,omitempty"`
}

type Seat struct {
	SeatDescription string `json:"seatDescription,omitempty"`
	SeatIdentifier  string `json:"seatIdentifier,omitempty"`
	SeatNumber      string `json:"seatNumber,omitempty"`
	SeatRow         string `json:"seatRow,omitempty"`
	SeatSection     string `json:"seatSection,omitempty"`
	SeatType        string `json:"seatType,omitempty"`
}

type WifiNetwork struct {
	Password string `json:"password,omitempty"`
	Ssid     string `json:"ssid,omitempty"`
}

type StoreCard struct {
	*PassFields
}
