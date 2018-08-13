package soteapi

import (
	"encoding/xml"
)

// GetOrderList returns orders list
func (api *SoteAPI) GetOrderList(sessionHash string, offset, limit int) ([]OrderWithTimestamp, error) {
	req := getOrderListRequest{SessionHash: sessionHash, Offset: offset, Limit: limit}
	resp := getOrderListResponse{}
	err := serviceOrder.Call(stOrder, req, &resp, nil)
	return resp.Items, err
}

type getOrderListRequest struct {
	XMLName     xml.Name `xml:"GetOrderList"`
	SessionHash string   `xml:"GetOrderList>_session_hash"`
	Offset      int      `xml:"GetOrderList>_offset"`
	Limit       int      `xml:"GetOrderList>_limit"`
}

type getOrderListResponse struct {
	XMLName xml.Name             `xml:"GetOrderListResponse"`
	Items   []OrderWithTimestamp `xml:"GetOrderListResponse>item"`
}
