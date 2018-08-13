package soteapi

import (
	"encoding/xml"
)

// DoLogin login to Soteshop and return session token
func (api *SoteAPI) DoLogin(username, password string) (string, error) {
	req := doLoginRequest{Username: username, Password: password}
	resp := doLoginResponse{}
	err := serviceWebAPIBackend.Call(stWebAPIBackend, req, &resp, nil)
	return resp.Hash, err
}

type doLoginRequest struct {
	XMLName  xml.Name `xml:"doLogin"`
	Username string   `xml:"doLogin>username"`
	Password string   `xml:"doLogin>password"`
}

type doLoginResponse struct {
	XMLName xml.Name `xml:"doLoginResponse"`
	Hash    string   `xml:"doLoginResponse>hash"`
}
