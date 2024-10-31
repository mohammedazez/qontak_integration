package utils

import (
	"encoding/json"
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

func ObjectToObject(in interface{}, out interface{}) {
	dataByte, _ := json.Marshal(in)
	json.Unmarshal(dataByte, &out)
}

func ObjectToString(data interface{}) string {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dataByte)
}

func StringToObject(in string, out interface{}) {
	json.Unmarshal([]byte(in), &out)
	return
}
func GetStringBetween(str, start, end string) string {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return ""
	}

	endIndex := strings.Index(str[startIndex+len(start):], end)
	if endIndex == -1 {
		return ""
	}

	result := str[startIndex+len(start) : startIndex+len(start)+endIndex]
	return result
}

func FulfillTemplate(template string, parameters map[string]interface{}) string {
	if template == "" {
		return template
	}
	for key, element := range parameters {
		template = strings.Replace(template, fmt.Sprintf("{{%v}}", key), element.(string), -1)
	}
	return template
}
func ConvertPhoneNumber(mobilePhoneNumber string) (newMobilePhoneNumber string, err error) {
	phoneNumber := strings.Replace(mobilePhoneNumber, " ", "", -1)
	if strings.HasPrefix(phoneNumber, "62") {
		newMobilePhoneNumber = strings.Replace(phoneNumber, "62", "0", 1)
	} else if strings.HasPrefix(phoneNumber, "+62") {
		newMobilePhoneNumber = strings.Replace(phoneNumber, "+62", "0", 1)
	} else if strings.HasPrefix(phoneNumber, "0") {
		newMobilePhoneNumber = phoneNumber
	} else {
		newMobilePhoneNumber = "0" + phoneNumber
	}
	valid := NewValidator()
	if err = valid.Var(newMobilePhoneNumber, "numeric"); err != nil {
		newMobilePhoneNumber = ""
		err = Error.New(http.StatusBadRequest, "FAILED", "Mobile Phone Number Tidak Valid")
		return
	}

	return
}

func ConvertDate(date, format_need string) (result string, err error) {
	listformats := []string{
		"02 January 2006",
		"02 Jan 2006",
		"02-01-2006",
		"02/01/2006",
		"01/02/2006",
		"2006-01-02",
		"January 02, 2006",
		"Jan 02, 2006",
		"2006-01-02 15:04:05",
	}

	// Parsing tanggal dengan berbagai format
	var t time.Time

	for _, format := range listformats {
		t, err = time.Parse(format, date)
		if err == nil {
			break
		}
	}

	// Cek apakah parsing berhasil
	if err != nil {
		err = errors.New(fmt.Sprintf("Gagal memparsing tanggal: %v", err.Error()))
		return
	}

	// Mengonversi tanggal ke format yang diinginkan
	result = t.Format(format_need)
	return
}
