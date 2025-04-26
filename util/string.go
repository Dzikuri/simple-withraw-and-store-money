package util

import (
	"encoding/json"
	"fmt"
	"time"
	"unicode"
)

func LogPretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b) + "\n")
}

func ValidPassword(s string) error {
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		//"numeric":	{unicode.Number, unicode.Digit},
		"special": {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range s {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}

func GenerateInvoiceID(prefix string) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%x", prefix, timestamp)
}

// GetStringPointer mengembalikan pointer ke string jika tidak kosong, atau nil jika kosong
func GetStringPointer(s *string) *string {
	if s == nil || *s == "" {
		return nil
	}
	return s
}
