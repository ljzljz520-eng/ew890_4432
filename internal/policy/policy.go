package policy

import "strings"

func PolicyRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func PolicyRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func PolicyRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func PolicyRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func PolicyRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func PolicyRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func PolicyRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func PolicyRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func PolicyRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func PolicyRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func PolicyRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func PolicyRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "policy-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
