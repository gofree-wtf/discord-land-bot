package main

import (
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {
	u := &url.URL{
		Scheme: "http",
		Host:   "openapi.molit.go.kr:8081",
		Path:   "/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTrade",
	}

	q := u.Query()
	q.Set("serviceKey", Cfg.ServiceKey.GetRTMSDataSvcAptTrade)
	q.Set("LAWD_CD", "11110")
	q.Set("DEAL_YMD", "202103")
	u.RawQuery = q.Encode()

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	log.Info(string(body))

	data := &GetRTMSDataSvcAptTrade{}
	err = xml.Unmarshal(body, data)
	assert.NoError(t, err)

	log.Info(data)
}
