package models

import (
	"regexp"
)

var (
	beelineReg, _ = regexp.Compile("(996|\\+996|0)?(77\\d{7}|31258\\d{4}|22[0-2]\\d{6})")
	megacom, _    = regexp.Compile("(996|\\+996|0)?(55\\d{7}|755\\d{6})")
	nurtelecom, _ = regexp.Compile("(996|\\+996|0)?([5,7]0\\d{7})")
)

// PhoneType la
type PhoneType int

const (
	// Megacom la
	Megacom PhoneType = iota

	// Beeline la
	Beeline

	// Nurtelecom la
	Nurtelecom

	// Unknown la
	Unknown
)

func (phone PhoneType) String() string {
	names := [...]string{"Megacom", "Beeline", "Nurtelecom"}
	if phone < Megacom || phone > Nurtelecom {
		return "Unknown"
	}
	return names[phone]
}

// CheckPhoneOperator la
func CheckPhoneOperator(phone string) PhoneType {
	if len(phone) != 12 {
		return Unknown
	}
	if beelineReg.MatchString(phone) {
		return Beeline
	} else if megacom.MatchString(phone) {
		return Megacom
	} else if nurtelecom.MatchString(phone) {
		return Nurtelecom
	} else {
		return Unknown
	}
}
