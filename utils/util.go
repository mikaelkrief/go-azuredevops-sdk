package utils

import (
	"encoding/base64"
)

func BasicAuth(token string) string {
	auth := token
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
