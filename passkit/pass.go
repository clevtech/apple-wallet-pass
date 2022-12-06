package passkit

import (
	"encoding/json"
	"errors"
	"time"
)

type BarcodeFormat string
type DataDetectorType string
type DateStyle string
type NumberStyle string
type PassPersonalizationField string
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

	PassPersonalizationFieldName         PassPersonalizationField = "PKPassPersonalizationFieldName"
	PassPersonalizationFieldPostalCode   PassPersonalizationField = "PKPassPersonalizationFieldPostalCode"
	PassPersonalizationFieldEmailAddress PassPersonalizationField = "PKPassPersonalizationFieldEmailAddress"
	PassPersonalizationFieldPhoneNumber  PassPersonalizationField = "PKPassPersonalizationFieldPhoneNumber"

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

func (p *Pass) SetAppLaunchURL(url string) error {
	if url == "" {
		return errors.New("URL can not be empty")
	}

	p.AppLaunchURL = url

	return nil
}

func (p *Pass) SetAssociatedStoreIdentifiers(ids []int64) error {
	if len(ids) == 0 {
		return errors.New("Store identifiers can not be empty")
	}

	p.AssociatedStoreIdentifiers = ids

	return nil
}

func (p *Pass) SetAuthenticationToken(token string) error {
	if token == "" {
		return errors.New("Authentication token can not be empty")
	}

	p.AuthenticationToken = token

	return nil
}

func (p *Pass) SetBackgroundColor(color string) error {
	if color == "" {
		return errors.New("Background color can not be empty")
	}

	p.BackgroundColor = color

	return nil
}

func (p *Pass) SetBarcodes(barcodes []Barcodes) error {
	if len(barcodes) == 0 {
		return errors.New("Barcodes can not be empty")
	}

	p.Barcodes = barcodes

	return nil
}

func (p *Pass) SetBeacons(beacons []Beacons) error {
	if len(beacons) == 0 {
		return errors.New("Beacons can not be empty")
	}

	p.Beacons = beacons

	return nil
}

func (p *Pass) SetBoardingPass(boardingPass *BoardingPass) error {
	if boardingPass == nil {
		return errors.New("Boarding pass can not be empty")
	}

	p.BoardingPass = boardingPass

	return nil
}

func (p *Pass) SetCoupon(coupon *Coupon) error {
	if coupon == nil {
		return errors.New("Coupon can not be empty")
	}

	p.Coupon = coupon

	return nil
}

func (p *Pass) SetDescription(description string) error {
	if description == "" {
		return errors.New("Description can not be empty")
	}

	p.Description = description

	return nil
}

func (p *Pass) SetEventTicket(eventTicket *EventTicket) error {
	if eventTicket == nil {
		return errors.New("Event ticket can not be empty")
	}

	p.EventTicket = eventTicket

	return nil
}

func (p *Pass) SetExpirationDate(date *time.Time) error {
	if date == nil {
		return errors.New("Expiration date can not be empty")
	}

	p.ExpirationDate = date

	return nil
}

func (p *Pass) SetForegroundColor(color string) error {
	if color == "" {
		return errors.New("Foreground color can not be empty")
	}

	p.ForegroundColor = color

	return nil
}

func (p *Pass) SetFormatVersion(version int64) error {
	if version != 1 {
		return errors.New("Format version can only be 1")
	}

	p.FormatVersion = version

	return nil
}

func (p *Pass) SetGeneric(generic *Generic) error {
	if generic == nil {
		return errors.New("Generic can not be empty")
	}

	p.Generic = generic

	return nil
}

func (p *Pass) SetGroupingIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Grouping identifier can not be empty")
	}

	p.GroupingIdentifier = identifier

	return nil
}

func (p *Pass) SetLabelColor(color string) error {
	if color == "" {
		return errors.New("Label color can not be empty")
	}

	p.LabelColor = color

	return nil
}

func (p *Pass) SetLocations(locations []Locations) error {
	if len(locations) == 0 {
		return errors.New("Locations can not be empty")
	}

	p.Locations = locations

	return nil
}

