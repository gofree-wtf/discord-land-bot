package main

import "encoding/xml"

type GetRTMSDataSvcAptTrade struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"` // 99
		ResultMsg  string `xml:"resultMsg"`  // SERVICE KEY IS NOT REGIST...
	} `xml:"header"`
}
