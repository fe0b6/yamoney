package yamoney

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var noNumReg = regexp.MustCompile("\\D")

// ParsePhone - Привозим телефон к нужному формату
func ParsePhone(phone string) int {
	phone = noNumReg.ReplaceAllString(phone, "")

	if strings.HasPrefix(phone, "8") {
		phone = "7" + strings.TrimPrefix(phone, "8")
	}

	ans, err := strconv.Atoi(phone)
	if err != nil {
		log.Println("[error]", err, phone)
		return ans
	}

	return ans
}
