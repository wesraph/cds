package messenger

import (
	"github.com/ovh/cds/engine/api"
	"github.com/ovh/cds/engine/service"
)

// Service is messenger service
type Service struct {
	service.Common
	Cfg    Configuration
	Router *api.Router
}

// Configuration is the vcs configuration structure
type Configuration struct {
	Name string `toml:"name" comment:"Name of this CDS Messenger Service\n Enter a name and token below to enable this service" json:"name"`
	HTTP struct {
		Addr string `toml:"addr" default:"" commented:"true" comment:"Listen address without port, example: 127.0.0.1" json:"addr"`
		Port int    `toml:"port" default:"8089" json:"port"`
	} `toml:"http" comment:"######################\n CDS Messenger HTTP Configuration \n######################" json:"http"`
	URL   string `default:"http://localhost:8089" json:"url"`
	Hubot struct {
		URL string `toml:"url" json:"url"`
	} `toml:"hubot" comment:"######################\n CDS Hubot Settings \n######################" json:"hubot"`
	API service.APIServiceConfiguration `toml:"api" comment:"######################\n CDS API Settings \n######################" json:"api"`
}
