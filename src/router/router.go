package router

import (
	"encoding/json"
	"github.com/br-parser/src/model"
	"github.com/bsm/openrtb"
	"github.com/mssola/user_agent"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	parserURL = "/parser"
)

type server struct {
	reader *geoip2.Reader
}

func NewServer(reader *geoip2.Reader) *server {
	s := new(server)
	s.reader = reader
	return s
}

func (s *server) Start() {
	router := mux.NewRouter()
	router.HandleFunc(parserURL, s.parserHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func (s *server) parserHandler(w http.ResponseWriter, r *http.Request) {
	var req *openrtb.BidRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		renderResponse(http.StatusForbidden, w, &model.ParserError{ErrMsg: err.Error()})
		log.Printf("Can't unmarshal request. Error %s", err)
		return
	}

	if req == nil {
		renderResponse(http.StatusForbidden, w, &model.ParserError{ErrMsg: "Empty request"})
		return
	}

	if req.Device == nil {
		renderResponse(http.StatusForbidden, w, &model.ParserError{ErrMsg: "Device information not found"})
		return
	}

	ua := user_agent.New(req.Device.UA)
	if ua == nil {
		renderResponse(http.StatusForbidden, w, &model.ParserError{"Can't parse user agent"})
		return
	}

	var resp model.ParserResponse
	resp.OS = ua.OS()
	resp.DeviceType = getDeviceType(req.Device.DeviceType)
	resp.Browser = getBrowser(ua)
	resp.Country = getCountry(s.reader, req.Device.IP)
	resp.Domain = getDomain(req.Device.IP)

	renderResponse(http.StatusOK, w, &resp)
}
