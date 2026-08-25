package validation

import "strings"

func ValidationRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func ValidationRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func ValidationRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func ValidationRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func ValidationRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func ValidationRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func ValidationRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func ValidationRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func ValidationRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func ValidationRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func ValidationRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func ValidationRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "validation-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
