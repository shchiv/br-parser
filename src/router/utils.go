package router

import (
	"encoding/json"
	"errors"
	"github.com/bsm/openrtb"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
)

const (
	undefined     = "undefined"
	defaultLocale = "en"
)

func getDeviceType(deviceType int) string {
	switch deviceType {
	case openrtb.DeviceTypeMobile:
		return "Mobile"
	case openrtb.DeviceTypePC:
		return "PC"
	case openrtb.DeviceTypeTV:
		return "TV"
	case openrtb.DeviceTypePhone:
		return "Phone"
	case openrtb.DeviceTypeTablet:
		return "Tablet"
	case openrtb.DeviceTypeConnected:
		return "Connected"
	case openrtb.DeviceTypeSetTopBox:
		return "SetTopBox"
	default:
		return undefined
	}
}

func getBrowser(ua *user_agent.UserAgent) string {
	if ua == nil {
		log.Println("Can't get browser. User agent is nil")
		return undefined
	}
	browser, _ := ua.Browser()
	return browser
}

func getCountry(reader *geoip2.Reader, IP string) string {
	if reader == nil {
		log.Printf("Can't get country for IP %s. Reader is nil", IP)
		return undefined
	}

	IPaddr := net.ParseIP(IP)
	if IPaddr != nil {
		record, err := reader.Country(IPaddr)
		if err != nil {
			log.Printf("Can't get country for IP %s. Error %s", IP, err.Error())
			return undefined
		} else if record != nil {
			if name, ok := record.Country.Names[defaultLocale]; ok {
				return name
			} else {
				log.Printf("Can't get country for locale %s and IP %s", defaultLocale, IP)
				return undefined
			}
		} else {
			log.Printf("Can't get country for IP %s", IP)
			return undefined
		}
	} else {
		log.Printf("Can't parse IP %s for getting country", IP)
		return undefined
	}
}

func getDomain(IP string) string {
	addr, err := net.LookupAddr(IP)
	if err != nil {
		log.Printf("Can't get domain for IP %s. Error %s", IP, err.Error())
		return undefined
	}

	if len(addr) > 0 {
		return addr[0]
	} else {
		log.Printf("Can't get domain for IP %s", IP)
		return undefined
	}
}

func renderResponse(statusCode int, w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("Can't render response with Status %d Value %+v. Error %s", statusCode, value, err.Error())
	}
}

func CreateReader() (*geoip2.Reader, error) {
	if reader, err := geoip2.Open("./resources/GeoLite2-Country.mmdb"); err != nil {
		return nil, err
	} else if reader == nil {
		return nil, errors.New("GeoIP reader is nil")
	} else {
		return reader, nil
	}
}
