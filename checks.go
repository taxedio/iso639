package iso639

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Removes Diacritics from submitted string
func convertDia(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)
	return s
}

// Check map for Alpha3 Match and return ISO
// example, "en" will return "ENG"
func Alpha3Match(s string) *string {
	var (
		returnStr string
	)
	if val, ok := isocodes3[s]; ok {
		returnStr = fmt.Sprintf("%v", val.alph3b)
		return &returnStr
	}
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "xx" || s == "xxx" {
		return nil
	}
	for _, val := range isocodes3 {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3b) || s == strings.ToLower(val.alph3t) {
			returnStr = fmt.Sprintf("%v", val.alph3b)
			return &returnStr
		}
		// check against english characters
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			returnStr = fmt.Sprintf("%v", val.alph3b)
			return &returnStr
		}
		// check against original text
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			returnStr = fmt.Sprintf("%v", val.alph3b)
			return &returnStr
		}
	}
	return nil
}

// Check map for Alpha2 Match and return ISO
// example, "en" will return "EN"
func Alpha2Match(s string) *string {
	var (
		returnStr string
	)
	if val, ok := isocodes2[s]; ok {
		returnStr = fmt.Sprintf("%v", val.alph2)
		return &returnStr
	}
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "xx" || s == "xxx" {
		return nil
	}
	for _, val := range isocodes2 {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3b) || s == strings.ToLower(val.alph3t) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
		// check against english characters
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
		// check against original text
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
	}
	return nil
}

// Check map for Alpha3 Match and return ISO
// example, "gb" will return GBR Struct
func StructMatch(s string) *ISOEntry {
	var (
		entry ISOEntry
	)
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "xx" || s == "xxx" {
		return nil
	}
	for _, val := range isocodes3 {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3b) || s == strings.ToLower(val.alph3t) {
			entry = ISOEntry{
				enName: val.enName,
				frName: val.frName,
				alph2:  val.alph2,
				alph3t: val.alph3t,
				alph3b: val.alph3b,
			}
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			entry = ISOEntry{
				enName: val.enName,
				frName: val.frName,
				alph2:  val.alph2,
				alph3t: val.alph3t,
				alph3b: val.alph3b,
			}
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			entry = ISOEntry{
				enName: val.enName,
				frName: val.frName,
				alph2:  val.alph2,
				alph3t: val.alph3t,
				alph3b: val.alph3b,
			}
			return &entry
		}

	}
	return nil
}
