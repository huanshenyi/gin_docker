package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"gin_docker/src/domain"
)

// デバイス情報をgin.contextのuserに追加
func NewDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := getDeviceFromHeaderValue(c.GetHeader(domain.ClientAppTypeHTTPHeaderKey))
		// user, found := authenticator.GetUser(c)
		// if !found {
		// 	return
		// }
		c.Set("device", device)
	}
}

func getDeviceFromHeaderValue(handerValue string) domain.Device {
	var device domain.Device
	switch v := strings.ToLower(strings.TrimSpace(handerValue)); v {
	case domain.IOS.String():
		device.Platform = domain.IOS
	case domain.Android.String():
		device.Platform = domain.Android
	case domain.Web.String():
		device.Platform = domain.Web
	default:
		device.Platform = domain.UndefinedPlatform
	}
	return device
}
