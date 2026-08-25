package exporter

import "strings"

func ExporterRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func ExporterRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func ExporterRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func ExporterRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func ExporterRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func ExporterRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func ExporterRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func ExporterRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func ExporterRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func ExporterRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func ExporterRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func ExporterRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "exporter-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
