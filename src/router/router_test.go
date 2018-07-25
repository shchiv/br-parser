package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var jsonStr = []byte(`
{
   "id":"FxlYjsHZs1",
   "at":2,
   "allimps":0,
   "imp":[
      {
         "id":"1",
         "video":{
            "w":320,
            "h":480,
            "companionad":[
               {
                  "h":420,
                  "w":320
               }
            ],
            "companiontype":[
               2
            ],
            "startdelay":0,
            "battr":[
               1,
               3,
               5,
               8,
               9,
               10,
               11
            ],
            "mimes":[
               "video/mp4",
               "video/3gpp"
            ],
            "linearity":1,
            "maxbitrate":4000,
            "minduration":5,
            "maxduration":60,
            "protocols":[
               2,
               5
            ],
            "pos":0
         },
         "ext":{
            "strictbannersize":1
         },
         "instl":0,
         "displaymanager":"SOMA",
         "tagid":"101000445",
         "secure":0,
         "didsha1":"31f455635a12252f8cedb39cd7533fc58a064bff",
         "didmd5":"af2d48f2495881aed1737bb21017f9b6",
         "dpidsha1":"a5ab171d2c5551c4b0552b6280b42b578e08e490",
         "dpidmd5":"285a23d862ebbec44bd5520643e2a9eb",
         "macsha1":"d76f4f17da290fb7b96d953a1b9a01a733cf4d9d",
         "macmd5":"f244cddbc7f5161e9dcd758d0c3d7796"
      }
   ],
   "device":{
      "geo":{
         "lat":-31.600006,
         "lon":-60.708298,
         "ipservice":3,
         "country":"ARG",
         "region":"21",
         "zip":"3000",
         "metro":"0",
         "city":"Santa Fe",
         "type":2
      },
      "make":"Google",
      "model":"Nexus 6P",
      "os":"Android",
      "osv":"6.0",
      "ua":"Mozilla/5.0 (Linux; Android 6.0; Nexus 6P Build/MDA83) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.94 Mobile Safari/537.36",
      "ip":"201.252.0.0",
      "js":0,
      "connectiontype":0,
      "devicetype":1
   },
   "app":{
      "id":"101000445",
      "name":"App name",
      "domain":"play.google.com",
      "cat":[
         "IAB3"
      ],
      "bundle":"com.ximad.snake",
      "storeurl":"https://play.google.com/store/apps/details?id=com.ximad.snake",
      "keywords":"",
      "publisher":{
         "id":"1001028764",
         "name":"Publisher name"
      }
   },
   "bcat":[
      "IAB17-18",
      "IAB7-42",
      "IAB23",
      "IAB7-28",
      "IAB26",
      "IAB25",
      "IAB9-9",
      "IAB24"
   ],
   "badv":[

   ],
   "ext":{
      "udi":{
         "androidid":"3840000210b96e40",
         "androididmd5":"285a23d862ebbec44bd5520643e2a9eb",
         "androididsha1":"a5ab171d2c5551c4b0552b6280b42b578e08e490",
         "imei":"356938035643809",
         "imeimd5":"af2d48f2495881aed1737bb21017f9b6",
         "imeisha1":"31f455635a12252f8cedb39cd7533fc58a064bff",
         "macmd5":"f244cddbc7f5161e9dcd758d0c3d7796",
         "macsha1":"d76f4f17da290fb7b96d953a1b9a01a733cf4d9d",
         "odin":"d76f4f17da290fb7b96d953a1b9a01a733cf4d9d"
      },
      "operaminibrowser":0,
      "carriername":"Telecom Argentina"
   },
   "regs":{
      "coppa":0,
      "ext":{
         "gdpr":1
      }
   },
   "user":{
      "keywords":"",
      "ext":{
         "consent":"BONZC5sONZC5sAAABAENAAoAAAAFIgAAAAAAAAAAAAI"
      }
   }
}`)

func TestParserHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/parser", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	r, err := CreateReader()
	if err != nil {
		//TODO add ability to use reader in tests
		//t.Fatal(err)
	}

	server := NewServer(r)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.parserHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"OS":"Android 6.0","DeviceType":"Mobile","Browser":"Chrome","Country":"undefined","Domain":"host0.201-252-0.telecom.net.ar."}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}