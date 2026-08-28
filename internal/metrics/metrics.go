package metrics

import "strings"

func MetricsRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func MetricsRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func MetricsRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func MetricsRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func MetricsRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func MetricsRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func MetricsRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func MetricsRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func MetricsRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func MetricsRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func MetricsRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func MetricsRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "metrics-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
