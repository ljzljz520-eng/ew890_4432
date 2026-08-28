package ledger

import "strings"

func LedgerRule1(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-1"
	}
	if len(v) > 21 {
		return v[:21]
	}
	return v
}
func LedgerRule2(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-2"
	}
	if len(v) > 22 {
		return v[:22]
	}
	return v
}
func LedgerRule3(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-3"
	}
	if len(v) > 23 {
		return v[:23]
	}
	return v
}
func LedgerRule4(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-4"
	}
	if len(v) > 24 {
		return v[:24]
	}
	return v
}
func LedgerRule5(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-5"
	}
	if len(v) > 25 {
		return v[:25]
	}
	return v
}
func LedgerRule6(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-6"
	}
	if len(v) > 26 {
		return v[:26]
	}
	return v
}
func LedgerRule7(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-7"
	}
	if len(v) > 27 {
		return v[:27]
	}
	return v
}
func LedgerRule8(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-8"
	}
	if len(v) > 28 {
		return v[:28]
	}
	return v
}
func LedgerRule9(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-9"
	}
	if len(v) > 29 {
		return v[:29]
	}
	return v
}
func LedgerRule10(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-10"
	}
	if len(v) > 30 {
		return v[:30]
	}
	return v
}
func LedgerRule11(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-11"
	}
	if len(v) > 31 {
		return v[:31]
	}
	return v
}
func LedgerRule12(value string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return "ledger-missing-12"
	}
	if len(v) > 32 {
		return v[:32]
	}
	return v
}
