package soteapi

import (
	"encoding/xml"
)

// Test test communication
func (api *SoteAPI) Test(echo string) (string, error) {
	req := testRequest{Echo: echo}
	resp := testResponse{}
	err := serviceWebAPIBackend.Call(stWebAPIBackend, req, &resp, nil)
	if err != nil {
		return "", err
	}
	return resp.Echo, nil
}

type testRequest struct {
	XMLName xml.Name `xml:"test"`
	Echo    string   `xml:"test>echo"`
}

type testResponse struct {
	XMLName xml.Name `xml:"testResponse"`
	Echo    string   `xml:"testResponse>echo"`
}
