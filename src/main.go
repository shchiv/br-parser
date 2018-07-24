package main

import (
	"encoding/json"
	"fmt"
	"github.com/br-parser/src/model"
	"github.com/bsm/openrtb"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
)

func main() {
	/*err := readDir("./testdata")
	if err != nil {
		log.Fatal(err)
	}*/

	http.HandleFunc("/parser", parserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parserHandler(w http.ResponseWriter, r *http.Request) {
	var req *openrtb.BidRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR %s", err))
		return
	}

	ua := user_agent.New(req.Device.UA)

	var resp model.ParserResponse
	resp.OS = ua.OS()
	resp.DeviceType = getDeviceType(req.Device.DeviceType)
	resp.Browser = getBrowser(ua)
	resp.Country = getCountry(req.Device.IP)
	resp.Domain = getDomain(req.Device.IP)

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(fmt.Sprintf("ERROR %s", err.Error()))
	}
}

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
		return "Unknown"
	}
}

func getBrowser(ua *user_agent.UserAgent) string {
	browser, version := ua.Browser()
	return fmt.Sprintf("%s %s", browser, version)
}

func getCountry(IP string) string {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	IPaddr := net.ParseIP(IP)
	if IPaddr != nil {
		record, err := db.Country(IPaddr)
		if err != nil {
			log.Println(fmt.Sprintf("Error %s", err.Error()))
			return "undefined"
		} else {
			log.Println(fmt.Sprintf("Country %s", record.Country.Names["en"]))
			return record.Country.Names["en"]
		}
	} else {
		//TODO handle this
	}
	return "undefined"
}

func getDomain(IP string) string {
	addr, err := net.LookupAddr(IP)
	if err != nil {
		return "undefined"
	}

	if len(addr) > 0 {
		return addr[0]
	} else {
		return "undefined"
	}
}

/*func readDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			if err := readDir(path + "/" + file.Name()); err != nil {
				log.Printf("Error on read dir %s Error %s", file.Name(), err)
				continue
			}
		} else {
			val, err := parseFile(path + "/" + file.Name())
			if err != nil {
				log.Printf("Error on parse file %s Error %s", file.Name(), err)
				continue
			}
			fmt.Println(fmt.Sprintf("File %s Browser %s", file.Name(), val))
		}
	}

	return nil
}

func parseFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var req *openrtb.BidRequest
	err = json.NewDecoder(file).Decode(&req)
	if err != nil {
		return "", nil
	}

	ua := user_agent.New(req.Device.UA)
	name, version := ua.Browser()
	ua.Mobile()
	return fmt.Sprintf("%s %s", name, version), nil
}*/
