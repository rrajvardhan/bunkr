package dashboard

import (
	"github.com/skip2/go-qrcode"
)

func GenerateCode(url string) string {
	ascii, _ := qrcode.New(url, qrcode.High)
	return ascii.ToSmallString(true)
}
