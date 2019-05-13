package yamoney

import (
	"regexp"
	"strings"
)

var noNumReg = regexp.MustCompile("\\D")

// ParsePhone - Привозим телефон к нужному формату
func ParsePhone(phone string) string {
	phone = noNumReg.ReplaceAllString(phone, "")

	if strings.HasPrefix(phone, "8") {
		phone = "7" + strings.TrimPrefix(phone, "8")
	}

	return phone
}
