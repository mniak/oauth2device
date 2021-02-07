package googledevice

import (
	"github.com/mniak/oauth2device"
)

// Endpoint specifies the endpoints used for the Google OAuth
// 2.0 for devices.
var DeviceEndpoint = oauth2device.DeviceEndpoint{
	CodeURL: "https://accounts.google.com/o/oauth2/device/code",
}
