package soteapi

import (
	"encoding/xml"
)

// CountProducer returns producers count
func (api *SoteAPI) CountProducer(sessionHash string) (int, error) {
	req := countProducerRequest{SessionHash: sessionHash}
	resp := countProducerResponse{}
	err := serviceCategory.Call(stProducer, req, &resp, nil)
	return resp.Count, err
}

type countProducerRequest struct {
	XMLName     xml.Name `xml:"CountProducer"`
	SessionHash string   `xml:"CountProducer>_session_hash"`
}

type countProducerResponse struct {
	XMLName xml.Name `xml:"CountProducerResponse"`
	Count   int      `xml:"CountProducerResponse>_count"`
}
