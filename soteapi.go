package soteapi

import (
	"fmt"

	soapc "github.com/achiku/soapc"
)

// SoteAPI represents a SoteShop API
type SoteAPI struct {
	SoteShopURL string
	IsHTTPS     bool
}

const (
	stWebAPIBackend = "stWebApiBackend"
	stProducer      = "stProducer"
	stCategory      = "stCategory"
	stOrder         = "stOrder"
)

var (
	serviceWebAPIBackend *soapc.Client
	serviceProducer      *soapc.Client
	serviceCategory      *soapc.Client
	serviceOrder         *soapc.Client
)

// New creates a new SoteAPI client
func New(url string, isHTTPS bool) *SoteAPI {
	api := &SoteAPI{SoteShopURL: url, IsHTTPS: isHTTPS}
	serviceWebAPIBackend = api.newSoapClient(stWebAPIBackend)
	serviceProducer = api.newSoapClient(stProducer)
	serviceCategory = api.newSoapClient(stProducer)
	serviceOrder = api.newSoapClient(stOrder)
	return api
}

func (api *SoteAPI) newSoapClient(module string) *soapc.Client {
	url := api.makeServiceURL(module)
	return soapc.NewClient(url, api.IsHTTPS, nil)
}

func (api *SoteAPI) makeServiceURL(module string) string {
	return fmt.Sprintf("%s/backend.php/%s/soap", api.SoteShopURL, module)
}
