package soteapi

import (
	"encoding/xml"
)

// GetProducerList returns producers list
func (api *SoteAPI) GetProducerList(sessionHash string, offset, limit int) ([]ProducerWithTimestamp, error) {
	req := getProducerListRequest{SessionHash: sessionHash, Offset: offset, Limit: limit}
	resp := getProducerListResponse{}
	err := serviceProducer.Call(stProducer, req, &resp, nil)
	return resp.Items, err
}

type getProducerListRequest struct {
	XMLName     xml.Name `xml:"GetProducerList"`
	SessionHash string   `xml:"GetProducerList>_session_hash"`
	Offset      int      `xml:"GetProducerList>_offset"`
	Limit       int      `xml:"GetProducerList>_limit"`
}

type getProducerListResponse struct {
	XMLName xml.Name                `xml:"GetProducerListResponse"`
	Items   []ProducerWithTimestamp `xml:"GetProducerListResponse>item"`
}
