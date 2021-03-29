package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type GetRTMSDataSvcAptTradeResponse struct {
	XMLName xml.Name `xml:"response"`
	Header  struct {
		ResultCode uint   `xml:"resultCode"` // 00
		ResultMsg  string `xml:"resultMsg"`  // NORMAL SERVICE.
	} `xml:"header"`
	Body struct {
		Items struct {
			Item []*AptTrade `xml:"item"`
		} `xml:"items"`
		NumOfRows  uint `xml:"numOfRows"`
		PageNo     uint `xml:"pageNo"`
		TotalCount uint `xml:"totalCount"`
	} `xml:"body"`
}

type AptTrade struct {
	ApartmentName string  `xml:"아파트"`  // 광화문풍림스페이스 본(9-0)
	BuildYear     uint    `xml:"건축년도"` // 2015
	Floor         uint    `xml:"층"`    // 11
	ExclusiveArea float32 `xml:"전용면적"` // 94.51

	RegionalCode string `xml:"지역코드"` // 11110
	Dong         string `xml:"법정동"`  // 사직동
	Jibun        string `xml:"지번"`   // 9

	DealYear      uint   `xml:"년"`    // 2015
	DealMonth     uint   `xml:"월"`    // 12
	DealDay       uint   `xml:"일"`    // 1
	DealAmountStr string `xml:"거래금액"` // 82,500 (만원)
	DealAmount    uint   // 82500 * 10000

	CancelDealType string `xml:"해제여부"`    // O
	CancelDealDay  string `xml:"해제사유발생일"` // 21.01.27
}

func (r *GetRTMSDataSvcAptTradeResponse) Validate() error {
	if r.Header.ResultCode != 0 {
		return fmt.Errorf("error code: %d", r.Header.ResultCode)
	}

	for _, item := range r.Body.Items.Item {
		item.ApartmentName = strings.Trim(item.ApartmentName, " ")
		item.RegionalCode = strings.Trim(item.RegionalCode, " ")
		item.Dong = strings.Trim(item.Dong, " ")
		item.Jibun = strings.Trim(item.Jibun, " ")

		item.DealAmountStr = strings.Trim(item.DealAmountStr, " ")
		{
			dealAmount, err := strconv.Atoi(strings.Replace(item.DealAmountStr, ",", "", -1))
			if err != nil {
				return err
			}
			item.DealAmount = uint(dealAmount) * 10_000
		}

		item.CancelDealType = strings.Trim(item.CancelDealType, " ")
		item.CancelDealDay = strings.Trim(item.CancelDealDay, " ")
	}
	return nil
}
