package soteapi

import (
	"encoding/xml"
)

// GetVersion returns Soteshop API version
func (api *SoteAPI) GetVersion() (string, error) {
	req := getVersionRequest{}
	resp := getVersionResponse{}
	err := serviceWebAPIBackend.Call(api.SoteShopURL, req, &resp, nil)
	if err != nil {
		return "", err
	}
	return resp.Version, nil
}

type getVersionRequest struct {
	XMLName xml.Name `xml:"getVersion"`
	Culture string   `xml:"getVersion>_culture"`
}

type getVersionResponse struct {
	XMLName xml.Name `xml:"getVersionResponse"`
	Version string   `xml:"getVersionResponse>version"`
}
