package soteapi

import (
	"encoding/xml"
)

// GetOrder returns order by id
func (api *SoteAPI) GetOrder(sessionHash string, id int) (*OrderWithTimestamp, error) {
	req := getOrderRequest{SessionHash: sessionHash, ID: id}
	resp := getOrderResponse{}
	err := serviceOrder.Call(stOrder, req, &resp, nil)
	return &resp.Data, err
}

type getOrderRequest struct {
	XMLName     xml.Name `xml:"GetOrder"`
	SessionHash string   `xml:"GetOrder>_session_hash"`
	ID          int      `xml:"GetOrder>id"`
}

type getOrderResponse struct {
	XMLName xml.Name           `xml:"GetOrderResponse"`
	Data    OrderWithTimestamp `xml:"GetOrderResponse"`
}
