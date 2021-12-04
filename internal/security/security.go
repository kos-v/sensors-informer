package security

import (
	"strings"
	"unicode/utf8"
)

func HideSecrets(secrets []string, target string) string {
	result := target
	for _, secret := range secrets {
		result = strings.Replace(result, secret, strings.Repeat("*", utf8.RuneCountInString(secret)), -1)
	}

	return result
}
