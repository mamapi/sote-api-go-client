package soteapi

import (
	"encoding/xml"
)

// GetProducer returns producer by id
func (api *SoteAPI) GetProducer(sessionHash string, id int) (*ProducerWithTimestamp, error) {
	req := getProducerRequest{SessionHash: sessionHash, ID: id}
	resp := getProducerResponse{}
	err := serviceProducer.Call(stProducer, req, &resp, nil)
	return &resp.Data, err
}

type getProducerRequest struct {
	XMLName     xml.Name `xml:"GetProducer"`
	SessionHash string   `xml:"GetProducer>_session_hash"`
	ID          int      `xml:"GetProducer>id"`
}

type getProducerResponse struct {
	XMLName xml.Name              `xml:"GetProducerResponse"`
	Data    ProducerWithTimestamp `xml:"GetProducerResponse"`
}
