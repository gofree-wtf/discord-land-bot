package main

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServiceKey struct {
		GetRTMSDataSvcAptTrade string `yaml:"getRTMSDataSvcAptTrade" env:"GET_RTMS_DATA_SVC_APT_TRADE"`
	} `yaml:"serviceKey" env:"SERVICE_KEY"`
}

var Cfg Config

func init() {
	loader := aconfig.LoaderFor(&Cfg, aconfig.Config{
		SkipFlags: true,
		Files:     []string{"config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})
	if err := loader.Load(); err != nil {
		log.Panic(err)
	}
}
