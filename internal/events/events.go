package events

import "strings"

func EventsRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func EventsRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func EventsRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func EventsRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func EventsRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func EventsRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func EventsRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func EventsRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func EventsRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func EventsRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func EventsRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func EventsRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "events-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
