package soteapi

import (
	"encoding/xml"
)

// GetOrderCount returns orders count
func (api *SoteAPI) GetOrderCount(sessionHash string) (int, error) {
	req := getOrderCountRequest{SessionHash: sessionHash}
	resp := getOrderCountResponse{}
	err := serviceOrder.Call(stOrder, req, &resp, nil)
	return resp.Count, err
}

type getOrderCountRequest struct {
	XMLName     xml.Name `xml:"GetOrderCount"`
	SessionHash string   `xml:"GetOrderCount>_session_hash"`
}

type getOrderCountResponse struct {
	XMLName xml.Name `xml:"GetOrderCountResponse"`
	Count   int      `xml:"GetOrderCountResponse>_count"`
}
