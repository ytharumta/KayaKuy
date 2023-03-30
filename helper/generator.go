package helper

import (
	"encoding/base64"
	"github.com/thanhpk/randstr"
	"strconv"
)

func GenerateCode(UserId int64) string {
	StringToEncode := randstr.String(10) + strconv.Itoa(int(UserId))

	Encoding := base64.StdEncoding.EncodeToString([]byte(StringToEncode))
	return Encoding
}
