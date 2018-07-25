package router

import (
	"github.com/bsm/openrtb"
	"github.com/mssola/user_agent"
	. "github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetDeviceType(t *testing.T) {
	result := getDeviceType(openrtb.DeviceTypeMobile)
	Equal(t, "Mobile", result)

	result = getDeviceType(openrtb.DeviceTypePC)
	Equal(t, "PC", result)

	result = getDeviceType(openrtb.DeviceTypeTV)
	Equal(t, "TV", result)

	result = getDeviceType(openrtb.DeviceTypePhone)
	Equal(t, "Phone", result)

	result = getDeviceType(openrtb.DeviceTypeTablet)
	Equal(t, "Tablet", result)

	result = getDeviceType(openrtb.DeviceTypeConnected)
	Equal(t, "Connected", result)

	result = getDeviceType(openrtb.DeviceTypeSetTopBox)
	Equal(t, "SetTopBox", result)

	result = getDeviceType(openrtb.DeviceTypeUnknown)
	Equal(t, "undefined", result)

	result = getDeviceType(-1)
	Equal(t, "undefined", result)

	result = getDeviceType(math.MaxInt64)
	Equal(t, "undefined", result)

	result = getDeviceType(00000)
	Equal(t, "undefined", result)

	result = getDeviceType(111111)
	Equal(t, "undefined", result)
}

func TestGetBrowser(t *testing.T) {
	uaStr := "Mozilla/5.0 (Linux; Android 6.0; Nexus 6P Build/MDA83) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.94 Mobile Safari/537.36"
	ua := user_agent.New(uaStr)
	if ua != nil {
		result := getBrowser(ua)
		Equal(t, "Chrome", result)
	} else {
		t.Errorf("User agent not valid")
	}
}

func TestGetBrowserNilUserAgent(t *testing.T) {
	result := getBrowser(nil)
	Equal(t, "undefined", result)
}

func TestGetBrowserWithoutVersion(t *testing.T) {
	uaStr := "Mozilla/5.0 (Linux; Android 6.0; Nexus 6P Build/MDA83) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome Mobile Safari/537.36"
	ua := user_agent.New(uaStr)
	if ua != nil {
		result := getBrowser(ua)
		Equal(t, "Chrome", result)
	} else {
		t.Errorf("User agent not valid")
	}
}

func TestGetBrowserEmptyUserAgent(t *testing.T) {
	ua := user_agent.New("")
	if ua != nil {
		result := getBrowser(ua)
		Equal(t, "", result)
	} else {
		t.Errorf("User agent not valid")
	}
}

func TestGetCountry(t *testing.T) {
	result := getCountry(nil, "")
	Equal(t, "undefined", result)
}

func TestGetDomainLocalIP(t *testing.T) {
	result := getDomain("192.168.0.11")
	Equal(t, "undefined", result)
}

func TestGetDomainRealIP(t *testing.T) {
	result := getDomain("77.123.139.189")
	Equal(t, "2ip.ua.", result)
}