func (p *Pass) SetLogoText(text string) error {
	if text == "" {
		return errors.New("Logo text can not be empty")
	}

	p.LogoText = text

	return nil
}

func (p *Pass) SetMaxDistance(distance int64) error {
	if distance == 0 {
		return errors.New("Max distance can not be empty")
	}

	p.MaxDistance = distance

	return nil
}

func (p *Pass) SetNFC(nfc *NFC) error {
	if nfc == nil {
		return errors.New("NFC can not be empty")
	}

	p.NFC = nfc

	return nil
}

func (p *Pass) SetOrganizationName(name string) error {
	if name == "" {
		return errors.New("Organization name can not be empty")
	}

	p.OrganizationName = name

	return nil
}

func (p *Pass) SetPassTypeIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Pass type identifier can not be empty")
	}

	p.PassTypeIdentifier = identifier

	return nil
}

func (p *Pass) SetRelevantDate(date *time.Time) error {
	if date == nil {
		return errors.New("Relevant date can not be empty")
	}

	p.RelevantDate = date

	return nil
}

func (p *Pass) SetSemantics(semantics *SemanticTags) error {
	if semantics == nil {
		return errors.New("Semantics can not be empty")
	}

	p.Semantics = semantics

	return nil
}

func (p *Pass) SetSerialNumber(number string) error {
	if number == "" {
		return errors.New("Serial number can not be empty")
	}

	p.SerialNumber = number

	return nil
}

func (p *Pass) SetSharingProhibited(prohibited bool) error {
	p.SharingProhibited = prohibited

	return nil
}

func (p *Pass) SetStoreCard(storeCard *StoreCard) error {
	if storeCard == nil {
		return errors.New("Store card can not be empty")
	}

	p.StoreCard = storeCard

	return nil
}

func (p *Pass) SetSuppressStripShine(suppress bool) error {
	p.SuppressStripShine = suppress

	return nil
}

func (p *Pass) SetTeamIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Team identifier can not be empty")
	}

	p.TeamIdentifier = identifier

	return nil
}

// func (p *Pass) SetUserInfo(userInfo []UserInfo) error {
// 	if len(userInfo) == 0 {
// 		return errors.New("User infos can not be empty")
// 	}

// 	p.UserInfo = userInfo

// 	return nil
// }

func (p *Pass) SetVoided(voided bool) error {
	p.Voided = voided

	return nil
}

func (p *Pass) SetWebServiceURL(url string) error {
	if url == "" {
		return errors.New("Web service url can not be empty")
	}

	p.WebServiceURL = url

	return nil
}

func (p *Pass) ToJson() ([]byte, error) {
	return json.Marshal(p)
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

func NewBoardingPass(transitType TransitType) *BoardingPass {
	return &BoardingPass{PassFields: NewPassFields(), TransitType: transitType}
}

type Coupon struct {
	*PassFields
}

func NewCoupon() *Coupon {
	return &Coupon{PassFields: NewPassFields()}
}

type EventTicket struct {
	*PassFields
}

func NewEventTicket() *EventTicket {
	return &EventTicket{PassFields: NewPassFields()}
}

type Generic struct {
	*PassFields
}

func NewGeneric() *Generic {
	return &Generic{PassFields: NewPassFields()}
}

type PassFields struct {
	AuxiliaryFields []PassFieldContent `json:"auxiliaryFields,omitempty"`
	BackFields      []PassFieldContent `json:"backFields,omitempty"`
	HeaderFields    []PassFieldContent `json:"headerFields,omitempty"`
	PrimaryFields   []PassFieldContent `json:"primaryFields,omitempty"`
	SecondaryFields []PassFieldContent `json:"secondaryFields,omitempty"`
}

func NewPassFields() *PassFields {
	return &PassFields{}
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

func NewStoreCard() *StoreCard {
	return &StoreCard{PassFields: NewPassFields()}
}

type Personalize struct {
	Description                   string                     `json:"description,omitempty"`
	RequiredPersonalizationFields []PassPersonalizationField `json:"requiredPersonalizationFields,omitempty"`
	TermsAndConditions            string                     `json:"termsAndConditions,omitempty"`
}
